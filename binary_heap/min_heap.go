package main

import "fmt"

type MinHeap struct {
	array []int
}

func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return (i * 2) + 1
}

func right(i int) int {
	return (i * 2) + 2
}

func (h *MinHeap) swap(first, second int) {

	h.array[first], h.array[second] = h.array[second], h.array[first]
}

func (h *MinHeap) insert(index int) {
	h.array = append(h.array, index)

	h.heapifyUp(len(h.array) - 1)
}

func (h *MinHeap) heapifyUp(index int) {
	for h.array[parent(index)] > h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) remove() int {
	removed := h.array[0]
	L := len(h.array) - 1
	if len(h.array) == 0 {
		return -1
	}
	//move the last node to the 1st position
	h.array[0] = h.array[L]

	//re-arrange the slice by removing last node i.e previous moved node
	h.array = h.array[:L]
	h.heapifyDown(0)
	return removed

}
func (h *MinHeap) heapifyDown(index int) {
	lastIndexToCheck := len(h.array) - 1

	L, R := left(index), right(index)
	childToCompare := index
	for L <= lastIndexToCheck {
		//compare the child
		if L == lastIndexToCheck || h.array[L] < h.array[R] {
			childToCompare = L
		} else {
			childToCompare = R
		}
		//swap with the smallest
		if h.array[index] > h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			L, R = left(index), right(index)

		} else {
			return
		}
	}

}

func main() {
	var mh MinHeap
	for i := 0; i < 10; i++ {

		mh.insert(i + 3*i)

	}
	fmt.Println(mh.array)
	for i := 0; i < 10; i++ {
		if len(mh.array)!=0 {
			fmt.Println(mh.array[i])

		}
		mh.remove()
		fmt.Println("then ", mh.array)

	}

}
