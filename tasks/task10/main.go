package main

import (
	"errors"
	"fmt"
	"strings"
)

// validateUsername checks if the provided username meets certain criteria.
func validateUsername(username string) error {
    if len(username) == 0 {
        return errors.New("username cannot be empty")
    }
    if len(username) < 3 {
        return errors.New("username must be at least 3 characters")
    }
    if len(username) > 20 {
        return errors.New("username cannot exceed 20 characters")
    }
    if strings.Contains(username, " ") {
        return errors.New("username cannot contain spaces")
    }
    return nil
}
func safeDivide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
//
func findInSlice(slice []int, target int) (int, error) {
    if slice == nil {
        return -1, errors.New("slice is nil")
    }

    for i, value := range slice {
        if value == target {
            return i, nil  // Found - return index, no error
        }
    }

    return -1, fmt.Errorf("value %d not found in slice", target)
}

func getUser(users map[int]string, id int) (string, error) {
    if users == nil {
        return "", errors.New("users map is nil")
    }

    name, exists := users[id]
    if !exists {
        return "", fmt.Errorf("user with id %d not found", id)
    }

    return name, nil
}

func main() {
	// Task 1: Validate usernames
    usernames := []string{"alice", "ab", "a very long username here", "bob smith"}

    for _, name := range usernames {
        err := validateUsername(name)
        if err != nil {
            fmt.Printf("Invalid username '%s': %v\n", name, err)
        } else {
            fmt.Printf("Valid username: %s\n", name)
        }
    }
	// Task 2: Safe division
	fmt.Println("=== Safe Division ===")
	  result, err := safeDivide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)  // 5
    }

    result, err = safeDivide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)  // division by zero
    } else {
        fmt.Println("Result:", result)
    }

	numbers := []int{10, 20, 30, 40, 50}

    index, err := findInSlice(numbers, 30)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Found at index: %d\n", index)
    }

    index, err = findInSlice(numbers, 99)
    if err != nil {
        fmt.Println("Error:", err)  // value 99 not found in slice
    }
// Task 4: Get user by ID from a map
    fmt.Println("=== Get User by ID ===")
	users := map[int]string{
        1: "Alice",
        2: "Bob",
        3: "Charlie",
    }

    name, err := getUser(users, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("User:", name)  // Bob
    }

    name, err = getUser(users, 99)
    if err != nil {
        fmt.Println("Error:", err)  // user with id 99 not found
    }
}