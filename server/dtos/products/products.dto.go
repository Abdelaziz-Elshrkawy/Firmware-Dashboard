package productsDtos

type AddProductBody struct {
	Name string `json:"name" validate:"required"`
}

type UpdateProductBody struct {
	Id   *int   `json:"id" validate:"required,max=10"`
	Name string `json:"name" validate:"required,max=10"`
}

type DeleteProductBody struct {
	Id *int `json:"id" validate:"required"`
}
