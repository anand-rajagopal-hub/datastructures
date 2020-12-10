package main

import (
	"fmt"
)

func myAtoi(s string) int {
	allowedChars := map[rune]bool{
		'+': true,
		'-': true,
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}

	f := false
	numberRunes := []rune{}
	var multiplier int
	multiplier = 1
	chars := []rune(s)
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		_, ok := allowedChars[c]
		if ok {
			if (c == '-' || c == '+') && f {
				break
			}
			f = true
			if c == '-' && i == len(chars)-1 {
				continue
			}
			if c == '-' {
				_, nextcharok := allowedChars[chars[i+1]]
				if nextcharok {
					multiplier = -1
				}
			} else if c != '+' {
				numberRunes = append(numberRunes, c)
			}
		} else if c != ' ' && c != '+' && !f {
			break
		} else {
			break
		}
	}
	return convert(multiplier, numberRunes)
}

func convert(m int, runes []rune) int {
	var total int
	minInt := 1 << 31
	maxInt := (1 << 31) - 1
	for i := 0; i < len(runes); i++ {
		if total > maxInt/10 || (total == maxInt/10 && int(runes[i]-48) > maxInt%10) {
			if m == -1 {
				return m * minInt
			}
			return m * maxInt
		}
		total = total*10 + int((runes[i] - 48))
	}
	return m * int(total)
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("-"))
	fmt.Println(myAtoi("+1"))
	fmt.Println(myAtoi("21474836460"))
}
