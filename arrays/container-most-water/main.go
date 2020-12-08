package main

import "fmt"

func maxArea(height []int) int {
	j := len(height) - 1
	i := 0

	mh := 0
	area := 0

	for {
		if i >= j {
			break
		}
		ih := height[i]
		jh := height[j]

		mh = 0

		if ih < jh {
			mh = ih
			a := mh * (j - i)
			if a > area {
				area = a
			}
			i++
		} else {
			mh = jh
			a := mh * (j - i)
			if a > area {
				area = a
			}
			j--
		}
	}
	return area
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 7, 3}))
	fmt.Println(maxArea([]int{8, 1, 6, 2, 5, 4, 8, 7, 3}))
}
