package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	n, _ := strconv.ParseInt("111", 10, 64)
	fmt.Println(n)

	n, _ = strconv.ParseInt("0x1000", 0, 64)
	fmt.Println(n)

	n2, _ := strconv.Atoi("123")
	fmt.Println(n2)

	n2, err := strconv.Atoi("AAA")
	fmt.Println(n2, err)
}
