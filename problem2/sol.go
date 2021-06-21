package problem2

import (
	"strconv"
	"strings"
)

/*
	LeetCode Problem 2: Add Two Number
	Level: Medium
	Description: You are given two non-empty linked lists representing two non-negative integers.
	The digits are stored in reverse order, and each of their nodes contains a single digit.
	Add the two numbers and return the sum as a linked list.
*/

// ListNode represents a Node for Singly-Linked List
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode return a ListNode
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

// NewListFromString build a List from a string
func NewListFromString(str string) *ListNode {
	if len(str) == 0 {
		return nil
	}
	v, err := strconv.Atoi(string(str[0]))
	if err != nil {
		return nil
	}
	tmp := NewListNode(v)
	ptrt := tmp
	for i := 1; i < len(str); i++ {
		v, err := strconv.Atoi(string(str[i]))
		if err != nil {
			return nil
		}
		ptrt.Next = NewListNode(v)
		ptrt = ptrt.Next
	}
	return tmp
}

// String return the string representation of ListNode
func (l ListNode) String() string {
	tmp := l.DeepCopy()
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(tmp.Val))
	for tmp.Next != nil {
		tmp = tmp.Next
		builder.WriteString(strconv.Itoa(tmp.Val))
	}
	return builder.String()
}

// DeepCopy return the copy of ListNode
func (l ListNode) DeepCopy() *ListNode {
	d := NewListNode(l.Val)
	ptrd := d
	for ptrl := &l; ptrl.Next != nil; ptrl = ptrl.Next {
		ptrd.Next = NewListNode(ptrl.Next.Val)
		ptrd = ptrd.Next
	}
	return d
}

// AddTwoNumbers add two numbers and return thee sum as a linked list
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	car := 0
	if l1.Val+l2.Val >= 10 {
		car = 1
	}
	v := (l1.Val + l2.Val) % 10
	l3 := NewListNode(v)
	t := l3
	for l1.Next != nil && l2.Next != nil {
		l1 = l1.Next
		l2 = l2.Next
		v := (l1.Val + l2.Val + car) % 10
		if l1.Val+l2.Val+car >= 10 {
			car = 1
		} else {
			car = 0
		}
		l3.Next = NewListNode(v)
		l3 = l3.Next
	}
	for l1.Next != nil {
		l1 = l1.Next
		v := (l1.Val + car) % 10
		if l1.Val+car >= 10 {
			car = 1
		} else {
			car = 0
		}
		l3.Next = NewListNode(v)
		l3 = l3.Next
	}
	for l2.Next != nil {
		l2 = l2.Next
		v := (l2.Val + car) % 10
		if l2.Val+car >= 10 {
			car = 1
		} else {
			car = 0
		}
		l3.Next = NewListNode(v)
		l3 = l3.Next
	}
	if car == 1 {
		l3.Next = NewListNode(1)
	}
	return t
}
