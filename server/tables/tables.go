package tables

type Product struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type Device struct {
	Id          uint `json:"id"`
	Serial      uint `json:"serial"`
	Product_Id  uint `json:"product_id"`
	Firmware_Id uint `json:"firmware_id"`
}

type Firmwares struct {
	Id         uint   `json:"id"`
	Version    string `json:"version"`
	Product_Id uint   `json:"product_id"`
}
