package main

import (
	"fmt"
)

// 1 array of numbers and find the maximum value
func findMax(numbers []int) int {
    if len(numbers) == 0 {
        return 0
    }

    max := numbers[0]
    for _, n := range numbers {
        if n > max {
            max = n
        }
    }
    return max
}

func filterEven(numbers []int) []int {
    result := make([]int, 0, len(numbers))

    for _, num := range numbers {
        if num%2 == 0 {
            result = append(result, num)
        }
    }

    return result
}
func reverse(slice []int) []int {
    reversed := make([]int, len(slice))

    for i, value := range slice {
        reversed[len(slice)-1-i] = value
    }

    return reversed
}
func main() {
	// Task 1: Calculate the average of an array of temperatures
	temperatures := []float64{20.5, 22.0, 19.5, 21.0, 23.5}

sum := 0.0
for _, t := range temperatures {
    sum += t
}

avg := sum / float64(len(temperatures))

fmt.Printf("Average temperature: %.2f\n", avg)

fmt.Println("Max temperature:", findMax([]int{1, 5, 3, 9, 2}))

// Task 2: Filter even numbers from an array
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    even := filterEven(numbers)
   fmt.Println("=== Even Numbers Filter ===")
    fmt.Println("Even numbers:", even)  // [2 4 6 8 10]
// Task 3: Reverse an array of numbers

    numbers = []int{1, 2, 3, 4, 5} // Re-initialize the numbers slice
    reversed := reverse(numbers)

    fmt.Println("Original:", numbers)   // [1 2 3 4 5]
    fmt.Println("Reversed:", reversed)  // [5 4 3 2 1]
// Task 4: Read numbers from user input and store them in an array
	 var count int
    fmt.Print("How many numbers? ")
    fmt.Scan(&count)

    numbers = make([]int, 0, count)

    for i := 0; i < count; i++ {
        var num int
        fmt.Printf("Enter number %d: ", i+1)
        fmt.Scan(&num)
        numbers = append(numbers, num)
    }

    fmt.Println("You entered:", numbers)



}