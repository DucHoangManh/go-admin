package routes

import (
	"github.com/DucHoangManh/go-admin/handlers"
	"github.com/gofiber/fiber/v2"
)
func Setup(app *fiber.App){
		app.Post("/api/register", handlers.Register)
		app.Post("/api/login", handlers.Login)
		

		//app.Use(middlewares.IsAuthenticated)
		app.Get("/api/user", handlers.User)
		app.Get("/api/logout", handlers.Logout)


		app.Get("/api/users", handlers.AllUsers)
		app.Post("api/users", handlers.CreateUser)
		app.Get("api/users/:id", handlers.GetUser)
		app.Put("api/users/:id", handlers.UpdateUser)
		app.Delete("api/users/:id", handlers.DeleteUser)

		app.Get("/api/roles", handlers.AllRoles)
		app.Post("api/roles", handlers.CreateRole)
		app.Get("api/roles/:id", handlers.GetRole)
		app.Put("api/roles/:id", handlers.UpdateRole)
		app.Delete("api/roles/:id", handlers.DeleteRole)

		app.Get("api/permissions", handlers.AllPermissions)


		app.Get("/api/products", handlers.AllProducts)
		app.Post("api/products", handlers.CreateProduct)
		app.Get("api/products/:id", handlers.GetProduct)
		app.Put("api/products/:id", handlers.UpdateProduct)
		app.Delete("api/products/:id", handlers.DeleteProduct)

		app.Post("api/upload", handlers.Upload)
}
