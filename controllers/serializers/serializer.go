package serializers

import (
	"github.com/fabian-ss/Fiber-GORM-Swagger-Sqlite/models"
)

// Serializer, no modelo
type UserSerializer struct {
	ID        uint   `json:"id"`
	FirtsName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func CreateUserResponse(userModel models.User) UserSerializer {
	return UserSerializer{
		ID:        userModel.ID,
		FirtsName: userModel.FirtsName,
		LastName:  userModel.LastName,
	}
}

type NotaSerializer struct {
	ID          uint        `json:"id"`
	User        models.User `json:"user"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
}

func CreateNotaResponse(notaModel models.Nota, user models.User) NotaSerializer {
	return NotaSerializer{
		ID:          notaModel.ID,
		User:        user,
		Title:       notaModel.Title,
		Description: notaModel.Description,
	}
}
