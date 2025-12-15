package productService

import (
	"firmware_server/database"
	"firmware_server/tables"
	"firmware_server/tables/tablesName"
	"fmt"
)

func GetProducts(id *int) ([]tables.Product, error) {
	var query string

	if id != nil {
		query = fmt.Sprintf("select * from %s where id=%d", tablesName.Products, *id)
	} else {
		query = fmt.Sprintf("select * from %s", tablesName.Products)
	}

	res, err := database.DB.Query(query)

	if err != nil {
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

func AddProduct(name string) error {

	query := fmt.Sprintf("insert into %s (name) VALUES (?)", tablesName.Products)

	_, err := database.DB.Exec(query, name)

	if err != nil {
		return err
	}

	return nil
}

func UpdateProduct(id int, name string) error {
	query := fmt.Sprintf("update %s set name = ? where id = ?", tablesName.Products)

	res, err := database.DB.Exec(query, name, id)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func DeleteProduct(id int) error {
	query := fmt.Sprintf("delete from %s where id=?", tablesName.Products)

	_, err := database.DB.Exec(query, id)

	if err != nil {
		return err
	}
	return nil
}
