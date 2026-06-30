package main

import (
	"fmt"
	"strings"
)

// 1. Все функции вынесены на уровень пакета, под объявление import
func add(a, b int) int {
	return a + b
}

func divide(a, b float64) float64 {
	return a / b
}

func greet(name string) string {
	return "Hello, " + name
}

func isAdult(age int) bool {
	return age >= 18
}

func getUserInfo(id int) (string, int, bool) {
	name := "Alice"
	age := 25
	active := true
	return name, age, active
}
// 2. Вариативная функция sum() 
func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func validateAge(age int) bool {
    return age >= 0 && age <= 120
}

func validateEmail(email string) bool {
    return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func validateUsername(username string) (bool, string) { 
    if len(username) < 3 {
        return false, "Имя пользователя слишком короткое"
    }
    if len(username) > 20 {
        return false, "Имя пользователя слишком длинное"
    }
    return true, "Имя пользователя корректно"
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
	// 2. Изменено имя переменной с 'sum' на 'resultAdd', чтобы не ломать вызов функции sum()
	resultAdd := add(10, 5)
	quotient := divide(10.0, 3.0)
	greeting := greet("Alice")
	adult := isAdult(25)

	fmt.Println("Сумма:", resultAdd)
	fmt.Println("Частное:", quotient)
	fmt.Println("Приветствие:", greeting)
	fmt.Println("Совершеннолетний?:", adult)

	name, age, active := getUserInfo(1)
	fmt.Printf("%s, %d years, active: %t\n", name, age, active)

	// 3.вариативной функции и вывод результатов
	fmt.Println("=== Вариативная функция sum() ===")
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3)) // 6
	fmt.Println(sum(10, 20))  // 30
	fmt.Println(sum())        // 0

	nums := []int{1, 2, 3, 4}

	// 4. Передача среза в вариативную функцию обернута в fmt.Println для вывода на экран
	fmt.Println(sum(nums...)) // важно: три точки

	// 5.  проверки валидации
	fmt.Println("=== Валидация данных ===")
    age = 25
    email := "user@example.com"
    username := "alice"

    if validateAge(age) {
        fmt.Println("Возраст корректный")
    }

    if validateEmail(email) {
        fmt.Println("Email корректный")
    }

    valid, message := validateUsername(username) //returns two values: a boolean and a string message
    if valid {
		fmt.Println(message)
	}
	fmt.Println("=== Дополнительные функции ===")
	fmt.Println("Максимум 10 и 5:", max(10, 5))
    fmt.Println("Минимум 10 и 5:", min(10, 5))
    fmt.Println("Модуль -42:", abs(-42))
}
