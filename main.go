package main

import (
	"github.com/DucHoangManh/go-admin/database"
	"github.com/DucHoangManh/go-admin/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
    AllowOrigins:     "*",
    AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))
	app.Use(logger.New())
	routes.Setup(app)
	app.Listen(":8000")
}
