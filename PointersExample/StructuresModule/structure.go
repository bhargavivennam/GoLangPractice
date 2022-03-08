package main

import "fmt"

type Point struct {
	x int
	y int
	// isOnGrid bool
}

func main() {
	// var p1 Point = Point{1, 2, false}
	// var p2 Point = Point{-4, 6, true}
	// fmt.Println(p1.x, p1.y, p2.x, p2.y)
	// p1.x = 10
	// fmt.Println(p1.x, p2.x)
	// fmt.Println(p1.isOnGrid)

	p3 := &Point{x: 3, y: 5}
	fmt.Println(p3)
	changeX(p3) //The value of x changes to {100} and if we dont use pointers * and & at Point to change, the value of x it doesn't change
	fmt.Println(p3)

}

func changeX(pt *Point) { // Here it only maintains a copy of the changed field
	pt.x = 100
	pt.y = 200
}

//Embedded struct

// type Point struct {
// 	x int32
// 	y int32
// }
// type Square struct { //Here Square struct is calling Point struct at length field
// 	area   float32
// 	length *Point
// }

// func main() {
// 	p1 := &Point{6, 8}
// 	c1 := Square{4.56, p1}
// 	fmt.Println(c1.length.x)

// }
