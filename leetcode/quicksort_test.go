package leetcode

import "testing"

func partition(data []int, low, high int) int {
	pivot := data[low] //导致 low 位置值为空
	for low < high {
		//high指针值 >= pivot high指针👈移
		for low < high && pivot <= data[high] {
			high--
		}
		//填补low位置空值
		//high指针值 < pivot high值 移到low位置
		//high 位置值空
		data[low] = data[high]
		//low指针值 <= pivot low指针👉移
		for low < high && pivot >= data[low] {
			low++
		}
		//填补high位置空值
		//low指针值 > pivot low值 移到high位置
		//low位置值空
		data[high] = data[low]
	}
	//pivot 填补 low位置的空值
	data[low] = pivot
	return low
}

func QuickSort(data []int, low, high int) {
	if high > low {
		//位置划分
		pivotIndex := partition(data, low, high)
		//左边部分排序
		QuickSort(data, low, pivotIndex-1)
		//右边排序
		QuickSort(data, pivotIndex+1, high)
	}
}

func TestQuickSort(t *testing.T) {
	data := []int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9}
	QuickSort(data, 0, len(data)-1)
	t.Log(data)
}
