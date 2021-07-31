package handlers

import (
	"fmt"
	"github.com/gofiber/fiber"
	"time"
)

type UserId struct {
	Id int `json:"id"`
}

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

type UserWithoutTime struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
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

	err = c.JSON(user)
	if err != nil {
		fmt.Println("error with JSON response in CreateUserHandler", err)
	}

}

func GetUsersHandler(c *fiber.Ctx) {

	users, err := StoreToDB.GetUsers()
	if err != nil {
		fmt.Println("error with getUsers", err)
	}

	err = c.JSON(users)
	if err != nil {
		fmt.Println("error with JSON response in GetUsersHandler", err)
	}

}

func DeleteUserHandler(c *fiber.Ctx) {
	user := UserId{}

	u := new(UserId)
	if err := c.BodyParser(u); err != nil {
		c.SendStatus(500)
		c.SendString("Body Parser error")
	}

	user.Id = u.Id

	err := StoreToDB.DeleteUser(&user)
	if err != nil {
		fmt.Println("error deleting from database", err)
	}

	err = c.JSON(user)
	if err != nil {
		fmt.Println("error with JSON response in CreateUserHandler", err)
	}

}

func UpdateUserHandler(c *fiber.Ctx) {

	user := UserWithoutTime{}

	u := new(UserWithoutTime)
	if err := c.BodyParser(u); err != nil {
		c.SendStatus(500)
		c.SendString("Body Parser error")
	}
	user.Id = u.Id
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Age = u.Age

	err := StoreToDB.UpdateUser(&user)
	if err != nil {
		fmt.Println("error storing to database", err)
	}

	err = c.JSON(user)
	if err != nil {
		fmt.Println("error with JSON response in UpdateUserHandler", err)
	}
}
