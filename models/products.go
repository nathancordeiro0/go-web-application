package models

import (
	db "go-web-application/database"
)

type Product struct {
	Id, Amount        int
	Name, Description string
	Price             float64
}

func SearchAllProducts() []Product {
	db := db.ConnectDatabase()

	selectAllProducts, err := db.Query("select * from products order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &amount)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, amount int) {
	db := db.ConnectDatabase()

	insertData, err := db.Prepare("insert into products(name, description, price, amount) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, price, amount)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDatabase()

	deleteData, err := db.Prepare("delete from products where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteData.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConnectDatabase()

	productData, err := db.Query("select * from products where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	updatedProduct := Product{}

	for productData.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = productData.Scan(&id, &name, &description, &price, &amount)

		if err != nil {
			panic(err.Error())
		}

		updatedProduct.Id = id
		updatedProduct.Name = name
		updatedProduct.Description = description
		updatedProduct.Price = price
		updatedProduct.Amount = amount

	}
	defer db.Close()
	return updatedProduct
}

func UpdateProduct(id int, name, description string, price float64, amount int) {
	db := db.ConnectDatabase()

	updatedProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, amount=$4 where id=$5 ")

	if err != nil {
		panic(err.Error())
	}

	updatedProduct.Exec(name, description, price, amount, id)
	defer db.Close()

}
