//task1
// package main

// import "fmt"

// func main() {
//     // Объявление переменных
//     var weight float64 = 75.0  // вес (кг)
//     var height float64 = 1.80  // рост (метры)

//     // Вычисление BMI
//     bmi := weight / (height * height)

//	    fmt.Println("BMI:", bmi)
//	}
//
// task2
// package main

// import "fmt"

// func main() {

//     const CELSIUS_TO_FAHRENHEIT = 1.8
//     const FAHRENHEIT_OFFSET = 32

//     celsius := 25.0

//     fahrenheit := celsius*CELSIUS_TO_FAHRENHEIT + FAHRENHEIT_OFFSET

//     fmt.Println(celsius, "°C =", fahrenheit, "°F")
// }
// task3
// package main

// import "fmt"

// func main() {

//     const PI = 3.14159

//     radius := 5.0

//     area := PI * radius * radius

//     circumference := 2 * PI * radius

//     fmt.Println("Radius:", radius)
//     fmt.Println("Area:", area)
//     fmt.Println("Circumference:", circumference)
// }
// task4
package main

import "fmt"

func main() {

    var totalItems int = 100

    var price float64 = 19.99

    totalCost := float64(totalItems) * price

    fmt.Println("Total cost:", totalCost)
}