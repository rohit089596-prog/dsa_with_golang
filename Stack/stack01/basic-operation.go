package main

import "fmt"

type Stack struct {
	arr  []int
	top  int
	size int
}

func NewStack(size int) *Stack {
	return &Stack{
		arr:  make([]int, size),
		top:  -1,
		size: size,
	}
}

func (s *Stack) Push(element int) {
	if s.size-s.top > 1 {
		s.top++
		s.arr[s.top] = element
	} else {
		fmt.Println("Stack Overflow")
	}
}

func (s *Stack) Pop() {
	if s.top >= 0 {
		s.top--
	} else {
		fmt.Println("Stack UnderFlow")
	}
}

func (s *Stack) Peek() int {
	if s.top >= 0 {
		return s.arr[s.top]
	}
	fmt.Println("Stack is empty")
	return -1
}

func (s *Stack) isEmpty() bool {
	return s.top == -1
}
func main() {
	st := NewStack(5)
	st.Push(22)
	st.Push(43)
	st.Push(44)
	st.Push(23)
	st.Push(44)
	st.Push(49)
	fmt.Println(st.Peek())
	st.Pop()
	fmt.Println(st.Peek())
	st.Pop()
	fmt.Println(st.Peek())
	st.Pop()
	fmt.Println(st.Peek())
	if st.isEmpty() {
		fmt.Println("String is empty mere dosth")
	} else {
		fmt.Println("String is not empty mere dosth")
	}
}
