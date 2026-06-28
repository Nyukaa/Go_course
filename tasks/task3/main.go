package main

import "fmt"

func main() {

	// -------------------------
	// Task 1: Calculate total price with tax
	// -------------------------
	price := 100.0
	taxRate := 0.15

	tax := price * taxRate
	total := price + tax

	fmt.Println("Task 1")
	fmt.Println("Price:", price)
	fmt.Println("Tax:", tax)
	fmt.Println("Total:", total)
	fmt.Println()

	// -------------------------
	// Task 2: Check if a number is even
	// -------------------------
	number := 8
	isEven := number%2 == 0

	fmt.Println("Task 2")
	fmt.Println("Number:", number)
	fmt.Println("Is even:", isEven)
	fmt.Println()

	// -------------------------
	// Task 3: Validate age range
	// -------------------------
	age := 25
	isValidAge := age >= 18 && age <= 65

	fmt.Println("Task 3")
	fmt.Println("Age:", age)
	fmt.Println("Valid age:", isValidAge)
	fmt.Println()

	// -------------------------
	// Task 4: Calculate discount
	// -------------------------
	originalPrice := 200.0
	discountPercent := 20.0

	discount := originalPrice * (discountPercent / 100)
	finalPrice := originalPrice - discount

	fmt.Println("Task 4")
	fmt.Println("Original price:", originalPrice)
	fmt.Println("Discount:", discount)
	fmt.Println("Final price:", finalPrice)
	fmt.Println()

	// -------------------------
	// Task 5: Combined example
	// -------------------------
	canBuy := isEven && isValidAge

	fmt.Println("Task 5")
	fmt.Println("Can buy:", canBuy)
}