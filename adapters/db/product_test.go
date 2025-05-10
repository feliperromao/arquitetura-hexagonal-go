package db_test

import (
	"arquitetura-hexagonal-go/adapters/db"
	"arquitetura-hexagonal-go/application"
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	insertProduct(Db, "1", "Product 1", 10.0, application.ENABLED)
	insertProduct(Db, "2", "Product 2", 20.0, application.DISABLED)
	insertProduct(Db, "3", "Product 3", 30.0, application.ENABLED)
}

func createTable(db *sql.DB) {
	// _, _ = Db.Exec("CREATE TABLE products (id TEXT PRIMARY KEY, name TEXT, price REAL, status TEXT)")
	table:= `CREATE TABLE IF NOT EXISTS products (
				"id" string,
				"name" string,
				"price" float,
				"status" string
			)`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func insertProduct(db *sql.DB, id string, name string, price float64, status string) {
	_, err := db.Exec("INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)", id, name, price, status)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	repository := db.NewProductRepository(Db)

	product, err := repository.Get("1")
	require.Nil(t, err)
	require.Equal(t, "1", product.GetID())
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, application.ENABLED, product.GetStatus())
}

func TestProductDb_Create(t *testing.T) {
	setUp()
	defer Db.Close()
	repository := db.NewProductRepository(Db)

	product := application.NewProduct()
	product.Name = "Product 4"
	product.Price = 45.97
	product.Status = application.ENABLED

	createdProduct, err := repository.Create(product)
	require.Nil(t, err)
	require.Equal(t, product.GetID(), createdProduct.GetID())
	require.Equal(t, product.GetName(), createdProduct.GetName())
	require.Equal(t, product.GetPrice(), createdProduct.GetPrice())
	require.Equal(t, product.GetStatus(), createdProduct.GetStatus())
}

func TestProductDb_Update(t *testing.T) {
	setUp()
	defer Db.Close()
	repository := db.NewProductRepository(Db)

	product := application.NewProduct()
	product.ID = "1"
	product.Name = "Product 1 Updated"
	product.Price = 15.0
	product.Status = application.DISABLED

	updatedProduct, err := repository.Update(product)
	require.Nil(t, err)
	require.Equal(t, product.GetID(), updatedProduct.GetID())
	require.Equal(t, product.GetName(), updatedProduct.GetName())
	require.Equal(t, product.GetPrice(), updatedProduct.GetPrice())
	require.Equal(t, product.GetStatus(), updatedProduct.GetStatus())
}