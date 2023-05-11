package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello"
	fmt.Println(strings.Contains(a, "ll"))
	fmt.Println(strings.Count(a, "l"))
	fmt.Println(strings.HasPrefix(a, "he"))
	fmt.Println(strings.HasSuffix(a, "llo"))
	fmt.Println(strings.Index(a, "ll"))
	fmt.Println(strings.Join([]string{"he", "llo"}, "-"))
	fmt.Println(strings.Repeat(a, 2))
	fmt.Println(strings.Replace(a, "e", "E", -1))
	fmt.Println(strings.Split("a-b-c", "-"))
	fmt.Println(strings.ToLower(a))
	fmt.Println(strings.ToUpper(a))
	fmt.Println(len(a))
	b := "yumuing 博客"
	fmt.Println(len(b))
}
