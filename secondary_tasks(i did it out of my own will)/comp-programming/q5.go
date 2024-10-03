//5. Design your own Stack

package main

import "fmt"

const maxsize = 5

type Stack struct {
	arr [maxsize]int
	top int
}

func (s *Stack) push(value int) {
	if s.top == maxsize {
		fmt.Println("Stack overflow")
		return
	}
	s.arr[s.top] = value
	s.top++
}

func (s *Stack) pop() int {
	if s.top == 0 {
		fmt.Println("Stack underflow")
		return -1
	}
	s.top--
	return s.arr[s.top]
}

func (s *Stack) peek() int {
	if s.top == 0 {
		fmt.Println("Stack empty")
		return -1
	}
	return s.arr[s.top-1]
}

func main() {
	stack := Stack{}
	stack.push(10)
	stack.push(20)
	stack.push(30)

	fmt.Println("Peek:", stack.peek())
	fmt.Println("Pop:", stack.pop())
	fmt.Println("Peek:", stack.peek())
	fmt.Println("Pop:", stack.pop())
}
