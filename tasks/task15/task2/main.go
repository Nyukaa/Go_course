// Middleware pattern Chain interface implementations:
package main

import "fmt"

type Handler interface {
    Handle(request string) string
}

type BaseHandler struct{}

func (h BaseHandler) Handle(request string) string {
    return fmt.Sprintf("Processed: %s", request)
}

type LoggingMiddleware struct {
    next Handler
}

func (m LoggingMiddleware) Handle(request string) string {
    fmt.Printf("[LOG] Request: %s\n", request)
    result := m.next.Handle(request)
    fmt.Printf("[LOG] Response: %s\n", result)
    return result
}

type AuthMiddleware struct {
    next Handler
}

func (m AuthMiddleware) Handle(request string) string {
    fmt.Println("[AUTH] Checking credentials...")
    return m.next.Handle(request)
}

func main() {
    handler := AuthMiddleware{
        next: LoggingMiddleware{
            next: BaseHandler{},
        },
    }

    result := handler.Handle("GET /users")
    fmt.Println("Final result:", result)
}