package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Post("/todos", models.CreateToDo)
	app.Put("/todos/:id", models.UpdateToDo)
}

func initDatabse() {
	var err error
	dsn := "host=127.0.0.1 user=postgres password=afvsa9899 dbname=goTodo port=5432"
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
	app.Use(cors.New())
	app.Get("/", helloWorld)
	setupRoutes(app)
	app.Listen(":3200")
}
