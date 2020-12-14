package main

import "fmt"

func mergeSort(items []int) []int {
	if len(items) == 0 {
		return []int{}
	}
	return sort(items)
}

func sort(items []int) []int {
	if len(items) == 1 {
		return items
	}
	low := 0
	mid := len(items) / 2

	return merge(sort(items[low:mid]), sort(items[mid:]))
}

func merge(i1, i2 []int) []int {
	sortedArray := make([]int, len(i1)+len(i2))
	k := 0
	i := 0
	j := 0
	for i < len(i1) && j < len(i2) {
		if i1[i] < i2[j] {
			sortedArray[k] = i1[i]
			i++
			k++
		} else {
			sortedArray[k] = i2[j]
			j++
			k++
		}
	}

	for i < len(i1) {
		sortedArray[k] = i1[i]
		i++
		k++
	}

	for j < len(i2) {
		sortedArray[k] = i2[j]
		j++
		k++
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
