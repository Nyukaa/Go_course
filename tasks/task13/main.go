package main

import "fmt"

// Payment processor interface
type PaymentProcessor interface {
    Process(amount float64) error
}

// Credit card payment
type CreditCard struct {
    Number string
}

func (c CreditCard) Process(amount float64) error {
    fmt.Printf("Processing $%.2f via credit card %s\n", amount, c.Number)
    return nil
}

// PayPal payment
type PayPal struct {
    Email string
}

func (p PayPal) Process(amount float64) error {
    fmt.Printf("Processing $%.2f via PayPal account %s\n", amount, p.Email)
    return nil
}

// Checkout function works with any payment method
func Checkout(processor PaymentProcessor, amount float64) error {
    fmt.Println("Starting checkout...")
    return processor.Process(amount)
}

func main() {
    card := CreditCard{Number: "1234-5678-9012-3456"}
    paypal := PayPal{Email: "user@example.com"}

    Checkout(card, 99.99)
    Checkout(paypal, 49.99)
}