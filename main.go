package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
	"go_be_db_app/handlers"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "postgres"
)

func setupRoutes(app *fiber.App) {
	app.Post("/api/v1/create_user", handlers.CreateUserHandler)
	app.Get("/api/v1/get_users", handlers.GetUsersHandler)
	app.Post("/api/v1/update_user", handlers.UpdateUserHandler)
	app.Post("/api/v1/delete_user", handlers.DeleteUserHandler)
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

	handlers.InitStore(&handlers.DBstore{Db: db})

	app := fiber.New()
	setupRoutes(app)
	err = app.Listen(4000)
	if err != nil {
		fmt.Println("error with listening:", err)
	}
}
