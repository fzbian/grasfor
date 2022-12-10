package routes

import (
	"fmt"

	"github.com/fzbian/grasfor/utils"
	"github.com/gofiber/fiber/v2"
)

func AddPublication(c *fiber.Ctx) error {
	// Gets the parameters author, message and ip
	author := c.Query("author")
	if author == "" {
		author = "Anonymous"
	}
	message := c.Query("message")
	if message == "" {
		err := c.SendString("message empty")
		if err != nil {
			return err
		}
	}
	ip := c.IP()

	// Calls the AddPublication function to upload parameters to the database
	status := utils.AddPublication(author, message, ip)
	fmt.Println(status)
	return nil
}
