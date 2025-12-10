package env

var DBuser = "root"
var DBpassword = ""
var DBname = "products_firmware"

var MqttHost = "localhost"
var MqttPort = 1883

var MqttTopics = map[string]string{
	"update": "Update",
}
