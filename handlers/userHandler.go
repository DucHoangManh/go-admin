package handlers

import (
	"strconv"

	"github.com/DucHoangManh/go-admin/database"
	"github.com/DucHoangManh/go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func AllUsers(c *fiber.Ctx) error {
	page, _:= strconv.Atoi(c.Query("page","1"))
	return c.JSON(models.Paginate(database.DB,&models.User{},page))
}
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message": "can't parse request",
		})
	}
	//check if email already in use
	var aux_user models.User
	database.DB.Where("Email = ?", user.Email).First(&aux_user)
	if aux_user.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already in use",
		})
	}
	user.SetPassword("123456")
	database.DB.Create(&user)
	return c.JSON(fiber.Map{
		"message" : "success",
	})
}

func GetUser (c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	user:= models.User{
		Id: uint(userId),
	}
	database.DB.Find(&user)
	return c.JSON(fiber.Map{
		"message" : "success",
		"payload" : fiber.Map{
			"user" : user,
		},
	})
}

func UpdateUser(c *fiber.Ctx) error{
	userId, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	user := models.User{
		Id: uint(userId),
	}
	err =database.DB.First(&user).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	err = c.BodyParser(&user)
	if err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when parsing user",
		})
	}
	err = database.DB.Model(&user).Updates(user).Error
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "wrong person info",
		})
	}
	return c.JSON(fiber.Map{
		"message" : "success",
		"payload" : fiber.Map{
			"user" : user,
		},
	})
}

func DeleteUser(c *fiber.Ctx) error{
	userId, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
		})
	}
	user := models.User{
		Id: uint(userId),
	}
	database.DB.Delete(&user)
	return nil
}