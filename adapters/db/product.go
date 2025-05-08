package db

import (
	"arquitetura-hexagonal-go/application"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt , err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// func (p *ProductDb) Create(product application.ProductInterface) (application.ProductInterface, error) {
	
// }

// func (p *ProductDb) Update(product application.ProductInterface) (application.ProductInterface, error) {
	
// }

// func (p *ProductDb) Delete(id string) error  {
	
// }
