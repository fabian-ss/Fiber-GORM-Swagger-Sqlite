package controllers

import (
	"errors"

	"github.com/fabian-ss/Fiber-GORM-Swagger-Sqlite/controllers/serializers"
	"github.com/fabian-ss/Fiber-GORM-Swagger-Sqlite/database"
	"github.com/fabian-ss/Fiber-GORM-Swagger-Sqlite/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Crear una nota
// @Description Endpoint: /apiNotas/notas/
// @Tags Notas
// @Param nota_attrs body models.NotaResponse true "Nota attributes"
// @Router /apiNotas/notas/ [post]
func CreateNota(c *fiber.Ctx) error {

	var nota models.Nota

	if err := c.BodyParser(&nota); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error al mandar la estructura")
	}

	var user models.User

	if err := findUser(nota.UserRefer, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error al validar el usuario")
	}

	database.Database.Db.Create(&nota)

	reponseNota := serializers.CreateNotaResponse(nota, user)

	return c.Status(fiber.StatusOK).JSON(reponseNota)

}

// @Summary Ver todas mis notas
// @Description Endpoint: /apiNotas/notas/
// @Tags Notas
// @Router /apiNotas/notas/ [get]
func GetNotas(c *fiber.Ctx) error {
	notas := []models.Nota{}

	database.Database.Db.Find(&notas)

	responseNotas := []serializers.NotaSerializer{}

	for _, nota := range notas {
		var user models.User

		database.Database.Db.Find(&user, "id=?", nota.UserRefer)
		responseNota := serializers.CreateNotaResponse(nota, user)
		responseNotas = append(responseNotas, responseNota)
	}
	return c.Status(200).JSON(responseNotas)

}

func findNotas(id int, nota *models.Nota) error {

	database.Database.Db.Find(&nota, "id=?", id)

	if nota.ID == 0 {
		return errors.New("usuario no existe")
	}

	return nil
}

// @Summary Ver un usuario por ID
// @Description Endpoint: /apiNotas/notas/id
// @Tags Notas
// @ID get-int-by-int
// @Param id path int true "Account ID"
// @Router /apiNotas/notas/{id} [get]
func GetNota(c *fiber.Ctx) error {

	var nota models.Nota

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("error al comparar el id")
	}

	if err := findNotas(id, &nota); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error al encontrar la nota")
	}

	var user models.User
	database.Database.Db.First(&user, nota.UserRefer)

	responseNota := serializers.CreateNotaResponse(nota, user)

	return c.Status(fiber.StatusOK).JSON(responseNota)

}

// @Summary Ver las notas de un usuario
// @Description Endpoint: /apiNotas/notasname/id  Descripcion: Busca las notas asociadas al id del usuario
// @Tags Notas
// @Param id path int true "Account ID"
// @Router /apiNotas/notasname/{id} [get]
func GetNotaByUser(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("error en la estructura")
	}
	var user models.User
	notas := []models.Nota{}

	database.Database.Db.Find(&notas)

	responseNotas := []serializers.NotaSerializer{}

	for _, nota := range notas {

		database.Database.Db.Find(&user, "id=?", nota.UserRefer)

		if nota.UserRefer == id {
			responseNota := serializers.CreateNotaResponse(nota, user)
			responseNotas = append(responseNotas, responseNota)
		}

	}
	return c.Status(200).JSON(responseNotas)

}
