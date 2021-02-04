package main

import "fmt"

type Heap struct {
	capacity  int
	elements  []int
	itemCount int
}

func NewHeap(k int) *Heap {
	h := &Heap{
		capacity:  k + 1,
		elements:  make([]int, k+1),
		itemCount: 0,
	}
	h.elements[h.itemCount] = 0
	h.itemCount++
	return h
}

func (h *Heap) add(e int) {
	// fmt.Println("adding", e)
	if len(h.elements) == h.itemCount {
		if e < h.elements[1] {
			return
		}
		h.remove()
	}
	h.elements[h.itemCount] = e
	h.swim(h.itemCount)
	h.itemCount++
}

func (h *Heap) remove() int {
	if len(h.elements) == 1 {
		return -1
	}
	top := h.elements[1]
	h.elements[1] = h.elements[h.itemCount-1]
	h.elements[h.itemCount-1] = 0
	h.itemCount--
	h.sink(1)
	return top
}

func (h *Heap) sink(index int) {
	if index > h.itemCount/2 {
		return
	}
	exchangeIndex := h.exchangeIndex(index)
	if exchangeIndex == -1 {
		return
	}
	h.elements[index], h.elements[exchangeIndex] = h.elements[exchangeIndex], h.elements[index]
	h.sink(exchangeIndex)
}

func (h *Heap) exchangeIndex(index int) int {
	lcIndex := index << 1
	rcIndex := (index << 1) + 1
	size := h.itemCount
	currentItem := h.elements[index]

	if lcIndex < size && rcIndex < size {
		if h.elements[lcIndex] < h.elements[rcIndex] && currentItem > h.elements[lcIndex] {
			return lcIndex
		}

		if currentItem > h.elements[rcIndex] {
			return rcIndex
		}
	}

	if lcIndex < size && currentItem > h.elements[lcIndex] {
		return lcIndex
	}

	return -1
}

func (h *Heap) swim(index int) {
	if index == 1 {
		return
	}
	parentIndex := index >> 1
	parent := h.elements[parentIndex]
	currentElement := h.elements[index]
	if parent > currentElement {
		h.elements[parentIndex], h.elements[index] = h.elements[index], h.elements[parentIndex]
	}
	h.swim(parentIndex)
}

func partition(intervals [][]int, lo, hi int) int {
	pivotIndex := lo
	i := lo + 1
	j := hi
	for {
		for ; i <= hi && intervals[i][0] < intervals[pivotIndex][0]; i++ {
		}
		for ; j > lo && intervals[j][0] > intervals[pivotIndex][0]; j-- {

		}
		if i >= j {
			break
		}
		intervals[i], intervals[j] = intervals[j], intervals[i]
	}
	intervals[pivotIndex], intervals[j] = intervals[j], intervals[pivotIndex]
	return j
}

func quickSort(intervals [][]int, lo, hi int) {
	if lo >= hi {
		return
	}
	j := partition(intervals, lo, hi)
	quickSort(intervals, lo, j-1)
	quickSort(intervals, j+1, hi)
}

func minMeetingRooms(intervals [][]int) int {
	quickSort(intervals, 0, len(intervals)-1)
	h := NewHeap(len(intervals))
	h.add(intervals[0][1])
	for i, v := range intervals {
		if i == 0 {
			continue //ignore the first element
		}
		if h.elements[1] <= v[0] {
			h.remove()
			h.add(v[1])
		} else {
			h.add(v[1])
		}
	}
	return h.itemCount - 1
}

func main() {
	intervals := [][]int{
		{5, 10},
		{0, 30},
		{15, 20},
	}
	fmt.Println(minMeetingRooms(intervals))

	intervals = [][]int{
		{7, 10},
		{2, 4},
	}
	fmt.Println(minMeetingRooms(intervals))
}
