package cli

import (
	"arquitetura-hexagonal-go/application"
	"errors"
	"fmt"
)

func Run(service application.ProductServiceInterface, action string, id string, name string, price float64) (string, error) {
	switch action {
		case "create":
			return create(service, name, price)
		case "enable":
			return enable(service, id)
		case "disable":
			return disable(service, id)
		default:
			return "", errors.New("invalid action")
	}
}

func create(service application.ProductServiceInterface, name string, price float64) (string, error) {
	product, err := service.Create(name, price)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("product (%s) was created", product.GetName()), nil
}

func enable(service application.ProductServiceInterface, id string) (string, error) {
	service.Enable(id)
	return fmt.Sprintf("product (%s) was enabled", id), nil
}

func disable(service application.ProductServiceInterface, id string) (string, error) {
	service.Disable(id)
	return fmt.Sprintf("product (%s) was disabled", id), nil
}