package model

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	IDPromotion int       `json:"id_promotion"`
	Promotion   Promotion `json:"promotion" gorm:"foreignKey:IDPromotion;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Active      bool      `json:"active"`
}

type Promotion struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Discount    float64 `json:"discount"`
	Active      bool    `json:"active"`
}
