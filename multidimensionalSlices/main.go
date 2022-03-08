package main

import "fmt"

func main() {
	a := make([]int, 0)
	a = append(a, 1, 2, 3, 4)

	b := make([]int, 0)
	b = append(b, 10, 11, 12, 13, 14)

	c := make([][]int, 0)
	c = append(c, a, b)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
