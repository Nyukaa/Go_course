package main

import (
	"fmt"
)
func swap(a *int, b *int) {
    temp := *a
    *a = *b
    *b = temp
}
func divide(a, b int, result *int, remainder *int) {
    *result = a / b
    *remainder = a % b
}

func initializeIfNil(p **int) {
    if *p == nil {
        newValue := 0
        *p = &newValue
    }
}


func main() {
	// Task 1: Swap two numbers using pointers
    x := 10
    y := 20

    swap(&x, &y)

    fmt.Println(x, y)
	// Task 2: Divide two numbers and return the quotient and remainder using pointers
	fmt.Println("=== Division using pointers ===")
	var quotient, rem int

    divide(17, 5, &quotient, &rem)

    fmt.Println("Quotient:", quotient)
    fmt.Println("Remainder:", rem)
	// Task 3: Initialize a pointer if it is nil
	fmt.Println("=== Initialize pointer if nil ===")
	var ptr *int  // nil

    fmt.Println("Before:", ptr)  // <nil>

    initializeIfNil(&ptr)

    fmt.Println("After:", ptr)   // 0xc00001234
    fmt.Println("Value:", *ptr)  // 0
}
