package main

type Address struct {
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
	Street  string `json:"street"`
	City    string `json:"city"`
}

type Customer struct {
	CustomerId string  `json:"customer_id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	Address    Address `json:"address"`
}

type OrderItem struct {
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Discount  float64 `json:"discount"`
	ProductId string  `json:"product_id"`
}

type Order struct {
	OrderId    string            `json:"order_id"`
	Customer   Customer          `json:"customer"`
	Items      []OrderItem       `json:"items"`
	TotalPrice float64           `json:"total_price"`
	Status     string            `json:"status"`
	Metadata   map[string]string `json:"metadata"`
}
