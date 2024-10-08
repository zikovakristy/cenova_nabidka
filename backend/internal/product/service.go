package product

type Service interface {
    CreateProduct(product *Product) error
    GetProductByID(id uint) (*Product, error)
    GetAllProducts() ([]Product, error)
    UpdateProduct(product *Product) error
    DeleteProduct(id uint) error
}

type service struct {
    repo Repository
}

func NewService(repo Repository) Service {
    return &service{repo}
}

func (s *service) CreateProduct(product *Product) error {
    return s.repo.CreateProduct(product)
}

func (s *service) GetProductByID(id uint) (*Product, error) {
    return s.repo.GetProductByID(id)
}

func (s *service) GetAllProducts() ([]Product, error) {
    return s.repo.GetAllProducts()
}

func (s *service) UpdateProduct(product *Product) error {
    return s.repo.UpdateProduct(product)
}

func (s *service) DeleteProduct(id uint) error {
    return s.repo.DeleteProduct(id)
}
