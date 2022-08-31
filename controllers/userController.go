package controllers

import (
	"errors"
	"log"

	"github.com/fabian-ss/Fiber-GORM-Swagger-Sqlite/controllers/serializers"
	"github.com/fabian-ss/Fiber-GORM-Swagger-Sqlite/database"
	"github.com/fabian-ss/Fiber-GORM-Swagger-Sqlite/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Crear un usuario
// @Description Endpoint: /apiUsers/user/
// @Tags User
// @Param nota_attrs body models.UserResponse true "User attributes"
// @Router /apiUsers/user/ [post]
func CreateUser(c *fiber.Ctx) error {

	var user models.User

	if err := c.BodyParser(&user); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	database.Database.Db.Create(&user)

	reponseUser := serializers.CreateUserResponse(user)

	return c.Status(200).JSON(reponseUser)

}

// @Summary Ver todos los usuarios
// @Description Endpoint: /apiUsers/user/
// @Tags User
// @Router /apiUsers/user/ [get]
func GetUsers(c *fiber.Ctx) error {

	users := []models.User{}

	database.Database.Db.Find(&users)

	responseUsers := []serializers.UserSerializer{}

	for _, user := range users {
		responseUser := serializers.CreateUserResponse(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)

}

func findUser(id int, user *models.User) error {

	database.Database.Db.Find(&user, "id=?", id)

	if user.ID == 0 {
		return errors.New("usuario no existe")
	}

	return nil
}

// @Summary Ver un usuario por ID
// @Description Endpoint: /apiUsers/user/id
// @Tags User
// @Param id path int true "Account ID"
// @Router /apiUsers/user/{id} [get]
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		log.Println(err)
		return c.Status(400).JSON("Usuario no encontradoDDDDDD")
	}

	var user models.User

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := serializers.CreateUserResponse(user)

	return c.Status(200).JSON(responseUser)

}

// @Summary Actualizar un usuario por ID
// @Description Endopoint: /apiUsers/user/id
// @Tags User
// @Param id path int true "Account ID"
// @Param nota_attrs body models.UserResponse true "User attributes"
// @Router /apiUsers/user/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		log.Println(err)
		return c.Status(400).JSON("Usuario no encontradoDDDDDD")
	}

	var user models.User

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	updateUser := models.UserResponse{}

	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	user.FirtsName = updateUser.FirtsName
	user.LastName = updateUser.LastName

	database.Database.Db.Save(&user)

	reponseUser := serializers.CreateUserResponse(user)

	return c.Status(fiber.StatusOK).JSON(reponseUser)

}

// @Summary Elimiar un usuario por ID
// @Description Endpoint: /apiUsers/getoneuser/id
// @Tags User
// @Param id path int true "Account ID"
// @Router /apiUsers/user/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON("Usuario no encontrado")
	}

	var user models.User
	if err := findUser(id, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("Usuario eliminado")

}
