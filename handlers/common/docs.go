package common

import (
	"os"
	"rpiSite/utils"

	"github.com/gofiber/fiber/v2"
)

func GetDocument(c *fiber.Ctx) error {
	fileName := c.Params("fileName")
	var fileContent []byte
	var err error
	switch fileName {
	case "tntn":
		fileContent, err = os.ReadFile("docs/TNTN_Algorithm.md")
	case "":
		fileContent, err = os.ReadFile("docs/index.md")
	default:
		err = os.ErrNotExist
	}
	if err != nil {
		return c.Render("err", fiber.Map{
			"Error": err.Error(),
		})
	}
	return c.Render("docs", fiber.Map{
		"content": utils.UnsafeStringConversion(fileContent),
	})
}
