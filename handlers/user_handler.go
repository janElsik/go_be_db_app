package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber"
	"time"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

type FullUser struct {
	Id           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Age          int       `json:"age"`
	CreationDate time.Time `json:"creation_date"`
}

func CreateUserHandler(c *fiber.Ctx) {

	user := User{}

	u := new(User)
	if err := c.BodyParser(u); err != nil {
		c.SendStatus(500)
		c.SendString("Body Parser error")
	}

	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Age = u.Age

	err := StoreToDB.CreateUser(&user)
	if err != nil {
		fmt.Println("error storing to database", err)
	}

	c.Status(200).SendString("all ok")

}

func GetUsersHandler(c *fiber.Ctx) {

	users, err := StoreToDB.GetUsers()
	if err != nil {
		fmt.Println("error with getUsers", err)
	}
	c.Send(json.Marshal(users))

}
