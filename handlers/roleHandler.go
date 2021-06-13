package handlers

import (
	"github.com/DucHoangManh/go-admin/database"
	"github.com/DucHoangManh/go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role
	database.DB.Find(&roles)
	return c.JSON(fiber.Map{
		"message": "success",
		"payload": fiber.Map{
			"roles" : roles, 
		},
	})
}
func CreateRole(c *fiber.Ctx) error {
	var role models.Role
	err := c.BodyParser(&role)
	if err != nil {
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message": "can't parse request",
		})
	}
	database.DB.Create(&role)
	return c.JSON(fiber.Map{
		"message" : "success",
	})
}

func GetRole (c *fiber.Ctx) error {
	roleId, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	role:= models.Role{
		Id: uint(roleId),
	}
	database.DB.Find(&role)
	return c.JSON(fiber.Map{
		"message" : "success",
		"payload" : fiber.Map{
			"role" : role,
		},
	})
}

func UpdateRole(c *fiber.Ctx) error{
	roleId, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	role := models.Role{
		Id: uint(roleId),
	}
	err =database.DB.First(&role).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "role not found",
		})
	}
	err = c.BodyParser(&role)
	if err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when parsing role",
		})
	}
	database.DB.Model(&role).Updates(role)
	return c.JSON(fiber.Map{
		"message" : "success",
		"payload" : fiber.Map{
			"role" : role,
		},
	})
}

func DeleteRole(c *fiber.Ctx) error{
	roleId, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
		})
	}
	role := models.Role{
		Id: uint(roleId),
	}
	database.DB.Delete(&role)
	return nil
}