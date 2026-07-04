package main

import (
	"fmt"
)
type Contact struct {
    Phone string
    Email string
}

type Social struct {
    Twitter  string
    LinkedIn string
}

type Person struct {
    Name string
    Contact
    Social
}

type User struct {
	Name string
	Age  int
}

func incrementAgeGood(u *User) {
    u.Age++
}
//bad practice, passing by value
// func incrementAgeBad(u User) {
//     u.Age++
// }

func main() {
// 1. Struct embedding
fmt.Println("=== Struct Embedding ===")	


person := Person{
    Name: "Bob",
    Contact: Contact{
        Phone: "123-456",
    },
    Social: Social{
        Twitter: "@bob",
    },
}



fmt.Println(person.Phone)    // 123-456
fmt.Println(person.Twitter)  // @bob
// 2. Pointer receiver
fmt.Println("=== Pointer Receiver ===")
user := User{
    Name:"Alice",
    Age:25,
}

incrementAgeGood(&user)

fmt.Println(user.Age) // 26
}