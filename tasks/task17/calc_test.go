package calculator

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5

    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

func TestDivide(t *testing.T) {
    result, err := Divide(10, 2)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if result != 5 {
        t.Errorf("Divide(10, 2) = %d; want 5", result)
    }
}

func TestDivideByZero(t *testing.T) {
    _, err := Divide(10, 0)
    if err == nil {
        t.Error("expected error for division by zero")
    }
}
//bash:
//go mod init calculator
// go test
// go run .
// go build


//Table-driven tests A common Go pattern:

// func TestAdd(t *testing.T) {
//     tests := []struct {
//         name     string
//         a, b     int
//         expected int
//     }{
//         {"positive numbers", 2, 3, 5},
//         {"negative numbers", -2, -3, -5},
//         {"mixed signs", -2, 3, 1},
//         {"with zero", 0, 5, 5},
//     }

//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//             result := Add(tt.a, tt.b)
//             if result != tt.expected {
//                 t.Errorf("Add(%d, %d) = %d; want %d",
//                     tt.a, tt.b, result, tt.expected)
//             }
//         })
//     }
// }
// package calculator

// import "fmt"

// // Example functions must start with "Example"
// func ExampleAdd() {
//     result := Add(2, 3)
//     fmt.Println(result)
//     // Output: 5
// }

// func ExampleDivide() {
//     result, err := Divide(10, 2)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }
//     fmt.Println(result)
//     // Output: 5
// }

// // Example for a specific method
// func ExampleCalculator_Multiply() {
//     c := Calculator{}
//     fmt.Println(c.Multiply(4, 5))
//     // Output: 20
// }