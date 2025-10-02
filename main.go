package main

import (
	"fmt"

	"prac03/delivery"
	"prac03/discount"
	"prac03/interfaces"
	"prac03/notification"
	"prac03/order"
	"prac03/payment"
	"prac03/utils"
)

func main() {
	fmt.Printf("Internet Store Order System:\n")

	order1 := order.NewOrder("ORDER-001")

	order1.AddProduct(utils.Product{ID: "PROD001", Name: "Laptop", Price: 999.99, Quantity: 1})
	order1.AddProduct(utils.Product{ID: "PROD002", Name: "Mouse", Price: 25.50, Quantity: 2})

	total := order1.CalculateTotal()
	fmt.Printf("Total before discount: $%.2f\n", total)

	calculator := discount.DiscountCalculator{
		Rules: []interfaces.DiscountRule{
			&discount.PercentageDiscount{Percent: 10},
			&discount.BulkDiscount{MinItems: 3, Discount: 5},
		},
	}
	discounted := order1.ApplyDiscount(calculator)
	fmt.Printf("Total after discounts: $%.2f\n", discounted)

	order1.Payment = &payment.CreditCardPayment{CardNumber: "4400 4000 4000 1234"}

	order1.Delivery = &delivery.CourierDelivery{Address: "Almaty, Satbayeva 22"}

	order1.Notification = &notification.EmailNotification{Email: "mad@thestilk.org"}

	if err := order1.ProcessPayment(); err != nil {
		fmt.Printf("Payment failed: %v\n", err)
		return
	}

	if err := order1.Notify("Your order has been paid successfully!"); err != nil {
		fmt.Printf("Notification failed: %v\n", err)
		return
	}

	if err := order1.Deliver(); err != nil {
		fmt.Printf("Delivery failed: %v\n", err)
		return
	}

	if err := order1.Notify("Your order is on the way!"); err != nil {
		fmt.Printf("Notification failed: %v\n", err)
		return
	}

	fmt.Printf("\nOrder %s status: %s\n", order1.ID, order1.GetStatus())
}
