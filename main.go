package main

import (
	db2 "arquitetura-hexagonal-go/adapters/db"
	"arquitetura-hexagonal-go/application"
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", ".files/sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	productDb := db2.NewProductDb(db)
	productService := application.NewProductService(productDb)

	productService.Create("Product 1", 10.0)
	productService.Create("Product 2", 20.0)
	productService.Create("Product 3", 30.0)
}