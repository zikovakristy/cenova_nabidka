package user

import (
    "time"
)

type User struct {
    ID           uint       `gorm:"primaryKey" json:"id"`
    FirstName    string     `json:"first_name"`
    LastName     string     `json:"last_name"`
    Phone        string     `json:"phone"`
    Email        string     `gorm:"unique" json:"email"`
    City         string     `json:"city"`
    CreatedAt    time.Time  `json:"created_at"`
    UpdatedAt    time.Time  `json:"updated_at"`
    PriceOffers  []PriceOffer `gorm:"foreignKey:UserID" json:"price_offers"`
}
