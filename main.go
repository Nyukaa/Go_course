package main

// Импортируем сразу несколько пакетов с помощью круглых скобок
import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// 1. Работаем с текстом через пакет strings
	message := "Welcome to Go programming!"
	
	upperMessage := strings.ToUpper(message)
	lowerMessage := strings.ToLower(message)

	// 2. Выводим результаты через пакет fmt
	fmt.Println("Оригинал:   ", message)
	fmt.Println("В верхнем:  ", upperMessage)
	fmt.Println("В нижнем:   ", lowerMessage)

	// Разделитель строк
	fmt.Println(strings.Repeat("-", 30))

	// 3. Работаем со временем через пакет time
	currentTime := time.Now()
	// Форматируем дату в привычный вид: День-Месяц-Год Часы:Минуты:Секунды
	formattedTime := currentTime.Format("02-01-2006 15:04:05")
	
	fmt.Println("Текущее время устройства:", formattedTime)
}
//GOOS=windows GOARCH=amd64 go build -o myapp.exe main.go
