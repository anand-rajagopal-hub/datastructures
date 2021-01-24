package main

import "fmt"

type Element struct {
	value     int
	frequency int
}

func (e *Element) String() string {
	return fmt.Sprintf("[%d, %d]", e.value, e.frequency)
}

type Heap struct {
	capacity  int
	elements  []*Element
	itemCount int
}

func NewHeap(k int) *Heap {
	h := &Heap{
		capacity:  k,
		elements:  make([]*Element, k),
		itemCount: 0,
	}
	h.elements[h.itemCount] = nil
	h.itemCount++
	return h
}

func (h *Heap) add(e *Element) {
	if len(h.elements) == h.itemCount {
		if e.frequency < h.elements[1].frequency {
			return
		}
		h.remove()
	}
	h.elements[h.itemCount] = e
	h.swim(h.itemCount)
	h.itemCount++
}

func (h *Heap) remove() *Element {
	if len(h.elements) == 1 {
		return nil
	}
	top := h.elements[1]
	h.elements[1] = h.elements[h.itemCount-1]
	h.elements[h.itemCount-1] = nil
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
	lcIndex := index * 2
	rcIndex := (index * 2) + 1
	size := h.itemCount
	currentItem := h.elements[index].frequency

	if lcIndex < size && rcIndex < size {
		if h.elements[lcIndex].frequency < h.elements[rcIndex].frequency && currentItem > h.elements[lcIndex].frequency {
			return lcIndex
		}

		if currentItem > h.elements[rcIndex].frequency {
			return rcIndex
		}
	}

	if lcIndex < size && currentItem > h.elements[lcIndex].frequency {
		return lcIndex
	}

	return -1
}

func (h *Heap) swim(index int) {
	if index == 1 {
		return
	}
	parentIndex := index / 2
	parent := h.elements[parentIndex]
	currentElement := h.elements[index]
	if parent.frequency > currentElement.frequency {
		h.elements[parentIndex], h.elements[index] = h.elements[index], h.elements[parentIndex]
	}
	h.swim(parentIndex)
}

func topKFrequent(nums []int, k int) []int {
	h := NewHeap(k + 1)
	fMap := make(map[int]int, len(nums))
	for _, item := range nums {
		v, ok := fMap[item]
		if !ok {
			fMap[item] = 1
		} else {
			fMap[item] = v + 1
		}

	}
	for key, v := range fMap {
		e := &Element{
			value:     key,
			frequency: v,
		}
		h.add(e)
	}
	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = h.remove().value
	}
	return result
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{5, 3, 1, 1, 1, 3, 73, 1}, 2))
	fmt.Println(topKFrequent([]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}, 10))
	fmt.Println(topKFrequent([]int{5, 1, -1, -8, -7, 8, -5, 0, 1, 10, 8, 0, -4, 3, -1, -1, 4, -5, 4, -3, 0, 2, 2, 2, 4, -2, -4, 8, -7, -7, 2, -8, 0, -8, 10, 8, -8, -2, -9, 4, -7, 6, 6, -1, 4, 2, 8, -3, 5, -9, -3, 6, -8, -5, 5, 10, 2, -5, -1, -5, 1, -3, 7, 0, 8, -2, -3, -1, -5, 4, 7, -9, 0, 2, 10, 4, 4, -4, -1, -1, 6, -8, -9, -1, 9, -9, 3, 5, 1, 6, -1, -2, 4, 2, 4, -6, 4, 4, 5, -5}, 7))
	// fmt.Println(to)
}
