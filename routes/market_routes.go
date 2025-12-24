package routes

import (
	"log"

	"github.com/ayonqfl/go-fiber-gorm/database"
	"github.com/ayonqfl/go-fiber-gorm/helpers"
	qdb "github.com/ayonqfl/go-fiber-gorm/models/qdb"
	"github.com/ayonqfl/go-fiber-gorm/utils"
	"github.com/gofiber/fiber/v2"
)

func MarketHandlers(route fiber.Router) {
	// Define market watchlist API function
	route.Get("/watchlist", func(c *fiber.Ctx) error {
		// tableID := c.Query("table_id")

		currentUser, err := helpers.GetCurrentUser(c)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		userID := currentUser.UserID
		username := currentUser.Username
		currentUserID := currentUser.ID
		// log.Printf("User: %s ID: %s", username, userID)

		watchlistResult := []string{}
		var customWatchlistNames []string
		err = database.GetQtraderDB().Model(&qdb.Watchlist{}).
			Select("DISTINCT watchlist_name").
			Where("cln_id = ? AND watchlist_name IS NOT NULL AND watchlist_name != ''", userID).
			Order("watchlist_name ASC").
			Pluck("watchlist_name", &customWatchlistNames).Error

		if err != nil {
			log.Printf("Error querying watchlist: %v", err)
			return utils.SendResponse(c, 500, utils.ResponseOptions{
				Errors: "Failed to fetch watchlist",
			})
		}

		watchlistResult = append(watchlistResult, customWatchlistNames...)

		// -- config data extraction
		if helpers.GetEnvBool("ASSIGNED_DEFAULT_WATCHLIST", false) {
			var groupID uint
			err := database.GetQtraderDB().
				Model(&qdb.RmsGroupList{}).
				Select("group_id").
				Where("group_value = ?", username).
				Limit(1).
				Scan(&groupID).Error

			if err != nil {
				log.Printf("Failed to fetch group_id: %v", err)
			} else {
				var defaultWatchlistNames []string

				err = database.GetQtraderDB().
					Model(&qdb.DefaultWatchlist{}).
					Select("DISTINCT default_watchlist.name").
					Joins("LEFT JOIN default_watchlist_mapping ON default_watchlist.id = default_watchlist_mapping.watchlist_id").
					Where(`default_watchlist.type = ?
						OR (default_watchlist_mapping.type = ? AND default_watchlist_mapping.group_id = ?)
						OR (default_watchlist_mapping.type = ? AND default_watchlist_mapping.group_id = ?)`,
						"all",
						qdb.WatchlistMappingUser, currentUserID,
						qdb.WatchlistMappingGroup, groupID,
					).
					Order("default_watchlist.name ASC").
					Pluck("default_watchlist.name", &defaultWatchlistNames).Error

				if err != nil {
					log.Printf("Failed to fetch default watchlists: %v", err)
				} else {
					watchlistResult = append(watchlistResult, defaultWatchlistNames...)
				}
			}
		}

		if currentUser.UsersRoles == "client" {
			watchlistResult = append(watchlistResult, "PORTFOLIO")
		}
		watchlistResult = append(watchlistResult,
			"BOND (Public)",
			"SC",
			"ATB",
			"BOND (Government)",
			"SPOT MKT",
			"BLOCK",
			"Z CATEGORY",
			"SUSPEND",
		)
		return utils.SendResponse(c, 200, utils.ResponseOptions{
			Message: "Success",
			Data:    watchlistResult,
		})
	})

}
