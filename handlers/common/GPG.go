package common

import (
	"io"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/markbates/pkger"
)

var GPGData []byte

func GPG(c *fiber.Ctx) error {
	if GPGData != nil {
		c.Response().SetBody(GPGData)
		return nil
	}
	f, err := pkger.Open("/public.asc")
	if err != nil {
		log.Println("Error reading public.asc:", err)
		return c.Render("err", fiber.Map{
			"Error": "Couldn't read GPG file.",
		})
	}
	gpgData, err := io.ReadAll(f)
	if err != nil {
		log.Println("Error reading GPG file:", err)
		return c.Render("err", fiber.Map{
			"Error": "Couldn't read GPG file.",
		})
	}
	GPGData = gpgData
	c.Response().SetBody(GPGData)
	return nil
}
