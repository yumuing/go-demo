package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	s := "hello"
	n := 123
	p := point{1, 2}
	fmt.Println(s, n)
	fmt.Println(p)

	fmt.Printf("s=%v\n", s)
	fmt.Printf("n=%v\n", n)
	fmt.Printf("p=%v\n", p)
	fmt.Printf("p=%+v\n", p)
	fmt.Printf("p=%#v\n", p)

	f := 3.141592653
	fmt.Println(f)
	fmt.Printf("%.2f\n", f)
}
