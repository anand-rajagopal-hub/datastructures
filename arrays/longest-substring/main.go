package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	chars := []rune(s)
	m := make(map[rune]int)
	longest := 0
	nsp := 0
	for i := 0; i < len(chars); i++ {
		loc, ok := m[chars[i]]
		if ok {
			if loc >= nsp {
				nsp = loc + 1
			}
			delete(m, chars[i])
		}
		m[chars[i]] = i
		t := i - nsp + 1
		if longest < t {
			longest = t
		}
	}
	return longest
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("abcabcgfhijk"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring("bcd"))
	fmt.Println(lengthOfLongestSubstring("abba"))
}
