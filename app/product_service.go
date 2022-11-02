package app

type ProductService struct {
	Persistence ProductPersistenceInterface
}

// Get(id string) (ProductInterface, error)
// Create(name string, price float64) (ProductInterface, error)
// Enable(product ProductInterface) (ProductInterface, error)
// Disable(product ProductInterface) (ProductInterface, error)

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
