package main

import "fmt"

type R interface {
	getMenu() int
}

type B struct {
	a int
	b int
}

type C struct {
	a int
	b int
	d string
}

func main() {

	X := C{
		a: 1,
		b: 2,
		d: "hi",
	}

	y := B{
		a: 1,
		b: 3,
	}

	fmt.Printf("the menu is %d \n", Menu(X))
	fmt.Printf("the menu is %d", Menu(y))
}

func Menu(r R) int {
	return r.getMenu()
}
func (b B) getMenu() int {
	return b.a + b.b
}

func (c C) getMenu() int {
	return c.a + c.b
}
