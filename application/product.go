package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuidv4 "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error

	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(product ProductInterface) (ProductInterface, error)
	Enable(id string) (ProductInterface, error)
	Disable(id string) (ProductInterface, error)
}

type ProductRepositoryInterface interface {
	Get(id string) (ProductInterface, error)
	Create(product ProductInterface) (ProductInterface, error)
	Update(product ProductInterface) (ProductInterface, error)
	Delete(id string) error
}

const (
	DISABLED = "DISABLED"
	ENABLED  = "ENABLED"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

func NewProduct() *Product {
	return &Product{
		ID: uuidv4.NewV4().String(),
		Status: DISABLED,
	}
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != DISABLED && p.Status != ENABLED {
		return false, errors.New("invalid product status, must be ENABLED or DISABLED")
	}

	if p.Price < 0 {
		return false, errors.New("product price must be greater than or equal to 0")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("product price must be greater than 0 to enable")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("product price must be 0 to disable")
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
