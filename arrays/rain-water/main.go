package main

import "fmt"

func trap(height []int) int {
	left := 0
	right := len(height) - 1
	leftMax := 0
	rightMax := 0
	ans := 0
	for left < right {
		if height[left] > height[right] {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				ans = ans + rightMax - height[right]
			}
			right--
			fmt.Println("right", ans)
		} else {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				ans = ans + leftMax - height[left]
			}
			left++
			fmt.Println("left", ans)
		}
	}

	return ans
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	// fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
	// fmt.Println(trap([]int{}))
	// fmt.Println(trap([]int{0, 2, 0}))
	// fmt.Println(trap([]int{9, 6, 8, 8, 5, 6, 3}))
}
