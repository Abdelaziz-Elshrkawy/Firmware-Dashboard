package main

import (
	"bufio"
	"firmware_server/appMqtt"
	"firmware_server/controllers"
	"firmware_server/database"
	"firmware_server/server"
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/fiber/v3"
)

func main() {

	// dtos.InitValidator()

	// initializing database
	if err := database.Connect(); err != nil {
		log.Fatal("Error Connecting to Database:\n", err)
		bufio.NewScanner(os.Stdin)
		return
	}

	if err := appMqtt.InitMqtt(); err != nil {
		log.Fatal("Error Connecting to MQTT broker:\n", err)
		bufio.NewScanner(os.Stdin)
		return
	}

	server.Init()

	controllers.RegisterControllers()

	server.App.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Bedo Firmware Remote Update Api")
	})

	appMqtt.Client.Subscribe("test/topic", 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message on topic: %s: %s\n", msg.Topic(), string(msg.Payload()))
	})

	if err := server.App.Listen(":3000"); err != nil {
		fmt.Println("Server failed to start:", err)
	}

}
