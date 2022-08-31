package routes

import (
	"github.com/fabian-ss/Fiber-GORM-Swagger-Sqlite/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	route := app.Group("/apiUsers")
	route.Post("/user", controllers.CreateUser)
	route.Get("/user", controllers.GetUsers)
	route.Get("/user/:id", controllers.GetUser)
	route.Put("/user/:id", controllers.UpdateUser)
	route.Delete("/user/:id", controllers.DeleteUser)
}
