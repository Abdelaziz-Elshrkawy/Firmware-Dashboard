package main

import (
	"firmware_server/appMqtt"
	"firmware_server/controllers"
	"firmware_server/database"
	"firmware_server/server"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/fiber/v3"
)

func main() {

	// initializing database
	database.Connect()

	server.Init()

	appMqtt.InitMqtt()

	server.App.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello")
	})

	appMqtt.Client.Subscribe("test/topic", 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message on topic: %s: %s %d\n", msg.Topic(), string(msg.Payload()), msg.Qos())
	})

	controllers.RegisterControllers()

	if err := server.App.Listen(":3000"); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
