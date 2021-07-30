package handlers

import (
	"fmt"
	"github.com/gofiber/fiber"
)

func UserHandler(c *fiber.Ctx) {
	fmt.Println(c.BaseURL())
	c.Status(200).SendString("all ok")
}
