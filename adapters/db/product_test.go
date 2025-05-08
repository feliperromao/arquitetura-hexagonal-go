package db_test

import (
	"arquitetura-hexagonal-go/adapters/db"
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	insertProduct(Db, "1", "Product 1", 10.0, "active")
	insertProduct(Db, "2", "Product 2", 20.0, "inactive")
	insertProduct(Db, "3", "Product 3", 30.0, "active")
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
	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("1")
	require.Nil(t, err)
	require.Equal(t, "1", product.GetID())
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, "active", product.GetStatus())
}