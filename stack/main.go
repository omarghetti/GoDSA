package main

import "fmt"

type Stackable interface {
	push(int)
	pop() (int, error)
}

type Stack struct {
	items []int
}

func (s *Stack) push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) pop() (int, error) {
	lastIndex := len(s.items) - 1
	if len(s.items) == 0 {
		return -1, fmt.Errorf("you cannot pop from an empty stack")
	}
	toRemove := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return toRemove, nil
}

func main() {
	stack := Stack{}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(5)

	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
}
