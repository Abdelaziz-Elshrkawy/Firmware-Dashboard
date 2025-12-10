package server

import "github.com/gofiber/fiber/v3"

// before using this variable the Init function need to called first

var App *fiber.App

func Init() {
	App = fiber.New()

}
