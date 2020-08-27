package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Product is a struct that defines a product type
type Product struct {
	gorm.Model
	Image       string  `json:"imgSrc"`
	ImagAlt     string  `json:"imgAlt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"`
	ProductName string  `json:"productName"`
	Description string  `json:"desc"`
}

// Customer is a struct that defines a customer type
type Customer struct {
	gorm.Model
	FirstName string  `json:"firstname"`
	LastName  string  `json:"lastname"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	LoggedIn  bool    `json:"loggedin"`
	Orders    []Order `json:"orders"`
}

// Order is a type that defines a order type
type Order struct {
	gorm.Model
	CustomerID   int       `json:"customer_id"`
	ProductID    int       `json:"product_id"`
	Price        float64   `json:"sell_price"`
	PurchaseDate time.Time `json:"purchase_date"`
}
