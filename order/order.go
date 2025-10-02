package order

import (
	"fmt"

	"prac03/discount"
	"prac03/interfaces"
	"prac03/utils"
)

type Order struct {
	ID           string
	Products     []utils.Product
	Payment      interfaces.IPayment
	Delivery     interfaces.IDelivery
	Notification interfaces.INotification
	Status       utils.OrderStatus
	totalPrice   float64
}

func NewOrder(id string) *Order {
	return &Order{
		ID:       id,
		Products: make([]utils.Product, 0),
		Status:   utils.OrderCreated,
	}
}

func (o *Order) AddProduct(product utils.Product) {
	o.Products = append(o.Products, product)
}

func (o *Order) CalculateTotal() float64 {
	var total float64
	for _, p := range o.Products {
		total += p.Price * float64(p.Quantity)
	}
	o.totalPrice = total
	return total
}

func (o *Order) ApplyDiscount(calculator discount.DiscountCalculator) float64 {
	discounted := calculator.Calculate(o.totalPrice, o.Products)
	o.totalPrice = discounted
	return discounted
}

func (o *Order) ProcessPayment() error {
	if o.Payment == nil {
		return fmt.Errorf("payment method not set")
	}
	err := o.Payment.ProcessPayment(o.totalPrice)
	if err != nil {
		return err
	}
	o.Status = utils.OrderPaid
	return nil
}

func (o *Order) Deliver() error {
	if o.Delivery == nil {
		return fmt.Errorf("delivery method not set")
	}
	err := o.Delivery.DeliverOrder(o.ID, o.Products)
	if err != nil {
		return err
	}
	o.Status = utils.OrderShipped
	return nil
}

func (o *Order) Notify(message string) error {
	if o.Notification == nil {
		return fmt.Errorf("notification method not set")
	}
	return o.Notification.SendNotification(message)
}

func (o *Order) GetStatus() utils.OrderStatus {
	return o.Status
}
