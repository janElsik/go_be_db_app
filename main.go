package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"go_be_db_app/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Post("/api/v1/create_user", handlers.UserHandler)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	err := app.Listen(4000)
	if err != nil {
		fmt.Println("error with listening:", err)
	}
}
