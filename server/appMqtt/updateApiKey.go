package appMqtt

import (
	"firmware_server/env/topics"
)

func UpdateApikey(msg string) {
	Client.Publish(topics.UpdateApikey, 0, false, msg)
}
