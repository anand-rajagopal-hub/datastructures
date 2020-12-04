package main

import "fmt"

func twoSum(nums []int, target int) []int {
	compliments := make(map[int]int)
	for i, value := range nums {
		ci, ok := compliments[value]
		if ok {
			return []int{ci, i}
		}
		compliments[target-value] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
