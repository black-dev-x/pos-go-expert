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
	// limit
	db.Limit(2).Find(&products)

	// offset por pagination
	db.Limit(2).Offset(1).Find(&products)

	// where
	db.Where("price > ?", 100).Find(&products)

	// like
	likeProducts := []Product{}
	db.Where("name LIKE ?", "%Hel%").Find(&likeProducts)
	println("Like")
	println(likeProducts[0].Name)

	p := Product{}
	db.First(&p, 1)
	p.Name = "New Hell"
	db.Save(&p)

	potato := Product{}
	db.First(&potato, 1)
	db.Delete(&potato)

}
