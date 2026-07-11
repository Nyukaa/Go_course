package user

// Экспортируемая структура
type User struct {
    Name string // экспортируемое поле
    age  int    // неэкспортируемое поле
}

// Экспортируемая функция
func NewUser(name string, age int) *User {
    return &User{Name: name, age: age}
}

// Неэкспортируемая функция
func validateAge(age int) bool {
    return age >= 0
}

//Использование из другого пакета
//package main

//import "myapp/user"

// func main() {
//     u := user.NewUser("Alice", 30)

//     fmt.Println(u.Name) // Работает

    // fmt.Println(u.age)      // Ошибка
    // user.validateAge(25)    // Ошибка
//}