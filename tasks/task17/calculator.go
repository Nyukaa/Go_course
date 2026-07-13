package calculator

import (
	"errors"
	"fmt"
)

func Add(a, b int) int {
    return a + b
}

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
func main() {
	// Example usage of the calculator functions
	sum := Add(10, 5)	
	fmt	.Println("Sum:", sum) // Output: Sum: 15
}