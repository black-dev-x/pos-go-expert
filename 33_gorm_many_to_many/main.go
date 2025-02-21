package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:product_categories;"`
	gorm.Model
}

// belongs to
type Category struct {
	ID       uint `gorm:"primarykey"`
	Name     string
	Products []Product `gorm:"many2many:product_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/database?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// kitchen := Category{Name: "Kitchen"}
	// db.Create(&kitchen)
	// electronics := Category{Name: "Electronics"}
	// db.Create(&electronics)

	// product := Product{Name: "Potato Killer", Price: 1.0, Categories: []Category{kitchen, electronics}}
	// db.Create(&product)

	products := []Product{}
	db.Preload("Categories").Find(&products)
	for _, product := range products {
		println(product.Name)
		for _, category := range product.Categories {
			println(" - ", category.Name)
		}
	}
	tx := db.Begin()
	category := Category{}
	// pessimistic lock
	tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&category, "name = ?", "Kitchen")
	category.Name = "Hell Kitchen"
	tx.Debug().Save(&category)
	tx.Commit()

}
