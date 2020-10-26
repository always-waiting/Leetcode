package sort

import (
	"fmt"
)

func test_quickSort() {
	fmt.Println("testing")
}

/*
// 原始开发
func QuickSort(a []int, left, right, idx int) {
	if left == idx {
		return
	}
	if left > right {
		return
	}
	nIdx := axisCmp(a, left, right, idx)
	QuickSort(a, left, nIdx-2, nIdx-1)
	QuickSort(a, nIdx+1, right, idx)
}

func axisCmp(a []int, left, right, idx int) int {
	axis := a[idx]
	i := left
	j := right
	for i < j {
		if a[i] > axis && a[j] < axis {
			swap(a, i, j)
			i++
			j--
		} else if a[i] > axis {
			j--
		} else if a[j] < axis {
			i++
		} else {
			i++
			j--
		}
	}
	ret := 0
	if a[i] > axis {
		swap(a, i, idx)
		ret = i
	} else {
		swap(a, i+1, idx)
		ret = i + 1
	}
	return ret
}
*/
// 学习后开发
func QuickSort(a []int, left, right int) {
	if left >= right {
		return
	}
	idx := partition(a, left, right)
	QuickSort(a, left, idx-1)
	QuickSort(a, idx+1, right)
}

func partition(a []int, left, right int) int {
	pivot := a[right]
	i := left
	j := right - 1
	for i < j {
		for a[i] < pivot && i < right-1 {
			i++
		}
		for a[j] > pivot && j > 0 {
			j--
		}
		if i < j {
			swap(a, i, j)
		}
	}
	if a[i] < pivot {
		swap(a, i+1, right)
		return i + 1
	} else {
		swap(a, i, right)
		return i
	}
}

// 快速排序复习
func QuickSort2(a []int, left, right int) {
	if left >= right {
		return
	}
	idx := partition2(a, left, right)
	QuickSort2(a, left, idx-1)
	QuickSort2(a, idx+1, right)
}

func partition2(a []int, left, right int) int {
	pivot := a[right]
	i := left
	j := right - 1
	for i <= j {
		for i <= right-1 {
			if a[i] > pivot {
				break
			}
			i++
		}
		for j >= left {
			if a[j] < pivot {
				break
			}
			j--
		}
		if i < j {
			swap(a, i, j)
		}
	}
	var ret int
	if a[i] < pivot {
		swap(a, i+1, right)
		ret = i + 1
	} else {
		swap(a, i, right)
		ret = i
	}
	return ret

}

// 双轴快排
func DualPivotQuickSort(a []int, left, right int) {
	if left >= right {
		return
	}
	if left == right-1 {
		if a[left] > a[right] {
			swap(a, left, right)
		}
		return
	}
	idxs := dualPartition(a, left, right)
	DualPivotQuickSort(a, left, idxs[0]-1)
	DualPivotQuickSort(a, idxs[0]+1, idxs[1]-1)
	DualPivotQuickSort(a, idxs[1]+1, right)
}

func dualPartition(a []int, left, right int) []int {
	ret := make([]int, 2)
	if a[right] < a[right-1] {
		swap(a, right, right-1)
	}
	pivot1, pivot2 := a[right-1], a[right]
	i := left
	j := right - 2
	for i <= j {
		for i <= right-2 {
			if a[i] > pivot1 {
				break
			}
			i++
		}
		for j >= left {
			if a[j] < pivot1 {
				break
			}
			j--
		}
		if i <= j {
			swap(a, i, j)
		}
	}
	if a[i] > pivot1 {
		swap(a, i, right-1)
		ret[0] = i
	} else {
		if i+1 < right-1 {
			swap(a, i+1, right-1)
			ret[0] = i + 1
		} else {
			ret[0] = i
		}
	}
	i = ret[0]
	j = right - 1
	for i <= j {
		for i <= right-1 {
			if a[i] > pivot2 {
				break
			}
			i++
		}
		for j >= ret[0] {
			if a[j] < pivot2 {
				break
			}
			j--
		}
		if i <= j {
			swap(a, i, j)
		}
	}
	if a[i] >= pivot2 {
		swap(a, i, right)
		ret[1] = i
	} else {
		swap(a, i+1, right)
		ret[1] = i + 1
	}
	return ret
}
