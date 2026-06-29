package main

import "fmt"

func main() {
//Task1 if 	
score := 75

if score >= 90 {
    fmt.Println("Excellent")
} else if score >= 70 {
    fmt.Println("Good")
} else if score >= 50 {
    fmt.Println("Satisfactory")
} else {
    fmt.Println("Unsatisfactory")
}
//Task2 switch
day := "Saturday"
switch day {
case "Saturday", "Sunday":
    fmt.Println("Weekend")
}
age := 30
switch {
case age < 18:
    fmt.Println("Minor")
case age >= 18 && age < 65:
    fmt.Println("Adult")
case age >= 65:
    fmt.Println("Senior")
}
//Task3 for as while loop
i := 0

for i < 5 {
    fmt.Println(i)
    i++
}
// task4 value and without index
numbers := []int{10, 20, 30, 40}

for index, value := range numbers {
    fmt.Printf("index=%d value=%d\n", index, value)
}


for _, value := range numbers {
    fmt.Println(value)
}
//Task5 for as range loop text and char
text := "Hello"

for index, char := range text {
    fmt.Printf("index=%d char=%c\n", index, char)
}
//Task6 for as for loop Метки (labels)
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

target := 5
found := false

search:
for row := 0; row < len(matrix); row++ {
    for col := 0; col < len(matrix[row]); col++ {
        if matrix[row][col] == target {
            fmt.Printf("Found %d at [%d][%d]\n", target, row, col)
            found = true
            break search // Выходим сразу из обоих циклов
        }
    }
}

if !found {
    fmt.Println("Value not found")
}
}