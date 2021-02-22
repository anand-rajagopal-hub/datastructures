package main

import (
	"fmt"
	"strings"
)

func paren(n int, cur string, arr *[]string) {
	if len(cur) == (n * 2) {
		*arr = append(*arr, cur)
		return
	}

	lp := strings.Count(cur, "(")
	rp := strings.Count(cur, ")")
	if lp < n {
		paren(n, cur+"(", arr)
	}
	if rp < lp {
		paren(n, cur+")", arr)
	}
}

func generateParenthesis(n int) []string {
	arr := make([]string, 0)
	paren(n, "", &arr)
	return arr
}

func main() {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(4))
}
