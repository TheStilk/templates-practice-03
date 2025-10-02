package discount

import (
	"math"

	"prac03/interfaces"
	"prac03/utils"
)

type (
	PercentageDiscount  struct{ Percent float64 }
	FixedAmountDiscount struct{ Amount float64 }
	BulkDiscount        struct {
		MinItems int
		Discount float64
	}
	DiscountCalculator struct{ Rules []interfaces.DiscountRule }
	SeasonalDiscount   struct {
		Season string
		Amount float64
	}
)

func (p *PercentageDiscount) Calculate(price float64, _ []utils.Product) float64 {
	return price * (1 - p.Percent/100)
}

func (f *FixedAmountDiscount) Calculate(price float64, _ []utils.Product) float64 {
	return math.Max(0, price-f.Amount)
}

func (b *BulkDiscount) Calculate(price float64, products []utils.Product) float64 {
	totalItems := 0
	for _, p := range products {
		totalItems += p.Quantity
	}
	if totalItems >= b.MinItems {
		return price * (1 - b.Discount/100)
	}
	return price
}

func (dc *DiscountCalculator) Calculate(originalPrice float64, products []utils.Product) float64 {
	price := originalPrice
	for _, rule := range dc.Rules {
		price = rule.Calculate(price, products)
	}
	return price
}

func (s *SeasonalDiscount) Calculate(price float64, _ []utils.Product) float64 {
	return price * (1 - s.Amount/100)
}
