package main

import "fmt"

var m = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func gen(combo string, remaining string, result *[]string) {
	if len(remaining) == 0 {
		*result = append(*result, combo)
		return
	}
	digit := rune(remaining[0])
	for _, v := range m[digit] {
		gen(combo+string(v), remaining[1:], result)
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := make([]string, 0)

	gen("", digits, &result)
	return result
}

func main() {
	fmt.Println(letterCombinations("23"))

	fmt.Println(m['2'], m['3'], m['4'])
	fmt.Println(letterCombinations("234"))

}
