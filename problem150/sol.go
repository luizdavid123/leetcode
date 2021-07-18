package problem150

import (
	"fmt"
	"strconv"
	"sync"
)

/*
	LeetCode Problem 150: Evaluate Reverse Polish Notation
	Level: Medium
	Description: Evaluate the value of an arithmetic expression in Reverse Polish Notation.
	Valid operators are +, -, *, and /. Each operand may be an integer or another expression.
	Note that division between two integers should truncate toward zero.
	It is guaranteed that the given RPN expression is always valid. That means the expression would always evaluate to a result, and there will not be any division by zero operation.
	Link: https://leetcode.com/problems/evaluate-reverse-polish-notation/
*/

// EvalRPN evalute the value of an arithmetic expression in Reverse Polish Notation
func EvalRPN(tokens []string) int {
	ops := map[string]func(string, string) string{
		"+": Add,
		"-": Minu,
		"*": Mutiply,
		"/": Divide,
	}
	stack := New()
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			a := stack.Pop().(string)
			b := stack.Pop().(string)
			stack.Push(ops[token](b, a))
		default:
			stack.Push(token)
		}
	}
	last := stack.Pop().(string)
	ans, _ := strconv.Atoi(last)
	return ans
}

// Add add two num
func Add(a, b string) string {
	n, _ := strconv.Atoi(a)
	m, _ := strconv.Atoi(b)
	return strconv.Itoa(n + m)
}

// Minu minu two num
func Minu(a, b string) string {
	n, _ := strconv.Atoi(a)
	m, _ := strconv.Atoi(b)
	return strconv.Itoa(n - m)
}

// Mutiply mutiply two num
func Mutiply(a, b string) string {
	n, _ := strconv.Atoi(a)
	m, _ := strconv.Atoi(b)
	return strconv.Itoa(n * m)
}

// Divide divide two num
func Divide(a, b string) string {
	n, _ := strconv.Atoi(a)
	m, _ := strconv.Atoi(b)
	return strconv.Itoa(n / m)
}

// Stack serve data as first-in-last-out manner
type Stack struct {
	items []interface{}
	size  int
	mutex *sync.Mutex
}

// New return a stack
func New() *Stack {
	return &Stack{
		items: []interface{}{},
		size:  0,
		mutex: &sync.Mutex{},
	}
}

// Empty check if the stack is empty
func (s *Stack) Empty() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.size == 0
}

// Size return the number of items in the stack
func (s *Stack) Size() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.size
}

// Pop remove the first element from the stack and return it
func (s *Stack) Pop() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.size == 0 {
		return nil
	}
	top := s.items[s.size-1]
	s.items = s.items[:s.size-1]
	s.size--
	return top
}

// Top return the first element from the stack
func (s *Stack) Top() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.size == 0 {
		return nil
	}
	return s.items[s.size-1]
}

// Push put an element at the top of the stack
func (s *Stack) Push(v interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.items = append(s.items, v)
	s.size++
}

// String return the string representation of the stack
func (s *Stack) String() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return fmt.Sprintf("%d %v", s.size, s.items)
}
