package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/database"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	db.Create(&Product{Name: "Potato", Price: 666.0})
	product := Product{}
	db.First(&product, 1)
	println(product.Name)
	secondProduct := Product{}
	db.First(&secondProduct, "name = ?", "Potato")
	println(secondProduct.Name)

	products := []Product{}
	db.Find(&products)
	for _, product := range products {
		println(product.Name)
	}

}
