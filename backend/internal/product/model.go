package product

type Product struct {
    ID          uint     `gorm:"primaryKey" json:"id"`
    Type        string   `json:"type"` // hw, sw, služba
    Program     []string `gorm:"type:text[]" json:"program"` // Multiple programs
    Default     bool     `json:"default"`
    ItemCode    string   `json:"item_code"`
    Name        string   `json:"name"`
    Description string   `json:"description"`
    URL         string   `json:"url"`
    ImageURL    string   `json:"image_url"`
    SalePrice   float64  `json:"sale_price"`
    RentPrice   float64  `json:"rent_price"`
    PurchasePrice float64 `json:"purchase_price"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    PriceOffers []PriceOfferProduct `gorm:"foreignKey:ProductID" json:"price_offers"`
}
