package application

type ProductService struct {
	ProductRepository ProductRepositoryInterface
}

func NewProductService(repository ProductRepositoryInterface) *ProductService {
	return &ProductService{
		ProductRepository: repository,
	}
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	return s.ProductRepository.Get(id)
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price
	_, err := product.IsValid()
	if err != nil {
		return nil, err
	}
	result, err := s.ProductRepository.Create(product)
	if err != nil {
		return nil, err
	}
	return result, nil
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


func (s *ProductService) Disable(id string) (ProductInterface, error) {
	product, err := s.ProductRepository.Get(id)
	if err != nil {
		return nil, err
	}

	err = product.Disable()
	if err != nil {
		return nil, err
	}

	return s.ProductRepository.Update(product)
}