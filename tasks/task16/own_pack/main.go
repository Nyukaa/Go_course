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