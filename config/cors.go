package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsHandler(app *fiber.App) {
	app.Use(cors.New())
}
