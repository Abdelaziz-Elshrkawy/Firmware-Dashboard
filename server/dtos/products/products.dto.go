package productsDtos

type AddProductBody struct {
	Name string `json:"name"`
}

type UpdateProductBody struct {
	Id   *int   `json:"id"`
	Name string `json:"name"`
}

type DeleteProductBody struct {
	Id *int `json:"id"`
}
