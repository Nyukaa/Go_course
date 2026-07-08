// Optional values with nil interfaces Use nil pointers to represent optional values:
package main

import "fmt"
func Compose[T any](f, g func(T) T) func(T) T {
    return func(x T) T {
        return f(g(x))
    }
}
func main() {
addOne := func(x int) int { return x + 1 }
double := func(x int) int { return x * 2 }
addThenDouble := Compose(double, addOne)
fmt.Println(addThenDouble(5))
}