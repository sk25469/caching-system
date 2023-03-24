package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sk25469/momoney-backend-assignment/controllers"
)

// This function registers the 2 get requests routes for [posts] and [todos]
func RegisterRoutes() {
	app := fiber.New()

	app.Get("/posts/:id", controllers.GetPosts)
	app.Get("/todos/:id", controllers.GetTodos)

	log.Fatal(app.Listen(":3000"))
}
