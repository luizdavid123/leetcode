package problem155

import "container/list"

/*
	LeetCode Problem 155: Min Stack
	Level: Easy
	Description: Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
	Implement the MinStack class:
	1. MinStack() initializes the stack object.
	2. void push(val) pushes the element val onto the stack.
	3. void pop() removes the element on the top of the stack.
  4. int top() gets the top element of the stack.
  5. int getMin() retrieves the minimum element in the stack.
	Link: https://leetcode.com/problems/min-stack/
*/

type Record struct {
	val int
	min int
}

type MinStack struct {
	data *list.List
	size int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: list.New(),
		size: 0,
	}
}

func (m *MinStack) Push(val int) {
	record := Record{}
	if m.size == 0 {
		record = Record{val: val, min: val}
	} else {
		data := m.data.Front().Value.(*Record)
		record = Record{val: val, min: MinInt(data.min, val)}
	}
	m.size++
	m.data.PushFront(&record)
}

func (m *MinStack) Pop() {
	if m.size == 0 {
		return
	}
	m.size--
	m.data.Remove(m.data.Front())
}

func (m *MinStack) Top() int {
	if m.size == 0 {
		return -1
	}
	return m.data.Front().Value.(*Record).val
}

func (m *MinStack) GetMin() int {
	if m.size == 0 {
		return -1
	}
	return m.data.Front().Value.(*Record).min
}

// MinInt return the smaller one
func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
