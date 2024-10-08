package offer

import (
    "time"

    "cenova_nabidka/backend/internal/user"
)

type PriceOffer struct {
    ID            uint       `gorm:"primaryKey" json:"id"`
    UserID        uint       `json:"user_id"`
    User          user.User  `json:"user"`
    OfferName     string     `json:"offer_name"`
    FirstName     string     `json:"first_name"`
    LastName      string     `json:"last_name"`
    Phone         string     `json:"phone"`
    Email         string     `json:"email"`
    City          string     `json:"city"`
    CreatedAt     time.Time  `json:"created_at"`
    UpdatedAt     time.Time  `json:"updated_at"`
    Products      []PriceOfferProduct `gorm:"foreignKey:PriceOfferID" json:"products"`
    TotalPrice    float64    `json:"total_price"`
    DiscountSW    float64    `json:"discount_sw"`
    DiscountHW    float64    `json:"discount_hw"`
    DiscountService float64 `json:"discount_service"`
}

type PriceOfferProduct struct {
    ID            uint    `gorm:"primaryKey" json:"id"`
    PriceOfferID  uint    `json:"price_offer_id"`
    ProductID     uint    `json:"product_id"`
    Product       product.Product `json:"product"`
    Quantity      int     `json:"quantity"`
    UnitPrice     float64 `json:"unit_price"`
    TotalPrice    float64 `json:"total_price"`
    IsZeroPrice   bool    `json:"is_zero_price"` // Indicates if the product price is zero
}
