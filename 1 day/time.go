package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	t := time.Date(2022, 3, 27, 1, 25, 36, 0, time.UTC)
	t2 := time.Date(2022, 3, 27, 2, 30, 36, 0, time.UTC)
	fmt.Println(t)
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	diff := t2.Sub(t)
	fmt.Println(diff)
	fmt.Println(diff.Minutes(), diff.Seconds())
	t3, err := time.Parse("2006-01-02 15:04:05", "2022-03-27 01:25:36")
	if err != nil {
		panic(err)
	}
	fmt.Println(t3 == t)
	fmt.Println(now.Unix())
}
