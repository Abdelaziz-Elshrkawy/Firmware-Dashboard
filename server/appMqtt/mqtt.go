package appMqtt

import (
	"firmware_server/env"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var Client mqtt.Client

func InitMqtt() error {
	options := mqtt.NewClientOptions()

	options.AddBroker(fmt.Sprintf("mqtt://%s:%d", env.MqttHost, env.MqttPort))
	options.ClientID = "bedo_firmware_server"
	options.AutoReconnect = true

	Client = mqtt.NewClient(options)

	if token := Client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	fmt.Println("Mqtt Connected")

	return nil
}
