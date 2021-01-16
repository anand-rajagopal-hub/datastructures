package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffle(n []int) {
	rand.Seed(time.Now().UnixNano())
	for i := len(n) - 1; i > 0; i-- {
		j := rand.Int31n(int32(i + 1))
		n[i], n[j] = n[j], n[i]
	}

}

func sortArray(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	shuffle(nums)
	quicksort(nums, 0, len(nums))
	return nums
}

func partition(n []int, lo, hi int) int {
	pivotIndex := lo
	i := lo + 1
	j := hi - 1
	for {
		for ; i < hi && n[i] <= n[pivotIndex]; i++ {
		}
		for ; j >= lo+1 && n[j] > n[pivotIndex]; j-- {
		}
		if i >= j {
			break
		}
		n[i], n[j] = n[j], n[i]
	}
	n[j], n[pivotIndex] = n[pivotIndex], n[j]
	return j
}

func quicksort(n []int, lo, hi int) {
	if lo == hi {
		return
	}
	pivotIndex := partition(n, lo, hi)
	quicksort(n, lo, pivotIndex)
	quicksort(n, pivotIndex+1, hi)
}

func main() {
	fmt.Println(sortArray([]int{7, 3, 9, 5, 4, 8, 0, 1}))
	fmt.Println(sortArray([]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}))
}
