package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *service) Store(name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}
	lastID++

	newProduct, err := s.repository.Store(lastID, name, color, price, stock, code, published, creationDate)
	if err != nil {
		return Product{}, err
	}

	return newProduct, nil
}
