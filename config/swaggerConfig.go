package config

import (
	docs "github.com/fabian-ss/Fiber-GORM-Swagger-Sqlite/docs"
)

func SwaggerConfig() {
	docs.SwaggerInfo.Title = "API con Fiber, Swagger y GORM"
	docs.SwaggerInfo.Description = "CRUD de usuarios y notas usando GORM"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
