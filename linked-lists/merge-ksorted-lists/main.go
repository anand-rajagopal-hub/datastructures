package main

import "fmt"

type Element struct {
	item       int
	listNumber int
}

func (e *Element) String() string {
	return fmt.Sprintf("[%d,%d]", e.item, e.listNumber)
}

type Heap struct {
	capacity  int
	itemCount int
	elements  []*Element
}

func NewHeap(k int) *Heap {
	h := &Heap{
		capacity:  k + 1,
		itemCount: 0,
		elements:  make([]*Element, k+1),
	}

	h.elements[h.itemCount] = nil
	h.itemCount++
	return h
}

type AddError struct {
	capacity  int
	itemCount int
}

func (a AddError) Error() string {
	return fmt.Sprintf("reached capacity %d %d", a.capacity, a.itemCount)
}

func (h *Heap) String() string {
	return fmt.Sprintf("%d, %v", h.itemCount, h.elements)
}

func (h *Heap) add(e *Element) error {
	if h.itemCount == h.capacity {
		return AddError{
			capacity:  h.capacity,
			itemCount: h.itemCount,
		}
	}
	h.elements[h.itemCount] = e
	h.swim(h.itemCount)
	h.itemCount++
	return nil
}

func (h *Heap) swim(index int) {
	if index == 1 {
		return
	}
	parentIndex := index / 2
	if h.elements[index].item < h.elements[parentIndex].item {
		h.elements[index], h.elements[parentIndex] = h.elements[parentIndex], h.elements[index]
	}
	h.swim(parentIndex)
}

func (h *Heap) sink(index int) {

	if index > h.itemCount/2 {
		return
	}
	lcIndex := index * 2
	rcIndex := index*2 + 1

	if rcIndex < h.itemCount && h.elements[rcIndex] != nil {
		if h.elements[rcIndex].item < h.elements[lcIndex].item && h.elements[index].item > h.elements[rcIndex].item {
			h.elements[index], h.elements[rcIndex] = h.elements[rcIndex], h.elements[index]
			h.sink(rcIndex)
		}
	}
	if lcIndex < h.itemCount && h.elements[index].item > h.elements[lcIndex].item {
		h.elements[index], h.elements[lcIndex] = h.elements[lcIndex], h.elements[index]
		h.sink(lcIndex)
	}
}

func (h *Heap) remove() *Element {
	if h.itemCount == 1 {
		return nil
	}
	min := h.elements[1]

	h.elements[1] = h.elements[h.itemCount-1]
	h.elements[h.itemCount-1] = nil
	h.itemCount--
	h.sink(1)
	return min
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var result, current *ListNode

	h := NewHeap(len(lists))
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			err := h.add(&Element{
				item:       lists[i].Val,
				listNumber: i,
			})
			if err != nil {
				panic(err)
			}
		}
	}
	for {
		element := h.remove()
		if element == nil {
			break
		}

		if result == nil {
			result = &ListNode{
				Val:  element.item,
				Next: nil,
			}
			current = result
		} else {
			nn := &ListNode{
				Val:  element.item,
				Next: nil,
			}
			current.Next = nn
			current = current.Next
		}
		lists[element.listNumber] = lists[element.listNumber].Next
		if lists[element.listNumber] != nil {
			err := h.add(&Element{
				item:       lists[element.listNumber].Val,
				listNumber: element.listNumber,
			})
			if err != nil {
				panic(err)
			}
		}
	}

	// for {
	// 	e := h.remove()
	// 	if e == nil {
	// 		break
	// 	}
	// 	nn := &ListNode{
	// 		Val:  h.remove().item,
	// 		Next: nil,
	// 	}
	// 	current.Next = nn
	// 	current = current.Next
	// }
	//remove rest of elements from the heap
	return result
}

func main() {

}
