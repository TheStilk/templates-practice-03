package delivery

import (
	"fmt"

	"prac03/utils"
)

type (
	CourierDelivery     struct{ Address string }
	PostDelivery        struct{ PostOffice string }
	PickUpPointDelivery struct{ PointName string }
	DroneDelivery       struct{ DroneID string }
)

func (c *CourierDelivery) DeliverOrder(orderID string, products []utils.Product) error {
	fmt.Printf("Courier delivering order %s to %s\n", orderID, c.Address)
	return nil
}

func (p *PostDelivery) DeliverOrder(orderID string, products []utils.Product) error {
	fmt.Printf("Sending order %s via postal service to %s\n", orderID, p.PostOffice)
	return nil
}

func (p *PickUpPointDelivery) DeliverOrder(orderID string, products []utils.Product) error {
	fmt.Printf("Order %s ready for pickup at %s\n", orderID, p.PointName)
	return nil
}

func (d *DroneDelivery) DeliverOrder(orderID string) error {
	fmt.Printf("Drone %s delivering order %s\n", d.DroneID, orderID)
	return nil
}
