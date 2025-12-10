package productService

import (
	"firmware_server/database"
	"firmware_server/tables"
	"fmt"
)

func GetProducts() ([]tables.Product, error) {
	res, err := database.DB.Query("select * from products;")

	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return nil, err
	}

	defer res.Close()

	var products []tables.Product

	for res.Next() {
		var product tables.Product
		if err := res.Scan(&product.Id, &product.Name); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
