package routes

import (
	"fmt"
	"net/http"

	"github.com/fzbian/grasfor/database"
	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	var (
		messageStatusDatabase string
		messageStatusServer   string
		messageOnWeb          string
	)

	// Call the PingDB function to make a request to the database and check its status.
	result, err := database.PingDB()
	if err != nil {
		return err
	}

	if result == true {
		messageStatusDatabase = fmt.Sprintf("Ok!")
	} else {
		messageStatusDatabase = fmt.Sprintf("Failed!")
	}

	// A request is made to the server (it calls itself) to check the status of the server and its response.
	resp, err := http.Get(c.BaseURL())
	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		messageStatusServer = fmt.Sprintf("Ok!")
	} else {
		messageStatusServer = fmt.Sprintf("Failed!")
	}

	// Prints request statuses
	messageOnWeb = fmt.Sprintf("Health\nServer is %s\nDatabase is %s", messageStatusServer, messageStatusDatabase)
	return c.SendString(messageOnWeb)
}
