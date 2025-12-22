package routes

import (
	"os"
	"log"
	"strings"

	"github.com/ayonqfl/go-fiber-gorm/database"
	"github.com/ayonqfl/go-fiber-gorm/models/qdb"
	"github.com/ayonqfl/go-fiber-gorm/utils"
	"github.com/gofiber/fiber/v2"
)

func MarketHandlers(route fiber.Router) {
	// Define market watchlist API function
	route.Get("/watchlist", func(c *fiber.Ctx) error {
		// tableID := c.Query("table_id")

		currentUser, err := utils.GetCurrentUser(c)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		userID := currentUser.UserID
		username := currentUser.Username
		// log.Printf("User: %s ID: %s", username, userID)

		watchlistResult := []string{}
		var customWatchlistNames []string
		err = database.GetQtraderDB().Model(&models.Watchlist{}).
			Select("DISTINCT watchlist_name").
			Where("cln_id = ? AND watchlist_name IS NOT NULL AND watchlist_name != ''", userID).
			Order("watchlist_name ASC").
			Pluck("watchlist_name", &customWatchlistNames).Error

		if err != nil {
			log.Printf("Error querying watchlist: %v", err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to fetch watchlist",
			})
		}

		watchlistResult = append(watchlistResult, customWatchlistNames...)

		// config data extraction
		assignedDefaultWatchlist := os.Getenv("ASSIGNED_DEFAULT_WATCHLIST")
		if strings.ToLower(strings.TrimSpace(assignedDefaultWatchlist)) == "true" {
			var groupID uint
			err := database.GetQtraderDB().
				Model(&models.RmsGroupList{}).
				Select("group_id").
				Where("group_value = ?", username).
				Limit(1).
				Scan(&groupID).Error

			if err != nil {
				log.Printf("Failed to fetch group_id: %v", err)
			} else {
				var defaultWatchlistNames []string

				err = database.GetQtraderDB().
					Model(&models.DefaultWatchlist{}).
					Select("DISTINCT default_watchlist.name").
					Joins("LEFT JOIN default_watchlist_mapping ON default_watchlist.id = default_watchlist_mapping.watchlist_id").
					Where(`default_watchlist.type = ? 
						OR (default_watchlist_mapping.type = ? AND default_watchlist_mapping.group_id = ?)
						OR (default_watchlist_mapping.type = ? AND default_watchlist_mapping.group_id = ?)`,
						"all",
						utils.WatchlistMappingUser, currentUser.UserID,
						utils.WatchlistMappingGroup, groupID,
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
		return c.Status(200).JSON(fiber.Map{
			"message": "Success",
			"username": username,
			"data":    watchlistResult,
		})
	})

}
