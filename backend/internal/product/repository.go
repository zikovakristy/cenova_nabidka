package product

import (
    "gorm.io/gorm"
)

type Repository interface {
    CreateProduct(product *Product) error
    GetProductByID(id uint) (*Product, error)
    GetAllProducts() ([]Product, error)
    UpdateProduct(product *Product) error
    DeleteProduct(id uint) error
}

type repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
    return &repository{db}
}

func (r *repository) CreateProduct(product *Product) error {
    return r.db.Create(product).Error
}

func (r *repository) GetProductByID(id uint) (*Product, error) {
    var product Product
    if err := r.db.Preload("PriceOffers").First(&product, id).Error; err != nil {
        return nil, err
    }
    return &product, nil
}

func (r *repository) GetAllProducts() ([]Product, error) {
    var products []Product
    if err := r.db.Preload("PriceOffers").Find(&products).Error; err != nil {
        return nil, err
    }
    return products, nil
}

func (r *repository) UpdateProduct(product *Product) error {
    return r.db.Save(product).Error
}

func (r *repository) DeleteProduct(id uint) error {
    return r.db.Delete(&Product{}, id).Error
}
