package handlers

import (
	"strconv"
	"time"

	"github.com/DucHoangManh/go-admin/database"
	"github.com/DucHoangManh/go-admin/models"
	"github.com/DucHoangManh/go-admin/util"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	requestData := make(map[string]string)
	if err := c.BodyParser(&requestData); err != nil {
		return err
	}
	//check if password and confirm password are the same
	if requestData["password"] != requestData["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Confirm password does not match",
		})
	}
	//check if email already in use
	var aux_user models.User
	database.DB.Where("Email = ?", requestData["email"]).First(&aux_user)
	if aux_user.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already in use",
		})
	}
	user := models.User{
		FirstName: requestData["first_name"],
		LastName:  requestData["last_name"],
		Email:     requestData["email"],
		RoleId: 1,
	}
	user.SetPassword(requestData["password"])
	database.DB.Create(&user)
	return c.JSON(fiber.Map{
		"message": "success",
		"payload": fiber.Map{
			"user": user,
		},
	})
}

func Login(c *fiber.Ctx) error {
	requestData := make(map[string]string)
	if err := c.BodyParser(&requestData); err != nil {
		return err
	}
	var user models.User
	database.DB.Where("Email = ?", requestData["email"]).First(&user)
	if user.Id == 0{
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	err := user.ComparePassword(requestData["password"])
	if err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "wrong password",
		})
	}

	token, err := util.Generate(strconv.Itoa(int(user.Id)))
	if err != nil{
		return c.SendString(err.Error())
	}
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
		"payload": fiber.Map{
			"user": user,
			"token": token,
		},
	})
}
func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)
	

	var user models.User
	database.DB.Where("Id = ?", id).First(&user)
	return c.JSON(user)
}

func Logout (c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}