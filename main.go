package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"log"
)


type Product struct {
	ID uint `gorm:"primary_key"`
	Name string `gorm:"type:varchar(100)"`
	Price float32
}


func conexion() *gorm.DB{
	db, err := gorm.Open("postgres", "host=localhost user=root dbname=warehouse sslmode=disable port=26257")
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func createTables() {
	db := conexion()

	if !db.HasTable(&Product{}) {
		db.CreateTable(&Product{})
	}

	defer db.Close()
}

func createProduct() {

	db:= conexion()

	product := Product{Name:"coca cola", Price:10.5}

	db.Create(&product)

	defer db.Close()

}

func main() {
	createTables()
	createProduct()
}
