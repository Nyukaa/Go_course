package main

import "fmt"

type Product struct {
    ID    int
    Name  string
    Price float64
}

type CartItem struct {
    Product  Product
    Quantity int
}

type Cart struct {
    Items []CartItem
}

func NewCart() *Cart {
    return &Cart{
        Items: []CartItem{},
    }
}

func (c *Cart) AddItem(product Product, quantity int) {
    // Check if product already exists
    for i, item := range c.Items {
        if item.Product.ID == product.ID {
            c.Items[i].Quantity += quantity
            return
        }
    }

    // Add new item
    c.Items = append(c.Items, CartItem{
        Product:  product,
        Quantity: quantity,
    })
}

func (c *Cart) RemoveItem(productID int) {
    for i, item := range c.Items {
        if item.Product.ID == productID {
            c.Items = append(c.Items[:i], c.Items[i+1:]...)
            return
        }
    }
}

func (c Cart) Total() float64 {
    total := 0.0
    for _, item := range c.Items {
        total += item.Product.Price * float64(item.Quantity)
    }
    return total
}

func (c Cart) String() string {
    if len(c.Items) == 0 {
        return "Cart is empty"
    }

    result := "Cart:\n"
    for _, item := range c.Items {
        result += fmt.Sprintf("  %s x%d - $%.2f\n",
            item.Product.Name,
            item.Quantity,
            item.Product.Price*float64(item.Quantity))
    }
    result += fmt.Sprintf("Total: $%.2f", c.Total())
    return result
}

func main() {
    cart := NewCart()

    laptop := Product{ID: 1, Name: "Laptop", Price: 999.99}
    mouse := Product{ID: 2, Name: "Mouse", Price: 29.99}

    cart.AddItem(laptop, 1)
    cart.AddItem(mouse, 2)
    cart.AddItem(mouse, 1)  // Adding more of the same product

    fmt.Println(cart)

    fmt.Println("\nRemoving laptop...")
    cart.RemoveItem(1)

    fmt.Println(cart)
}