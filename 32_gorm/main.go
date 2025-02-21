package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	Name         string
	Price        float64
	CategoryId   uint
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

// has one
type SerialNumber struct {
	ID        uint `gorm:"primarykey"`
	Number    string
	ProductID uint
}

// belongs to
type Category struct {
	ID       uint `gorm:"primarykey"`
	Name     string
	Products []Product
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/database?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// db.Create(&Product{Name: "Hell", Price: 666.0})
	// product := Product{}
	// db.First(&product, 1)
	// println(product.Name)
	// secondProduct := Product{}
	// db.First(&secondProduct, "name = ?", "Potato")
	// println(secondProduct.Name)

	// products := []Product{}
	// db.Find(&products)
	// for _, product := range products {
	// 	println(product.Name)
	// }
	// // limit
	// db.Limit(2).Find(&products)

	// // offset por pagination
	// db.Limit(2).Offset(1).Find(&products)

	// // where
	// db.Where("price > ?", 100).Find(&products)

	// // like
	// likeProducts := []Product{}
	// db.Where("name LIKE ?", "%Hel%").Find(&likeProducts)
	// println("Like")
	// println(likeProducts[0].Name)

	// p := Product{}
	// db.First(&p, 1)
	// p.Name = "New Hell"
	// db.Save(&p)

	// potato := Product{}
	// db.First(&potato, 1)
	// db.Delete(&potato)

	// allProducts := []Product{}
	// db.Find(&allProducts)
	// println("All products")
	// for _, product := range allProducts {
	// 	println(product.Name)
	// }

	// category := Category{Name: "Electronics"}
	// db.Create(&category)
	serialNumber := SerialNumber{Number: "123456", ProductID: 1}
	db.Create(&serialNumber)
	// db.Create(&Product{Name: "Laptop", Price: 1000, CategoryId: category.ID})
	products_empty := []Product{}
	// db.Find(&products_empty)
	db.Preload("Category").Preload("SerialNumber").Find(&products_empty)
	println("All products / category")
	for _, product := range products_empty {
		println(product.Name)
		println(product.Category.Name)
		println(product.SerialNumber.Number)
	}

	categories := []Category{}
	db.Model(&Category{}).Preload("Products").Find(&categories)
	println("All categories")
	for _, category := range categories {
		println(category.Name)
		for _, product := range category.Products {
			println(product.Name)
		}
	}

}
