package handlers

import (
	"strconv"

	"github.com/DucHoangManh/go-admin/database"
	"github.com/DucHoangManh/go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func AllProducts(c *fiber.Ctx) error {
	page, _:= strconv.Atoi(c.Query("page","1"))
	return c.JSON(models.Paginate(database.DB, &models.Product{}, page))
}


func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	err := c.BodyParser(&product)
	if err != nil {
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message": "can't parse request",
		})
	}
	database.DB.Create(&product)
	return c.JSON(fiber.Map{
		"message" : "success",
	})
}

func GetProduct (c *fiber.Ctx) error {
	productId, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	product:= models.Product{
		Id: uint(productId),
	}
	database.DB.Find(&product)
	return c.JSON(fiber.Map{
		"message" : "success",
		"payload" : fiber.Map{
			"product" : product,
		},
	})
}

func UpdateProduct(c *fiber.Ctx) error{
	productId, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"message" : "invalid id",
		})
	}
	product := models.Product{
		Id: uint(productId),
	}
	err =database.DB.First(&product).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "product not found",
		})
	}
	err = c.BodyParser(&product)
	if err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error when parsing product",
		})
	}
	err = database.DB.Model(&product).Updates(product).Error
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "wrong product info",
		})
	}
	return c.JSON(fiber.Map{
		"message" : "success",
		"payload" : fiber.Map{
			"product" : product,
		},
	})
}

func DeleteProduct(c *fiber.Ctx) error{
	productId, err := c.ParamsInt("id")
	if err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid id",
		})
	}
	product := models.Product{
		Id: uint(productId),
	}
	database.DB.Delete(&product)
	return nil
}