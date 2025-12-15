package deviceDtos

type AddDeviceBody struct {
	Serial      *string `json:"serial" validate:"required"`
	Product_Id  *uint   `json:"product_id" validate:"required"`
	Firmware_Id *uint   `json:"firmware_id" validate:"required"`
}

type UpdateDeviceBody struct {
	Id          *uint  `json:"id"  validate:"required"`
	Serial      string `json:"serial" validate:"required"`
	Product_Id  uint   `json:"product_id" validate:"required"`
	Firmware_Id uint   `json:"firmware_id" validate:"required"`
}

type DeleteDeviceBody struct {
	Id *uint `json:"id"  validate:"required"`
}
