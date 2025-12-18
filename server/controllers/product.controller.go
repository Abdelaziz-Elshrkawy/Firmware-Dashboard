package controllers

import (
	productsDtos "firmware_server/dtos/products"
	"firmware_server/env/routes"
	"firmware_server/server"
	"firmware_server/services/productService"
	"firmware_server/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func getProducts(c fiber.Ctx) error {
	var id *int

	var query = c.Query("id")

	if query != "" {
		value, err := strconv.Atoi(query)

		if err != nil {
			return utils.BadRequestResponse(c, "invalid id query value")
		}
		id = &value
	}

	products, sqlErr := productService.GetProducts(id)
	println("get request")
	if sqlErr != nil {
		return utils.BadRequestResponse(c, sqlErr.Error())
	}

	return utils.ResponeConstructor(c, fiber.StatusOK, products)
}

func addProduct(c fiber.Ctx) error {
	body, err := utils.ParseBody[productsDtos.AddProductBody](c)

	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	if body.Name == "" {
		return utils.BadRequestResponse(c, "name is required")
	}

	sqlErr := productService.AddProduct(body.Name)

	if sqlErr != nil {
		return utils.BadRequestResponse(c, sqlErr.Error())
	}

	return utils.ResponeConstructor(c, fiber.StatusCreated, fiber.Map{
		"message": "Product added successfully",
	})

}

func updateProduct(c fiber.Ctx) error {
	body, err := utils.ParseBody[productsDtos.UpdateProductBody](c)

	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	if body.Id == nil {
		return utils.BadRequestResponse(c, "product id must not be empty")
	}

	if body.Name == "" {
		return utils.BadRequestResponse(c, "product name must not be empty")
	}

	if err = productService.UpdateProduct(*body.Id, body.Name); err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	return utils.ResponeConstructor(c, fiber.StatusOK, fiber.Map{
		"message": "product updated",
	})
}

func deleteProduct(c fiber.Ctx) error {
	body, err := utils.ParseBody[productsDtos.DeleteProductBody](c)

	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	if body.Id == nil {
		return utils.BadRequestResponse(c, "id is required")
	}

	if err = productService.DeleteProduct(*body.Id); err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	return utils.ResponeConstructor(c, fiber.StatusOK, fiber.Map{
		"message": "deleted",
	})
}

func productRoute() {
	var ProductGroup = server.App.Group(routes.Product)
	ProductGroup.Get("", getProducts)
	ProductGroup.Post("", addProduct)
	ProductGroup.Put("", updateProduct)
	ProductGroup.Delete("", deleteProduct)
}
