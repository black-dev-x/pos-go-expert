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
	product := NewProduct("1", "Product 2", 9.99)
	insertProduct(db, product)
	product, error := selectProduct(db, product.ID)
	if error != nil {
		panic(error)
	}
	println(product.Name)
	defer db.Close()
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	product := &Product{}
	err := db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return product, nil
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

func updateProduct(db *sql.DB, product *Product) {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, error := stmt.Exec(product.Name, product.Price, product.ID)
	if error != nil {
		panic(error)
	}
}
