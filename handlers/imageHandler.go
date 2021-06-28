package handlers

import "github.com/gofiber/fiber/v2"

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil{
		return err
	}
	files := form.File["image"]
	var fileName string
	for _, file := range files{
		fileName = file.Filename
		if err := c.SaveFile(file,"./uploads/"+fileName); err != nil{
			return err
		}
	}
	return c.JSON(fiber.Map{
		"url": "https://127.0.0.1:8000/uploads/" + fileName,
	})
}