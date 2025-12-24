package tables

type Product struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type Device struct {
	Id          uint   `json:"id"`
	Serial      string `json:"serial"`
	Product_Id  uint   `json:"product_id"`
	Firmware_Id uint   `json:"firmware_id"`
	Api_Key     string `json:"api_key"`
}

type Firmwares struct {
	Id         uint   `json:"id"`
	Version    string `json:"version"`
	Product_Id uint   `json:"product_id"`
}

type Users struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
