package main

import (
	"log"

	"github.com/fzbian/grasfor/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// We instantiate fiber in a variable
	app := fiber.New()

	// The path "/" is called to return the health of the server
	app.Get("/", routes.Health)
	// The path "/addPublication" receives Query parameters, it doesn't hurt anything graphically
	app.Get("/addPublication", routes.AddPublication)

	// We listen to the server on a port
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}
}
