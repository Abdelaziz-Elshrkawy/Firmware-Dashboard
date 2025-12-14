package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

// before using this variable the Init function need to called first

var App *fiber.App

func Init() {
	App = fiber.New()
	App.Use(cors.New(cors.Config{}))
}
