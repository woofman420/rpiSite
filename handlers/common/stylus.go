package common

import (
	"rpiSite/utils"
	"sync"
	"time"

	"github.com/cespare/xxhash"
	"github.com/gofiber/fiber/v2"
)

var (
	// Limiter variables
	mux        = &sync.RWMutex{}
	expiration = 15 * time.Second
	// Minus the expiration time, so on startup of the server we don't have to wait expiration time.
	timestamp = time.Now()
)

func StylusEvalGet(c *fiber.Ctx) error {
	return c.Render("stylus", fiber.Map{})
}

// Ensure that this function can only be executed every 10 seconds
func StylusEvalPost(c *fiber.Ctx) error {
	mux.RLock()
	defer mux.RUnlock()

	css, site := c.FormValue("css"), c.FormValue("url")

	if css == "" || site == "" {
		return c.Render("stylus", fiber.Map{
			"Error": "Missing css or site",
			"URL":   site,
			"CSS":   css,
		})
	}

	if time.Now().Before(timestamp) {
		return c.Render("stylus", fiber.Map{
			"Error": "Please wait 15 seconds before executing this function again.",
			"URL":   site,
			"CSS":   css,
		})
	}

	timestamp = time.Now().Add(expiration)

	// Execute the stylus compiler
	// Hash the name of css and site to SHA1.
	// This will be used as the filename of the compiled css.

	newHash := xxhash.New()
	_, err := newHash.Write([]byte(css + site))
	if err != nil {
		return err
	}
	fileName := utils.HashForFileName(newHash.Sum(nil))

	err = utils.TakeScreenshot(css, site, fileName)
	if err != nil {
		return c.Render("stylus", fiber.Map{
			"Error": err.Error(),
			"URL":   site,
			"CSS":   css,
		})
	}

	return c.Render("show_stylus_image", fiber.Map{
		"URL":      site,
		"fileName": fileName,
	})
}
