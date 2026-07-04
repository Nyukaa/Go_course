package main

import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

func NewRectangle(width, height float64) *Rectangle {
    return &Rectangle{
        Width:  width,
        Height: height,
    }
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func (r Rectangle) String() string {
    return fmt.Sprintf("Rectangle(%.1f x %.1f)", r.Width, r.Height)
}

func main() {
    rect := NewRectangle(10, 5)

    fmt.Println(rect)
    fmt.Printf("Area: %.1f\n", rect.Area())
    fmt.Printf("Perimeter: %.1f\n", rect.Perimeter())

    rect.Scale(2)

    fmt.Println(rect)
    fmt.Printf("Area: %.1f\n", rect.Area())
}
