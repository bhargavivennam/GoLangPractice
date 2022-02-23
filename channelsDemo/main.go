package main

import (
	"fmt"
	"time"
)

// func ChannelsDemo() {
// 	ch := make(chan int)
// 	ch <- 22
// 	fmt.Println(ch)
// }

// func ChannelsDemo2(){
// 	go ChannelsDemo(ch)

// }

func Add(num int, ch chan int) {
	res := num + 3
	fmt.Println(res)
	ch <- res
}

func Sub(num int, ch chan int) {
	res1 := num - 3
	time.Sleep(time.Second)
	fmt.Println(res1)
	ch <- res1

}
func main() {
	ch := make(chan int)
	go Add(3, ch)
	go Sub(10, ch)
	output := <-ch
	fmt.Println(output)
}
