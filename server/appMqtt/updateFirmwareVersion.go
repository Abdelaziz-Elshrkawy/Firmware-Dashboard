package appMqtt

import "firmware_server/env/topics"

func UpdateFirmwareVersion(msg string) {
	Client.Publish(topics.UpdateFirmwareVersion, 0, false, msg)
}
