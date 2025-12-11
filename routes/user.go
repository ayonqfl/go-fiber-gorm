package routes

import (
	"github.com/ayonqfl/go-fiber-gorm/database"
	"github.com/ayonqfl/go-fiber-gorm/models"
	"github.com/gofiber/fiber/v2"
)

type UserSerializer struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(user models.User) UserSerializer {
	return UserSerializer{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName}
}

func CreateUser(C *fiber.Ctx) error {
	var user models.User

	if err := C.BodyParser(&user); err != nil {
		return C.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	return C.Status(200).JSON(responseUser)
}
