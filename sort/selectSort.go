package sort

import (
	"fmt"
)

func SelectSort(a []int) {
	l := 0
	n := len(a) - 1
	for l < n {
		minPos := l
		maxPos := n
		for j := l; j <= n; j++ {
			if a[minPos] > a[j] {
				minPos = j
			}
		}
		for j := l; j < n; j++ {
			if a[maxPos] < a[j] {
				maxPos = j
			}
		}
		if maxPos == l && minPos == n {
			swap(a, minPos, maxPos)
		} else if maxPos == l {
			swap(a, n, maxPos)
			swap(a, l, minPos)
		} else {
			swap(a, minPos, l)
			swap(a, n, maxPos)
		}
		l++
		n--
	}
}

func swap(a []int, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func printArr(a []int) {
	fmt.Println(a)
}
