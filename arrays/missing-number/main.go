package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for _, v := range nums {
		sum = sum - v
	}
	return sum
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber([]int{0}))
}
