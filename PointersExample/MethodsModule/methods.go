package main

import "fmt"

type Student struct {
	name  string
	grade []int
	age   int
}

func (s *Student) getAvggrade() float32 {
	sum := 0
	for _, v := range s.grade {
		sum += v
	}
	return float32(sum) / float32(len(s.grade))
}

func (st *Student) getMaxGrade() int {
	studmax := 0
	for _, v := range st.grade {
		if studmax < v {
			studmax = v
		}
	}
	return studmax
}

func main() {
	s1 := Student{"abc", []int{70, 65, 80}, 20}
	s2 := Student{"def", []int{50, 35, 40}, 21}
	average := s1.getAvggrade()
	fmt.Println(average)
	average2 := s2.getAvggrade()
	fmt.Println(average2)
	s1.grade[1] = 82 //By using pointer in the above method we can change the values
	fmt.Println(s1.grade)

	max := s1.getMaxGrade()
	fmt.Println(max)
	max2 := s2.getMaxGrade()
	fmt.Println(max2)
}
