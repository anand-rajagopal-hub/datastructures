package main

import "fmt"

func reorderLogFiles(logs []string) []string {
	digitLogs := []string{}
	strLogs := []string{}
	for i := 0; i < len(logs); i++ {
		log := logs[i]

		lr := []rune(log)
		for j := 0; j < len(lr); j++ {
			if lr[j] == ' ' {
				if lr[j+1] >= 96 {
					strLogs = append(strLogs, log)
				} else {
					digitLogs = append(digitLogs, log)
				}
				break
			}
		}
	}
	sortedLogs := mergeSort(strLogs)
	for _, v := range digitLogs {
		sortedLogs = append(sortedLogs, v)
	}
	return sortedLogs
}

func mergeSort(items []string) []string {
	if len(items) == 0 {
		return []string{}
	}
	return sort(items, 0, len(items))
}

func sort(items []string, low, high int) []string {
	if len(items) == 1 {
		return items
	}
	low = 0
	high = len(items)

	return merge(sort(items[low:high/2], low, high/2), sort(items[high/2:high], high/2, high))
}

func merge(i1, i2 []string) []string {
	sortedArray := []string{}
	i := 0
	j := 0
	for i < len(i1) && j < len(i2) {
		i1Runes := []rune(i1[i])
		i2Runes := []rune(i2[j])

		i1Str := ""
		i2Str := ""
		i1ID := ""
		i2ID := ""
		for k, v := range i1Runes {
			if v == ' ' {
				i1Str = string(i1Runes[k:])
				i1ID = string(i1Runes[0:k])
				break
			}
		}
		for k, v := range i2Runes {
			if v == ' ' {
				i2Str = string(i2Runes[k:])
				i2ID = string(i2Runes[0:k])
				break
			}
		}
		if i1Str < i2Str {
			sortedArray = append(sortedArray, i1[i])
			i++
		} else if i1Str == i2Str {
			if i1ID < i2ID {
				sortedArray = append(sortedArray, i1[i])
				i++
			} else {
				sortedArray = append(sortedArray, i2[j])
				j++
			}

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
	fmt.Println(reorderLogFiles([]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}))
}
