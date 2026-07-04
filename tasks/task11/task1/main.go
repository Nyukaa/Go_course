package main

import (
    "errors"
    "fmt"
)

type User struct {
    ID       int
    Username string
    Email    string
    Active   bool
}

func NewUser(id int, username, email string) (*User, error) {
    if username == "" {
        return nil, errors.New("username cannot be empty")
    }
    if email == "" {
        return nil, errors.New("email cannot be empty")
    }

    return &User{
        ID:       id,
        Username: username,
        Email:    email,
        Active:   true,
    }, nil
}

func (u *User) Deactivate() {
    u.Active = false
}

func (u *User) UpdateEmail(newEmail string) error {
    if newEmail == "" {
        return errors.New("email cannot be empty")
    }
    u.Email = newEmail
    return nil
}

func (u User) String() string {
    status := "active"
    if !u.Active {
        status = "inactive"
    }
    return fmt.Sprintf("User %d: %s <%s> [%s]",
        u.ID, u.Username, u.Email, status)
}

func main() {
    user, err := NewUser(1, "alice", "alice@example.com")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(user)

    user.UpdateEmail("alice.new@example.com")
    fmt.Println(user)

    user.Deactivate()
    fmt.Println(user)
}