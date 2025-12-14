package productService

import (
	"database/sql"
	"firmware_server/database"
	"firmware_server/tables"
	"firmware_server/tables/tablesName"
	"fmt"
)

func GetProducts(id *int) ([]tables.Product, error) {
	var res *sql.Rows
	var err error

	if id != nil {
		res, err = database.DB.Query("select * from ? where id=?", tablesName.Products, id)
	} else {
		res, err = database.DB.Query("select * from ?", tablesName.Products)
	}

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
	query := "insert into ? (name) VALUES (?)"

	_, err := database.DB.Exec(query, tablesName.Products, name)

	if err != nil {
		return err
	}

	return nil
}

func UpdateProduct(id int, name string) error {
	query := "update ? set name = ? where id = ?"

	res, err := database.DB.Exec(query, tablesName.Products, name, id)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func DeleteProduct(id int) error {
	query := "delete from ? where id=?"

	_, err := database.DB.Exec(query, tablesName.Products, id)

	if err != nil {
		return err
	}
	return nil
}
