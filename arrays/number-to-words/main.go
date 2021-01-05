package main

import (
	"bytes"
	"fmt"
	"strings"
)

type Node struct {
	text string
	Next *Node
}

var groupQualifier = map[int]string{
	1: "",
	2: "Thousand",
	3: "Million",
	4: "Billion",
}

var oneToNineteen = map[int]string{
	0:  "Zero",
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
}

var tensMap = map[int]string{
	2: "Twenty",
	3: "Thirty",
	4: "Forty",
	5: "Fifty",
	6: "Sixty",
	7: "Seventy",
	8: "Eighty",
	9: "Ninety",
}

func numberToWords(num int) string {
	group := 0
	var head *Node
	if num == 0 {
		return "Zero"
	}
	for {
		r := num % 1000
		num = num / 1000
		group++
		s := w(r, group)
		if s != "" {
			head = &Node{
				text: s,
				Next: head,
			}
		}
		if num <= 0 {
			break
		}
	}
	b := &bytes.Buffer{}
	for node := head; node != nil; node = node.Next {
		b.WriteString(node.text)
		fmt.Println("'" + node.text + "'")
		if node.Next != nil {
			b.WriteString(" ")
		}
	}
	return strings.TrimSpace(b.String())
}

func w(num int, group int) string {
	if num == 0 {
		return ""
	}
	t := num / 100
	result := ""
	if t > 0 {
		result = oneToNineteen[t] + " Hundred"
	}
	t = num % 100
	if t == 0 {
		return result + " " + groupQualifier[group]
	}
	v, ok := oneToNineteen[t]
	if ok {
		if result == "" {
			result = result + v
		} else {
			result = result + " " + v
		}
	} else {
		if result == "" {
			result = result + tensMap[t/10]
		} else {
			result = result + " " + tensMap[t/10]
		}
		if t%10 > 0 {
			if result == "" {
				result = result + oneToNineteen[t%10]
			} else {
				result = result + " " + oneToNineteen[t%10]
			}
		}
	}
	result = result + " " + groupQualifier[group]
	return result
}

func main() {

}
