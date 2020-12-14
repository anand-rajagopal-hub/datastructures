package main

import "fmt"

func mergeSort(items []int) []int {
	if len(items) == 0 {
		return []int{}
	}
	return sort(items, 0, len(items))
}

func sort(items []int, low, high int) []int {
	if len(items) == 1 {
		return items
	}
	low = 0
	high = len(items)

	return merge(sort(items[low:high/2], low, high/2), sort(items[high/2:high], high/2, high))
}

func merge(i1, i2 []int) []int {
	sortedArray := []int{}
	i := 0
	j := 0
	for i < len(i1) && j < len(i2) {
		if i1[i] < i2[j] {
			sortedArray = append(sortedArray, i1[i])
			i++
		} else {
			sortedArray = append(sortedArray, i2[j])
			j++
		}
	}

	for i < len(i1) {
		sortedArray = append(sortedArray, i1[i])
		i++
	}

	for j < len(i2) {
		sortedArray = append(sortedArray, i2[j])
		j++
	}
	return sortedArray
}

func main() {
	fmt.Println(mergeSort([]int{7, 3, 9, 5, 4, 8, 0, 1}))
}

/*
* 7 3 9 5 4 8 0 1
* 7 3 9 5
* 7 3
* 7
 */
