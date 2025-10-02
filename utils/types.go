package utils

type Product struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}

type OrderStatus string

const (
	OrderCreated   OrderStatus = "created"
	OrderPaid      OrderStatus = "paid"
	OrderShipped   OrderStatus = "shipped"
	OrderDelivered OrderStatus = "delivered"
)
