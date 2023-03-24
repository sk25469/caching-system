package controllers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sk25469/momoney-backend-assignment/middleware"
	"github.com/sk25469/momoney-backend-assignment/utils"
)

func GetTodos(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	todoId, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Error in parsing Todo id: [%v]", err)
		return ctx.Status(500).SendString(err.Error())
	}
	log.Printf("Successfully parsed id: [%v]", todoId)

	response, err := middleware.InitMiddleWare(todoId, utils.Todo)
	if err != nil {
		log.Printf("Error occured in middleware: [%v]", err)
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.Status(200).JSON(response)
}

func GetPosts(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	postId, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Error in parsing Post id: [%v]", err)
		return ctx.Status(500).SendString(err.Error())
	}
	log.Printf("Successfully parsed id: [%v]", postId)
	response, err := middleware.InitMiddleWare(postId, utils.Post)
	if err != nil {
		log.Printf("Error occured in middleware: [%v]", err)
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.Status(200).JSON(response)
}
