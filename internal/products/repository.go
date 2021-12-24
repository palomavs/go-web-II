package products

var products []Product
var lastID int

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Code         string  `json:"code"`
	Published    bool    `json:"published"`
	CreationDate string  `json:"creationDate"`
}

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
	return products, nil
}

func (r *repository) Store(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {
	newProduct := Product{id, name, color, price, stock, code, published, creationDate}
	products = append(products, newProduct)
	lastID++

	return newProduct, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}
