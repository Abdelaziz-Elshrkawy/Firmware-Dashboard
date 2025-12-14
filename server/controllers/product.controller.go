package controllers

import (
	productsDtos "firmware_server/dtos/products"
	"firmware_server/env/routes"
	"firmware_server/server"
	"firmware_server/services/productService"
	"firmware_server/utils"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func getProducts(c fiber.Ctx) error {
	var id *int

	var query = c.Query("id")

	if query != "" {
		value, err := strconv.Atoi(query)

		if err != nil {
			return utils.ResponeConstructor(c, fiber.StatusBadRequest, fiber.Map{
				"error": "invalid id query value",
			})
		}
		id = &value
	}

	products, sqlErr := productService.GetProducts(id)

	if sqlErr != nil {
		return utils.ResponeConstructor(c, fiber.StatusBadRequest, fiber.Map{
			"error": sqlErr,
		})
	}

	return utils.ResponeConstructor(c, fiber.StatusOK, fiber.Map{"products": products})
}

func addProduct(c fiber.Ctx) error {
	res, err := utils.ParseBody[productsDtos.AddProductBody](c)
	fmt.Println(res, err)
	if err != nil {
		return c.JSON(fiber.Map{
			"res": err.Error(),
		})
	}

	if res.Name == "" {
		return c.JSON(fiber.Map{
			"res": "Error product must have a name",
		})
	}

	sqlErr := productService.AddProduct(res.Name)
	if sqlErr != nil {
		return c.JSON(sqlErr)
	}

	return utils.ResponeConstructor(c, fiber.StatusCreated, fiber.Map{
		"message": "Product added successfully",
	})

}

func updateProduct(c fiber.Ctx) error {

	body, err := utils.ParseBody[productsDtos.UpdateProductBody](c)

	if err != nil {
		return utils.ResponeConstructor(c, fiber.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}

	if body.Id == nil {
		return utils.ResponeConstructor(c, fiber.StatusBadRequest, fiber.Map{
			"error": "product id must not be empty",
		})
	}

	if body.Name == "" {
		return utils.ResponeConstructor(c, fiber.StatusBadRequest, fiber.Map{
			"error": "product name must not be empty",
		})
	}

	if err = productService.UpdateProduct(*body.Id, body.Name); err != nil {
		return utils.ResponeConstructor(c, fiber.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}

	return utils.ResponeConstructor(c, fiber.StatusOK, fiber.Map{
		"message": "product updated",
	})
}

func deleteProduct(c fiber.Ctx) error {
	body, err := utils.ParseBody[productsDtos.DeleteProductBody](c)

	if err != nil {
		return utils.ResponeConstructor(c, fiber.StatusBadRequest, fiber.Map{
			"error": err.Error(),
		})
	}

	if body.Id == nil {
		return utils.ResponeConstructor(c, fiber.StatusBadRequest, fiber.Map{
			"error": "product id must not be empty",
		})
	}

	if err = productService.DeleteProduct(*body.Id); err != nil {
		return utils.ResponeConstructor(c, fiber.StatusInternalServerError, fiber.Map{
			"error": "internal server error",
		})
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
