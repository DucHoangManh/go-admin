package handlers

import (
	"github.com/DucHoangManh/go-admin/database"
	"github.com/DucHoangManh/go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func AllUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(fiber.Map{
		"message": "success",
		"payload": users,
	})
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
	database.DB.Model(&user).Updates(user)
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