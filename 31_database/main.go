package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(id, name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/database")
	if err != nil {
		panic(err)
	}
	insertProduct(db, NewProduct("1", "Product 1", 9.99))
	defer db.Close()
}

func insertProduct(db *sql.DB, product *Product) {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, error := stmt.Exec(product.ID, product.Name, product.Price)
	if error != nil {
		panic(error)
	}

}
