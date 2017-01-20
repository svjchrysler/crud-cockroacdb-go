package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"log"
)


type Product struct {

}


func conexion() *gorm.DB{
	db, err := gorm.Open("postgres", "host=localhost user=root dbname=warehouse sslmode=disable port=26257")
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func createTables() {
	conexion:= conexion()

}

func main() {
	createTables()


}
