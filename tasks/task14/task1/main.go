//Polymorphic collection  Store different types that share an interface:
package main

import "fmt"

type Renderer interface {
    Render() string
}

type Text struct {
    Content string
}

func (t Text) Render() string {
    return t.Content
}

type Image struct {
    URL string
}

func (i Image) Render() string {
    return fmt.Sprintf("<img src='%s'>", i.URL)
}

func main() {
    elements := []Renderer{
        Text{Content: "Hello"},
        Image{URL: "photo.jpg"},
        Text{Content: "World"},
    }

    for _, elem := range elements {
        fmt.Println(elem.Render())
    }
}