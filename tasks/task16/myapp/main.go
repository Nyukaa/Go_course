// Using popular external packages
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.Run()
}
// Создай модуль go mod init myapp

// Будет создан файл: go.mod
// Скачай библиотеку go mod tidy  обновляет go.mod и go.sum. go.mod содержит имя модуля и список его зависимостей.
//Файл go.sum хранит криптографические контрольные суммы и обеспечивает воспроизводимые и безопасные сборки.
// Запусти программу
// go run main.go

//  в терминале появится 
// [GIN-debug] Listening and serving HTTP on :8080