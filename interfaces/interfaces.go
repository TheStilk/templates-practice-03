package interfaces

import (
	"prac03/utils"
)

type IPayment interface {
	ProcessPayment(amount float64) error
}

type IDelivery interface {
	DeliverOrder(orderID string, products []utils.Product) error
}

type INotification interface {
	SendNotification(message string) error
}

type DiscountRule interface {
	Calculate(price float64, products []utils.Product) float64
}
