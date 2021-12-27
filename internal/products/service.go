package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (Product, error)
	Update(id int, name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (Product, error)
	UpdateNameAndPrice(id int, name string, price float64) (Product, error)
	HardDelete(id int) ([]Product, error)
	Delete(id int) ([]Product, error)
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

func (s *service) Store(name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}
	lastID++

	newProduct, err := s.repository.Store(lastID, name, color, price, stock, code, published, creationDate, active)
	if err != nil {
		return Product{}, err
	}

	return newProduct, nil
}

func (s *service) Update(id int, name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (Product, error) {
	return s.repository.Update(id, name, color, price, stock, code, published, creationDate, active)
}

func (s *service) HardDelete(id int) ([]Product, error) {
	return s.repository.HardDelete(id)
}

func (s *service) Delete(id int) ([]Product, error) {
	return s.repository.Delete(id)
}

func (s *service) UpdateNameAndPrice(id int, name string, price float64) (Product, error) {
	return s.repository.UpdateNameAndPrice(id, name, price)
}
