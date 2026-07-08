// 1. Encapsulation and private state:
package main

import "fmt"
func NewAccount(initialBalance int) (deposit func(int), withdraw func(int), balance func() int) {
    bal := initialBalance

    deposit = func(amount int) {
        bal += amount
    }

    withdraw = func(amount int) {
        if amount <= bal {
            bal -= amount
        }
    }

    balance = func() int {
        return bal
    }

    return
}
func main() {
deposit, withdraw, balance := NewAccount(100)

deposit(50)
fmt.Println(balance())  // 150

withdraw(30)
fmt.Println(balance())  // 120
}