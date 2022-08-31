package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Retorna un string con la variable de entorno del puerto
func FiberPort() string {
	err := godotenv.Load()
	Errorhandle(err, "Error al importar el puerto", 2)

	return os.Getenv("MAINPORT")
}
