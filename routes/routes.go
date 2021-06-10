package routes

import (
	"github.com/DucHoangManh/go-admin/handlers"
	"github.com/DucHoangManh/go-admin/middlewares"
	"github.com/gofiber/fiber/v2"
)
func Setup(app *fiber.App){
		app.Post("/api/register", handlers.Register)
		app.Post("/api/login", handlers.Login)

		app.Use(middlewares.IsAuthenticated)
		app.Get("/api/user", handlers.User)
		app.Get("/api/logout", handlers.Logout)
		app.Get("/other", handlers.Others)
}
