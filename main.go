package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
	"go_be_db_app/helpers"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "postgres"
)

func setupRoutes(app *fiber.App) {
	app.Post("/api/v1/create_user", helpers.CreateUserHandler)
	app.Get("/api/v1/get_users", helpers.GetUsersHandler)
	app.Post("/api/v1/update_user", helpers.UpdateUserHandler)
	app.Post("/api/v1/delete_user", helpers.DeleteUserHandler)
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	helpers.InitStore(&helpers.DBstore{Db: db})

	app := fiber.New()
	setupRoutes(app)
	err = app.Listen(4000)
	if err != nil {
		fmt.Println("error with listening:", err)
	}
}
