package routes

import (
	"github.com/gofiber/fiber/v2"
)

func MarketHandlers(route fiber.Router) {
	// Define market watchlist API function
	route.Get("/watchlist", func (C *fiber.Ctx) error  {

		return C.Status(200).JSON(fiber.Map{
			"message": "Success",
		})
	})

}