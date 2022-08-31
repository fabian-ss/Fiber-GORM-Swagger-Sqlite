package routes

import (
	"github.com/fabian-ss/Fiber-GORM-Swagger-Sqlite/controllers"

	"github.com/gofiber/fiber/v2"
)

func NotasRoutes(app *fiber.App) {
	routes := app.Group("/apiNotas")
	routes.Get("/notas", controllers.GetNotas)
	routes.Post("/notas", controllers.CreateNota)
	routes.Get("/notas/:id", controllers.GetNota)
	routes.Get("/notasname/:id", controllers.GetNotaByUser)

}
