package payment

import "fmt"

type (
	CreditCardPayment   struct{ CardNumber string }
	PayPalPayment       struct{ Email string }
	BankTransferPayment struct{ AccountNumber string }
	CryptoPayment       struct{ WalletAddress string }
)

func (c *CreditCardPayment) ProcessPayment(amount float64) error {
	fmt.Printf("Processing payment of %.2f via Credit Card %s\n", amount, c.CardNumber)
	return nil
}

func (p *PayPalPayment) ProcessPayment(amount float64) error {
	fmt.Printf("Processing payment of %.2f via PayPal (%s)\n", amount, p.Email)
	return nil
}

func (b *BankTransferPayment) ProcessPayment(amount float64) error {
	fmt.Printf("Processing payment of %.2f via Bank Transfer to account %s\n", amount, b.AccountNumber)
	return nil
}

func (c *CryptoPayment) ProcessPayment(amount float64) error {
	fmt.Printf("Processing payment of %.2f via Crypto to wallet %s\n", amount, c.WalletAddress)
	return nil
}
