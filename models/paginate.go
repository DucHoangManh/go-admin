package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)
func Paginate(db *gorm.DB, page int) fiber.Map{
	limit :=5
	offset := (page-1)*limit
	var total int64
	var products []Product
	db.Model(&products).Count(&total)
	db.Offset(offset).Limit(limit).Find(&products)
	return fiber.Map{
		"message": "success",
		"meta": fiber.Map{
			"page": page,
			"total": total,
			"last_page": float64(int(total)/limit),
		},
		"payload": products,
	}
}