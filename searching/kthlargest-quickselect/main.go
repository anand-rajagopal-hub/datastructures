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

func partition(n []int, lo, hi int) int {
	i := lo + 1
	j := hi
	pivot := lo

	for {
		for i <= hi && n[i] <= n[pivot] {
			i++
		}
		for j >= lo+1 && n[j] > n[pivot] {
			j--
		}
		if j <= i {
			break
		}
		n[i], n[j] = n[j], n[i]
	}
	n[j], n[pivot] = n[pivot], n[j]
	return j
}

func quickSelect(n []int, lo, hi, k int) int {
	pivotIndex := partition(n, lo, hi)
	if k == pivotIndex {
		return n[pivotIndex]
	} else if k < pivotIndex {
		return quickSelect(n, lo, pivotIndex-1, k)
	}
	return quickSelect(n, pivotIndex+1, hi, k)
}

func findKthLargest(nums []int, k int) int {
	shuffle(nums)
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func main() {
	fmt.Println("------------>", findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println("------------>", findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	fmt.Println("------------>", findKthLargest([]int{2, 1}, 2))
	fmt.Println("------------>", findKthLargest([]int{7, 6, 5, 4, 3, 2, 1}, 2))
}
