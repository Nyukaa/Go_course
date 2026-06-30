package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
//Task1 print and scan

    var weight, height float64

    fmt.Println("=== BMI Calculator ===")

    fmt.Print("Enter your weight (kg): ")
    fmt.Scan(&weight)

    fmt.Print("Enter your height (m): ")
    fmt.Scan(&height)

    bmi := weight / (height * height)

    fmt.Printf("\nYour BMI: %.2f\n", bmi)

    if bmi < 18.5 {
        fmt.Println("Category: Underweight")
    } else if bmi < 25 {
        fmt.Println("Category: Normal weight")
    } else if bmi < 30 {
        fmt.Println("Category: Overweight")
    } else {
        fmt.Println("Category: Obese")
    }

//Task2 count letters, digits, spaces, and other characters in a string

    text := "Hello, 世界! 🌍"

    var letters, digits, spaces, other int

    for _, char := range text { //in Go, we use range over the string, which gives us runes (Unicode code points) instead of bytes.

        switch {

        case unicode.IsLetter(char):
            letters++

        case unicode.IsDigit(char):
            digits++

        case unicode.IsSpace(char):
            spaces++

        default:
            other++
        }
    }
	fmt.Println("\n=== Character Count ===")
    fmt.Printf("Text: %s\n", text)
    fmt.Printf("Total characters: %d\n", len([]rune(text)))
    fmt.Printf("Letters: %d\n", letters)
    fmt.Printf("Digits: %d\n", digits)
    fmt.Printf("Spaces: %d\n", spaces)
    fmt.Printf("Other: %d\n", other)

	//Task3 check if a string contains a specific substring and manipulate the string
    var input string

	fmt.Println("\n=== Task3 ===")

    fmt.Print("Enter your name: ")
    fmt.Scanln(&input)

    // Удаляем лишние пробелы
    input = strings.TrimSpace(input)

    // Переводим в нижний регистр
    input = strings.ToLower(input)

    // Проверяем, содержит ли имя слово "admin"
    if strings.Contains(input, "admin") {

        fmt.Println("Admin access granted!")

    } else {

        // Делаем первую букву заглавной
        capitalized := strings.ToUpper(string(input[0])) + input[1:]

        fmt.Printf("Welcome, %s!\n", capitalized)
    }
}
