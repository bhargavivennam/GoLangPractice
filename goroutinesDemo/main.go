package main

import (
	"fmt"
	"sync"
	"time"
)

// var wg = sync.WaitGroup{}

// func main() {
// 	wg.Add(2)
// 	go SampleRoutineFunc()
// 	go SampleRoutineFunc1()
// 	wg.Wait()
// }

// func SampleRoutineFunc() {
// 	time.Sleep(time.Millisecond)
// 	fmt.Println("Inside simplest goroutine")
// 	wg.Done()
// }
// func SampleRoutineFunc1() {
// 	time.Sleep(time.Second)
// 	fmt.Println("Inside simplest goroutine function 1")
// 	wg.Done()
// }
var m sync.Mutex
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	m.Lock()
	go SampleRoutineFunc()
	go SampleRoutineFunc1()
	wg.Wait()
	m.Unlock()
}

func SampleRoutineFunc() {
	time.Sleep(time.Millisecond)
	fmt.Println("Inside simplest goroutine")
	wg.Done()
}
func SampleRoutineFunc1() {
	time.Sleep(time.Second)
	fmt.Println("Inside simplest goroutine function 1")
	wg.Done()
}

// Panic message example

// package main

// import "fmt"

// func main() {
// 	var action int
// 	fmt.Println("Enter 1 for Student and 2 for Professional")
// 	fmt.Scanln(&action)
// 	//  Use of Switch Case in Golang
// 	switch action {
// 	case 1:
// 		fmt.Printf("I am a  Student")
// 	case 2:
// 		fmt.Printf("I am a  Professional")
// 	default:
// 		panic(fmt.Sprintf("I am a  %d", action))
// 	}
// 	fmt.Println("")
// 	fmt.Println("Enter 1 for US and 2 for UK")
// 	fmt.Scanln(&action)
// 	//  Use of Switch Case in Golang
// 	switch action {
// 	case 1:
// 		fmt.Printf("US")
// 	case 2:
// 		fmt.Printf("UK")
// 	default:
// 		panic(fmt.Sprintf("I am a  %d", action))
// 	}
// }
