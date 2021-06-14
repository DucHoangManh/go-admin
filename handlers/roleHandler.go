package handlers

import (
	"github.com/DucHoangManh/go-admin/database"
	"github.com/DucHoangManh/go-admin/models"
	"github.com/gofiber/fiber/v2"
)


func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role
	database.DB.Preload("Permissions").Find(&roles)
	return c.JSON(fiber.Map{
		"message": "success",
		"payload": fiber.Map{
			"roles" : roles, 
		},
	})
}
func CreateRole(c *fiber.Ctx) error {
	var roleDTO fiber.Map
	err := c.BodyParser(&roleDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message": "can't parse request",
		})
	}
	perList := roleDTO["permissions"].([]interface{})
	permissions := make([]models.Permission, len(perList))
	for i, permissionId := range perList {
		id:= permissionId.(float64)
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}
	role := models.Role{
		Name: roleDTO["name"].(string),
		Permissions: permissions,
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
	database.DB.Preload("Permissions").Find(&role)
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
	var roleDTO fiber.Map
	err = c.BodyParser(&roleDTO)
	if err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when parsing role",
		})
	}
	perList := roleDTO["permissions"].([]interface{})
	permissions := make([]models.Permission, len(perList))
	for i, permissionId := range perList {
		id:= permissionId.(float64)
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}
	role := models.Role{
		Id: uint(roleId),
		Name: roleDTO["name"].(string),
		Permissions: permissions,
	}
	checkRole := models.Role{
		Id: uint(roleId),
	}
	err =database.DB.First(&checkRole).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "role not found",
		})
	}
	var deleteResult interface{}
	database.DB.Table("role_permissions").Where("role_id", roleId).Delete(&deleteResult)
	err = database.DB.Model(&role).Updates(role).Error
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
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