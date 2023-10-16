package models

import (
	"github.com/gofiber/fiber/v2"

	"myproject/database"
)

type Todo struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetTodos(c *fiber.Ctx) error {
	db := database.DBConn
	var todos []Todo
	db.Find(&todos)
	return c.JSON(&todos)
}

func GetTodoById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var todo Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Could Not Find Todo"})
	}
	return c.JSON(&todo)
}

func UpdateToDo(c *fiber.Ctx) error {

	type updatedToDo struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	id := c.Params("id")
	db := database.DBConn
	var todo Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "could not find todo list", "data": err})
	}
	var updateTodo updatedToDo
	err = c.BodyParser(&updateTodo)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "please Review your input", "data": err})
	}
	todo.Title = updateTodo.Title
	todo.Completed = updateTodo.Completed
	db.Save(&todo)
	return c.JSON(&todo)
}

func CreateToDo(c *fiber.Ctx) error {
	db := database.DBConn
	todo := new(Todo)
	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "check your input"})

	}

	err = db.Create(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "could not create to"})
	}

	return c.JSON(&todo)

}
