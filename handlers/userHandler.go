package handlers

import (
	"strconv"

	"github.com/DucHoangManh/go-admin/database"
	"github.com/DucHoangManh/go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func AllUsers(c *fiber.Ctx) error {
	limit :=2
	page, _:= strconv.Atoi(c.Query("page","1"))
	offset := (page-1)*limit
	var total int64
	var users []models.User
	database.DB.Model(&users).Count(&total)
	database.DB.Preload("Role").Offset(offset).Limit(limit).Find(&users)
	return c.JSON(fiber.Map{
		"message": "success",
		"meta": fiber.Map{
			"page": page,
			"total": total,
			"last_page": float64(int(total)/limit),
		},
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