package leetcode

import "testing"

func heapAdjust(data []int, s int, m int) {
	pivot := data[s]
	for j := 2 * s + 1; j <= m; j = 2 * s + 1 {
		if j < m && data[j] < data[j + 1] {
			j++
		}
		if pivot >= data[j] {
			break
		}
		data[s] = data[j]
		s = j
	}
	data[s] = pivot
}

func HeapSort(data []int) {
	for i := len(data) / 2; i >= 0; i-- {
		heapAdjust(data, i, len(data) - 1)
	}
	for i := len(data) - 1; i > 0; i-- {
		data[0], data[i] = data[i], data[0]
		heapAdjust(data, 0, i - 1)
	}
}

func TestHeapSort1(t *testing.T) {
	data := []int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9}
	t.Log(data)
	HeapSort(data)
	t.Log(data)
}