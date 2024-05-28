package main

import "fmt"

// Holding array for the values to be inserted in the heap
type MaxHeap struct {
	array []int
}

// Insert function
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Heapify Logic: parent should always be bigger than childrens
func (h *MaxHeap) maxHeapifyUp(idx int) {
	for h.array[getParentIdx(idx)] < h.array[idx] {
		h.swap(getParentIdx(idx), idx)
		idx = getParentIdx(idx)
	}
}

func (h *MaxHeap) Extract() (int, error) {
	currMax := h.array[0]
	lastIndex := len(h.array) - 1

	if len(h.array) == 0 {
		return -1, fmt.Errorf("You cannot Extract from an Empty Heap")
	}
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]
	// We Start from the root going to the childrens
	h.maxHeapifyDown(0)
	return currMax, nil
}

// Heapify Logic from top to bottom
func (h *MaxHeap) maxHeapifyDown(idx int) {
	lastIndex := len(h.array) - 1
	left, right := getLeftIdx(idx), getRightIdx(idx)

	toCompare := 0

	for left <= lastIndex {
		if left == lastIndex {
			toCompare = left
		} else if h.array[left] > h.array[right] {
			toCompare = left
		} else {
			toCompare = right
		}

		if h.array[idx] < h.array[toCompare] {
			h.swap(idx, toCompare)
			idx = toCompare
			left, right = getLeftIdx(idx), getRightIdx(idx)
		} else {
			return
		}

	}
}

// Get Parent Index
func getParentIdx(i int) int {
	return (i - 1) / 2
}

// Get Left Index
func getLeftIdx(i int) int {
	return 2*i + 1
}

// Get Right Index
func getRightIdx(i int) int {
	return 2*i + 2
}

// Swap helper function
func (h *MaxHeap) swap(idx1, idx2 int) {
	h.array[idx1], h.array[idx2] = h.array[idx2], h.array[idx1]
}

func main() {
	mHeap := &MaxHeap{}
	fmt.Println(mHeap)
	// Add values to the heap
	heapValues := []int{100, 200, 300, 400, 500, 600}
	for _, v := range heapValues {
		mHeap.Insert(v)
		fmt.Println(mHeap)
	}

	for i := 0; i < 5; i++ {
		_, err := mHeap.Extract()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(mHeap)
	}
}
