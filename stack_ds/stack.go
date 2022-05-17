package main

import (
	"fmt"
)

type Stack []string

func IsEmpty(s *Stack) bool {

	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)

}

func (s *Stack) Pop() (string, bool) {
	if IsEmpty(s) {
		return "", false

	}

	index := len(*s) - 1
	element := (*s)[index]

	*s = (*s)[:index]
	return element, true

}
func main() {
	var s Stack

	s.Push("Data")
	s.Push("Structure")
	s.Push("And")
	s.Push("Algorithm")

	fmt.Println(s)
	for !IsEmpty(&s) {
		pop, is_data := s.Pop()
		if is_data == true {
			fmt.Println(pop)
 
		}

	}

}
