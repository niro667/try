package main

import (
	"fmt"
	"strconv"
)

func main() {
	z := "100.5"
	num1, _ := strconv.ParseFloat(z, 32)
	fmt.Println("Hello, wor", num1, "bye!!")

}
