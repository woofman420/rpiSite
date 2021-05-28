package user

import (
	"rpiSite/handlers/jwt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// AccountGet is the handler for `GET /account`
func AccountGet(c *fiber.Ctx) error {
	user, _ := jwt.User(c)

	// TODO, have a list of some sort.
	// Add a MOTD?
	var dayMessage string
	currentHour := time.Now().Hour()
	if currentHour < 6 {
		dayMessage = "Good night, maybe time to get to bed?"
	} else if currentHour > 6 && currentHour < 12 {
		dayMessage = "Good morning, had you got some coffee yet?"
	} else if currentHour > 12 && currentHour < 18 {
		dayMessage = "Good afternoon, make sure you don't forget to drink water."
	} else {
		dayMessage = "Good evening, it's lovely outside don't you think so?"
	}

	return c.Render("account", fiber.Map{
		"User":           user,
		"TimingOfTheDay": dayMessage,
	})
}
