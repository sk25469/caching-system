package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sk25469/momoney-backend-assignment/middleware"
	"github.com/sk25469/momoney-backend-assignment/utils"
)

// Toggles caching for todos
var cachingEnabledForTodos = true

// Toggles caching for posts
var cachingEnabledForPosts = true

// Fetches a todo with given id
func GetTodos(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	todoId, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Error in parsing Todo id: [%v]", err)
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	log.Printf("Successfully parsed id: [%v]", todoId)

	response, err := middleware.InitMiddleWare(todoId, utils.Todo, cachingEnabledForTodos)
	if err != nil {
		log.Printf("Error occured in middleware: [%v]", err)
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

// Toggles caching for todos
func ToggleCachingForTodos(ctx *fiber.Ctx) error {
	flagParam := ctx.Params("flag")
	if err := utils.CheckTypeAndValues(flagParam, "string", []string{"true", "false"}); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	log.Printf("Successfully parsed flag for todos: [%v]", flagParam)
	if flagParam == "true" {
		cachingEnabledForTodos = true
		return ctx.Status(http.StatusOK).JSON("Caching enabled for todos")
	} else {
		cachingEnabledForTodos = false
		return ctx.Status(http.StatusOK).JSON("Caching disabled for todos")
	}
}

// Fetches post for a given id
func GetPosts(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	postId, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Error in parsing Post id: [%v]", err)
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	log.Printf("Successfully parsed id: [%v]", postId)
	response, err := middleware.InitMiddleWare(postId, utils.Post, cachingEnabledForPosts)
	if err != nil {
		log.Printf("Error occured in middleware: [%v]", err)
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

// Toggles caching for posts
func ToggleCachingForPosts(ctx *fiber.Ctx) error {
	flagParam := ctx.Params("flag")
	if err := utils.CheckTypeAndValues(flagParam, "string", []string{"true", "false"}); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	log.Printf("Successfully parsed flag for todos: [%v]", flagParam)

	if flagParam == "true" {
		cachingEnabledForPosts = true
		return ctx.Status(http.StatusOK).JSON("Caching enabled for posts")
	} else {
		cachingEnabledForPosts = false
		return ctx.Status(http.StatusOK).JSON("Caching disabled for posts")

	}

}

// Toggles caching for both posts and todos
func ToggleCachingForAll(ctx *fiber.Ctx) error {
	flagParam := ctx.Params("flag")
	if err := utils.CheckTypeAndValues(flagParam, "string", []string{"true", "false"}); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	log.Printf("Successfully parsed flag for all routes: [%v]", flagParam)
	if flagParam == "true" {
		cachingEnabledForPosts = true
		cachingEnabledForTodos = true
		return ctx.Status(http.StatusOK).JSON("Caching enabled for all routes")
	} else {
		cachingEnabledForPosts = false
		cachingEnabledForTodos = false
		return ctx.Status(http.StatusOK).JSON("Caching disabled for all routes")
	}

}
