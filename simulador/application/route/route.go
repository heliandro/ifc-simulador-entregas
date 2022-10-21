package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Route struct {
	ID string
	ClientID string
	Positions []Position
}

type Position struct {
	Lat float64
	Long float64
}

func teste(r *Route) error {
	if r.ID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "ID is required")
	}

	if r.ClientID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "ClientID is required")
	}

	if len(r.Positions) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Positions is required")
	}

	return nil
	// file, err := os.OpenFile("routes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}

func Routes(app *fiber.App) {
	app.Use(logger.New())

	app.Get("/health", GetHealth)
}

func GetHealth(c *fiber.Ctx) error {
	
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "API Simulador is up and running",
	})
}