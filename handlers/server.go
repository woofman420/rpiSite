package handlers

import (
	"html/template"
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
	md "github.com/russross/blackfriday/v2"
)

var ext = md.CommonExtensions | md.AutoHeadingIDs

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

	engine.AddFunc("mdSafe", func(s string) template.HTML {
		gen := md.Run([]byte(s), md.WithExtensions(ext))
		return template.HTML(gen)
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

// JSONEncoder is a simple wrapper for a fast JSON Encoder.
func JSONEncoder(v interface{}) ([]byte, error) {
	return oj.Marshal(v, &oj.Options{OmitNil: true, HTMLUnsafe: false})
}

// JSONDecoder is a simple wrapper for a fast JSON Decoder.
func JSONDecoder(data []byte, v interface{}) error {
	return oj.Unmarshal(data, v)
}

// Initialize the server code and listen for it.
func Initialize() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: config.IsProduction,
		Views:                 renderEngine(),
		ProxyHeader:           proxyHeader(),
		JSONEncoder:           JSONEncoder,
		JSONDecoder:           JSONDecoder,
	})

	if config.IsDebug {
		app.Use(logger.New())
	}

	app.Use(middleware.AddHeaders)
	app.Use(compress.New())
	if config.IsProduction {
		app.Use(limiter.New(limiter.Config{Max: 150}))
	}

	app.Get("/", common.Index)
	app.Get("/.gpg", common.GPG)
	app.Get("/gusted.gpg", common.GPG)

	app.Get("/callback_helper", api.CallbackGet)
	app.Get("/stylus", common.StylusEvalGet)
	app.Post("/stylus", common.StylusEvalPost)
	app.Get("/docs/:fileName?", common.GetDocument)

	year2021 := app.Group("/2021")
	year2021.Get("/eindgesprek", common.EindgesprekGet)
	year2021.Get("/portfolio", common.PortfolioGet)

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
