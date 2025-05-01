package application

type ProductService struct {
	ProductRepository ProductRepositoryInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	return s.ProductRepository.Get(id)
}

func (s *ProductService) Create(product ProductInterface) (ProductInterface, error) {
	return s.ProductRepository.Create(product)
}

func (s *ProductService) Enable(id string) (ProductInterface, error) {
	product, err := s.ProductRepository.Get(id)
	if err != nil {
		return nil, err
	}

	err = product.Enable()
	if err != nil {
		return nil, err
	}

	return s.ProductRepository.Update(product)
}