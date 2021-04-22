package handlers

import (
	"log"
	"rpiSite/config"
	"rpiSite/handlers/common"
	"rpiSite/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/template/html"
	"github.com/markbates/pkger"
	"github.com/ohler55/ojg/oj"
)

func renderEngine() *html.Engine {
	engine := html.NewFileSystem(pkger.Dir("/views"), ".html")
	return engine
}

func proxyHeader() (s string) {
	if config.IS_DEBUG != "true" {
		s = "X-Real-IP"
	}

	return s
}

func JsonEncoder(v interface{}) ([]byte, error) {
	return utils.S2b(oj.JSON(v, oj.Options{OmitNil: true})), nil
}

func Initalize() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: config.IS_DEBUG == "false",
		Views:                 renderEngine(),
		ProxyHeader:           proxyHeader(),
		Prefork:               true,
		JSONEncoder:           JsonEncoder,
	})

	app.Use(compress.New())
	if config.IS_DEBUG == "false" {
		app.Use(limiter.New(limiter.Config{Max: 150}))
	}

	app.Use("/", common.Index)

	log.Fatal(app.Listen(config.PORT))
}
