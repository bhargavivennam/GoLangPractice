package main

import "fmt"

func main() {
	//Example for integers on & and *, & -> referencing a variable and * -> dereferencing a variable
	// var x int
	// x = 7
	// y := &x

	// fmt.Println(x, y)

	// *y = 10
	// fmt.Println(x, y)

	//Example for strings
	// toChange := "Hello"
	// var pointer *string = &toChange
	// fmt.Print(pointer, &pointer)

	toChange := "Hello"
	fmt.Println(toChange)
	changeValue(&toChange) //Here the value Hello is modified to Changed using pointers
	fmt.Println(toChange)
	changeValue2(toChange) //The string here doesn't change because strings are immutable
	fmt.Println(toChange)

}

//Example to create functions for pointers
func changeValue(str *string) {
	*str = "Changed!"
}
func changeValue2(str string) {
	str = "changed"

}
