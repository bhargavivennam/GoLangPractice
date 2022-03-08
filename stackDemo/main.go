package main

import "fmt"

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func main() {
	var stack Stack

	stack.Push("Stack")
	stack.Push("Implementation")
	stack.Push("Bhargavi")

	for !stack.IsEmpty() {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
}
