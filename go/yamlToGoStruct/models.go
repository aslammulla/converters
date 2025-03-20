package main

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}

type Customer struct {
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	Address    Address `json:"address"`
	CustomerId string  `json:"customer_id"`
	Name       string  `json:"name"`
}

type OrderItem struct {
	Price     float64 `json:"price"`
	Discount  float64 `json:"discount"`
	ProductId string  `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
}

type Order struct {
	Items      []OrderItem       `json:"items"`
	TotalPrice float64           `json:"total_price"`
	Status     string            `json:"status"`
	Metadata   map[string]string `json:"metadata"`
	OrderId    string            `json:"order_id"`
	Customer   Customer          `json:"customer"`
}
