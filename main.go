package main

import (
	"github.com/fabian-ss/Fiber-GORM-Swagger-Sqlite/config"
	"github.com/fabian-ss/Fiber-GORM-Swagger-Sqlite/database"

	"log"

	"github.com/fabian-ss/Fiber-GORM-Swagger-Sqlite/routes"

	"github.com/gofiber/fiber/v2"
)

// @contact.name   API Support - My Github
// @contact.url    https://github.com/fabian-ss
// @contact.email  randommusicd@gmail.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	database.ConnectDB()

	// Creación de servidor
	app := fiber.New()

	// Configuración de cors y swagger
	config.CorsHandler(app)
	config.SwaggerConfig()

	// Rutas
	routes.SwaggerRoute(app)
	routes.UserRoutes(app)
	routes.NotasRoutes(app)
	routes.NotFoundRoute(app)

	// Levantar el servidor
	log.Fatal(app.Listen(config.FiberPort()))
}
