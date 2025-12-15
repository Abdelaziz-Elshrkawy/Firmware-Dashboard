package firmwareDtos

type AddFirmwareBody struct {
	Version    *string `json:"version" validate:"required"`
	Product_Id *uint   `json:"product_id" validate:"required"`
}

type UpdateFirmwareBody struct {
	Id         *uint   `json:"id" validate:"required"`
	Version    *string `json:"version" validate:"required"`
	Product_Id *uint   `json:"product_id" validate:"required"`
}

type DeleteFirmwareBody struct {
	Id *uint `json:"id" validate:"required"`
}
