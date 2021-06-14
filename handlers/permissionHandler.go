package handlers

import (
	"github.com/DucHoangManh/go-admin/database"
	"github.com/DucHoangManh/go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func AllPermissions(c *fiber.Ctx) error {
	var pers []models.Permission
	database.DB.Find(&pers)
	return c.JSON(fiber.Map{
		"message": "success",
		"payload": fiber.Map{
			"roles" : pers, 
		},
	})
}