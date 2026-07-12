package main

import (
	"fmt"

	"own_pack/calc"
)

func main() {
    result := calc.Add(10, 5)
    fmt.Println(result)
}
//bash// go mod init own_pack 

// Organizing larger projects
// Recommended structure:

// myapp/
// +-- cmd/
// |   +-- myapp/
// |       +-- main.go
// +-- internal/
// |   +-- user/
// |   |   +-- user.go
// |   +-- auth/
// |       +-- auth.go
// +-- pkg/
// |   +-- api/
// |       +-- api.go
// +-- go.mod
// +-- go.sum
// Directories explained:

// cmd/ - application entry points. Each subdirectory is a separate binary
// internal/ - private application code. Cannot be imported by other projects. Go enforces this at the compiler level
// pkg/ - public library code. Can be imported by other projects