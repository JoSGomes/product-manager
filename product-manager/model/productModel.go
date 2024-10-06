package model

type Product struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Price       float64    `json:"price"`
	Description string     `json:"description"`
	Promotion   *Promotion `json:"promotion"`
	Active      bool       `json:"active"`
}

type Promotion struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Discount    float64 `json:"discount"`
	Active      bool    `json:"active"`
}
