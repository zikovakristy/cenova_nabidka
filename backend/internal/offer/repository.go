package offer

import (
    "gorm.io/gorm"
)

type Repository interface {
    CreateOffer(offer *PriceOffer) error
    GetOfferByID(id uint) (*PriceOffer, error)
    GetAllOffers() ([]PriceOffer, error)
    GetOffersByUserID(userID uint) ([]PriceOffer, error)
    UpdateOffer(offer *PriceOffer) error
    DeleteOffer(id uint) error
}

type repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
    return &repository{db}
}

func (r *repository) CreateOffer(offer *PriceOffer) error {
    return r.db.Create(offer).Error
}

func (r *repository) GetOfferByID(id uint) (*PriceOffer, error) {
    var offer PriceOffer
    if err := r.db.Preload("User").Preload("Products.Product").First(&offer, id).Error; err != nil {
        return nil, err
    }
    return &offer, nil
}

func (r *repository) GetAllOffers() ([]PriceOffer, error) {
    var offers []PriceOffer
    if err := r.db.Preload("User").Preload("Products.Product").Find(&offers).Error; err != nil {
        return nil, err
    }
    return offers, nil
}

func (r *repository) GetOffersByUserID(userID uint) ([]PriceOffer, error) {
    var offers []PriceOffer
    if err := r.db.Preload("User").Preload("Products.Product").Where("user_id = ?", userID).Find(&offers).Error; err != nil {
        return nil, err
    }
    return offers, nil
}

func (r *repository) UpdateOffer(offer *PriceOffer) error {
    return r.db.Save(offer).Error
}

func (r *repository) DeleteOffer(id uint) error {
    return r.db.Delete(&PriceOffer{}, id).Error
}
