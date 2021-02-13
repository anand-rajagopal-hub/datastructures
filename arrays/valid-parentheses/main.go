package main

import "fmt"

type ListNode struct {
	C    rune
	Next *ListNode
}

type Stack struct {
	head *ListNode
}

func (s *Stack) push(ch rune) {
	s.head = &ListNode{
		C:    ch,
		Next: s.head,
	}
}

func (s *Stack) pop() rune {
	if s.head == nil {
		return ' '
	}
	c := s.head.C
	s.head = s.head.Next
	return c
}

func isValid(s string) bool {
	st := &Stack{}
	m := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			st.push(c)
		} else {
			t := st.pop()
			if t != m[c] {
				return false
			}
		}
	}
	return st.pop() == ' '
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
}
