package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
    ID    int    `json:"id,omitempty"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Fetch user from API
func getUser(id int) (*User, error) {
    url := fmt.Sprintf("https://api.example.com/users/%d", id)

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var user User
    err = json.Unmarshal(body, &user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

// Create user via API
func createUser(user User) (*User, error) {
    url := "https://api.example.com/users"

    jsonData, err := json.Marshal(user)
    if err != nil {
        return nil, err
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var created User
    err = json.Unmarshal(body, &created)
    if err != nil {
        return nil, err
    }

    return &created, nil
}

func main() {
    // Create new user
    newUser := User{
        Name:  "Alice",
        Email: "alice@example.com",
    }

    created, err := createUser(newUser)
    if err != nil {
        fmt.Println("Error creating user:", err)
        return
    }

    fmt.Printf("Created user with ID: %d\n", created.ID)

    // Fetch user
    fetched, err := getUser(created.ID)
    if err != nil {
        fmt.Println("Error fetching user:", err)
        return
    }

    fmt.Printf("Fetched: %+v\n", fetched)
}
