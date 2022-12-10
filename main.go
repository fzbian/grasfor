package main

import (
	"fmt"
	"log"

	"github.com/fzbian/grasfor/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// We instantiate fiber in a variable
	app := fiber.New()

	api := app.Group("/api")
	// The path "/" is called to return the health of the server
	api.Get("/health", routes.Health)
	// The path "/new-post" receives Query parameters, it doesn't hurt anything graphically
	api.Post("/new-post", routes.NewPost)
	api.Get("/posts", routes.GetPosts)

	// Render index.html
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendFile("views/index.html")
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})

	// We render the form to send the posts graphically
	app.Get("/new-post", func(c *fiber.Ctx) error {
		err := c.SendFile("views/new-post.html")
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})

	// We listen to the server on a port
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}
}
