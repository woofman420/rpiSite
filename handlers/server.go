package handlers

import (
	"log"
	"net/http"
	"time"

	"rpiSite/config"
	"rpiSite/handlers/api"
	"rpiSite/handlers/common"
	"rpiSite/handlers/middleware"
	"rpiSite/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/ohler55/ojg/oj"
)

func renderEngine() *html.Engine {
	cssHash := utils.GetCSSHash()
	engine := html.NewFileSystem(http.Dir("./views"), ".html")

	engine.AddFunc("cssHash", func() string {
		if config.IsDebug {
			return utils.GetCSSHash()
		} else {
			return cssHash
		}
	})
	engine.Reload(config.IsDebug)
	return engine
}

func proxyHeader() (s string) {
	if config.IsProduction {
		s = "X-Real-IP"
	}

	return s
}

// JSONEncoder is a simple wrapper for a fast JSON ENcoder.
func JSONEncoder(v interface{}) ([]byte, error) {
	return utils.UnsafeByteConversion(oj.JSON(v, oj.Options{OmitNil: true, HTMLUnsafe: false})), nil
}

// Initialize the server code and listen for it.
func Initialize() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: config.IsProduction,
		Views:                 renderEngine(),
		ProxyHeader:           proxyHeader(),
		JSONEncoder:           JSONEncoder,
	})

	if config.IsDebug {
		app.Use(logger.New())
	}

	app.Use(middleware.AddHeaders)
	app.Use(compress.New())
	if config.IsProduction {
		app.Use(limiter.New(limiter.Config{Max: 150}))
	}
	//	app.Use(jwt.New())

	app.Get("/", common.Index)
	app.Get("/.gpg", common.GPG)
	app.Get("/gusted.gpg", common.GPG)

	app.Get("/callback_helper", api.CallbackGet)
	//app.Get("/register", common.RegisterGet)
	//app.Post("/register", common.RegisterPost)
	//app.Get("/login", common.LoginGet)
	//app.Post("/login", common.LoginPost)
	app.Get("/stylus", common.StylusEvalGet)
	app.Post("/stylus", common.StylusEvalPost)

	year2021 := app.Group("/2021")
	year2021.Get("/eindgesprek", common.EindgesprekGet)

	//app.Get("/account", jwt.Protected, user.AccountGet)

	if config.IsDebug {
		app.Static("/", "/static")
	}
	app.Use("/", filesystem.New(filesystem.Config{
		MaxAge: int(time.Hour) * 2,
		Root:   http.Dir("./static"),
	}))

	app.Use(common.NotFound)

	log.Fatal(app.Listen(config.Port))
}
