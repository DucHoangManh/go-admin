package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)
func Paginate(db *gorm.DB, entity Entity, page int) fiber.Map{
	limit :=5
	offset := (page-1)*limit
	total := entity.Count(db)
	data := entity.Take(db, limit, offset) 
	return fiber.Map{
		"message": "success",
		"meta": fiber.Map{
			"page": page,
			"total": total,
			"last_page": float64(int(total)/limit)+1,
		},
		"payload": data,
	}
}