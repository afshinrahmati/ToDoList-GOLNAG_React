package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"myproject/database"
	"myproject/models"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Afshinm")
}

func setupRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
	app.Get("/todos/:id", models.GetTodoById)
	app.Post("/todos", models.GetTodos)
}

func initDatabse() {
	var err error
	dsn := "host=127.0.0.1 user=postgres password=afvsa dbname=goTodo port=5432"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Faild to connect to database!")
	}
	fmt.Println("Database Connected!")
	database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Print("Migrate Db!")
}
func main() {
	app := fiber.New()
	initDatabse()
	app.Get("/", helloWorld)
	setupRoutes(app)
	app.Listen(":3200")
}
