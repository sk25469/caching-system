package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sk25469/momoney-backend-assignment/controllers"
)

// This function registers the 2 GET requests routes for [posts] and [todos]
// and 2 GET requests for toggling caching in todos and posts
func RegisterRoutes(port string) {
	app := fiber.New()

	app.Get("/posts/:id", controllers.GetPosts)
	app.Get("/todos/:id", controllers.GetTodos)
	app.Get("/caching/posts=:flag", controllers.ToggleCachingForPosts)
	app.Get("/caching/todos=:flag", controllers.ToggleCachingForTodos)
	app.Get("/caching=:flag", controllers.ToggleCachingForAll)

	log.Printf("🚀Server started at post: [%v]", port)
	log.Fatal(app.Listen(":" + port))
}
