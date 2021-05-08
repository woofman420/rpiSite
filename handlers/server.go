package handlers

import (
	"log"
	"rpiSite/config"
	"rpiSite/handlers/api"
	"rpiSite/handlers/common"
	"rpiSite/handlers/jwt"
	"rpiSite/handlers/monitor"
	"rpiSite/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

func JSONEncoder(v interface{}) ([]byte, error) {
	return utils.S2b(oj.JSON(v, oj.Options{OmitNil: true, HTMLUnsafe: false})), nil
}

func Initalize() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: config.IS_DEBUG == "false",
		Views:                 renderEngine(),
		ProxyHeader:           proxyHeader(),
		Prefork:               true,
		JSONEncoder:           JSONEncoder,
	})

	if config.IS_DEBUG == "true" {
		app.Use(logger.New())
	}

	app.Use(compress.New())
	if config.IS_DEBUG == "false" {
		app.Use(limiter.New(limiter.Config{Max: 150}))
	}
	app.Use(jwt.New())

	app.Get("/", common.Index)
	app.Group("/monitor", jwt.Protected, monitor.ProxyMonitor)

	app.Get("/callback_helper/:type?", api.CallbackGet)
	app.Post("/usw/access_token", api.CallbackHelperUSWPost)
	app.Get("/register", common.RegisterGet)
	app.Post("/register", common.RegisterPost)
	app.Get("/login", common.LoginGet)
	app.Post("/login", common.LoginPost)

	app.Get("/login", common.LoginGet)
	app.Post("/login", common.LoginPost)

	if config.IS_DEBUG == "true" {
		app.Static("/", "/static")
	}
	app.Use("/", filesystem.New(filesystem.Config{
		MaxAge: int(time.Hour) * 2,
		Root:   pkger.Dir("/static"),
	}))

	app.Use(common.NotFound)

	log.Fatal(app.Listen(config.PORT))
}
