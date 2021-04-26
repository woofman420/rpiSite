package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"rpiSite/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/ohler55/ojg/oj"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

var JSONParser = oj.Parser{
	Reuse: true,
}

func CallbackHelperUSWPost(c *fiber.Ctx) error {
	refer, code, state, clientID, clientSecret := c.FormValue("refer"), c.Query("code"), c.Query("state"), c.FormValue("clientID"), c.FormValue("clientSecret")

	if refer == "" || code == "" || clientID == "" || clientSecret == "" {
		return c.Render("err", fiber.Map{
			"Error": "Wrong inputs.",
		})
	}
	var referURL string
	switch refer {
	case "localhost":
		referURL = "http://localhost:3000"
	default:
		referURL = "https://userstyles.world"
	}

	url := referURL + "/oauth/access_token"
	url += "?client_id=" + clientID
	url += "&client_secret=" + clientSecret
	url += "&code=" + code
	if state != "" {
		url += "&state=" + state
	}

	req, err := http.Post(url, "application/x-www-form-urlencoded", nil)

	if err != nil {
		log.Println("Error fetching URL:", err)
		return c.Render("err", fiber.Map{
			"Error": "Couldn't fetch URL",
		})
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Error reading body:", err)
		return c.Render("err", fiber.Map{
			"Error": "Couldn't fetch body",
		})
	}

	if req.StatusCode != 200 {
		log.Println(req.StatusCode, utils.B2s(body))
		return c.Render("err", fiber.Map{
			"Error": "Didn't return 200 code",
		})
	}

	var returnJSON AccessToken
	err = oj.Unmarshal(body, &returnJSON)
	if err != nil {
		log.Println(err)
		return c.Render("err", fiber.Map{
			"Error": "Couldn't encode body to JSON",
		})
	}

	return c.JSON(returnJSON)
}
