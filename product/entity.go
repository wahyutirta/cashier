package product

import "time"

type Product struct {
	ID          int
	ProductCode string
	ProductName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ProductUnit string
	Price       []Price
}

type Price struct {
	ID        int
	ProductID int
	BuyPrice  int
	SellPrice int
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// input entity

type IDParamsInput struct {
	ID int `uri:"id" binding:"required"`
}

type JsonProductInput struct {
	ProductName string `json:"product_name" binding:"required"`
	ProductCode string `json:"product_code" binding:"required"`
	ProductUnit string `json:"product_unit" binding:"required"`
}

type JsonPriceInput struct {
	ProductID int `json:"product_id" binding:"required"`
	BuyPrice  int `json:"buy_price" binding:"required,numeric"`
	SellPrice int `json:"sell_price" binding:"required,numeric"`
	IsPrimary int `json:"is_primary" binding:"omitempty"`
}

// formatter

type ProductFormatter struct {
	ID          int    `json:"id"`
	ProductCode string `json:"product_code"`
	ProductName string `json:"product_name"`
	// CreatedAt   time.Time        `json:"created_at"`
	// UpdatedAt   time.Time        `json:"update_at"`
	ProductUnit string           `json:"product_unit"`
	Prices      []PriceFormatter `json:"product_prices"`
}

type PriceFormatter struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	BuyPrice  int `json:"buy_price"`
	SellPrice int `json:"sell_price"`
	IsPrimary int `json:"is_primary"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"update_at"`
}
