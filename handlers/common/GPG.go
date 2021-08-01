package common

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

var GPGData []byte

func GPG(c *fiber.Ctx) error {
	if GPGData != nil {
		c.Response().SetBody(GPGData)
		return nil
	}
	gpgData, err := os.ReadFile("/public.asc")
	if err != nil {
		log.Println("Error reading public.asc:", err)
		return c.Render("err", fiber.Map{
			"Error": "Couldn't read GPG file.",
		})
	}
	GPGData = gpgData
	c.Response().SetBody(GPGData)
	return nil
}
