package sort

/*
func ShellSort(a []int) {
	for gap := 4; gap > 0; gap = gap >> 1 {
		for m := 0; m < gap; m++ {
			for k := m; k < len(a); k = k + gap {
				for l := k; l >= gap; l = l - gap {
					if a[l] < a[l-gap] {
						swap(a, l, l-gap)
					}
				}
			}
		}
	}
}
*/
/*
// 采用了希尔序列
func ShellSort(a []int) {
	for gap := len(a)/2; gap > 0; gap = gap >> 1 {
		for i := gap; i < len(a); i++ {
			for j := i; j >= gap; j -= gap {
				if a[j] < a[j-gap] {
					swap(a, j, j-gap)
				}
			}
		}
	}
}
*/

func ShellSort(a []int) {
	h := 1
	for 3*h+1 <= len(a)/3 {
		h = 3*h + 1
	}
	for gap := h; gap > 0; gap = (gap - 1) / 3 {
		for i := gap; i < len(a); i++ {
			for j := i; j >= gap; j -= gap {
				if a[j] < a[j-gap] {
					swap(a, j, j-gap)
				}
			}
		}
	}
}

// 希尔排序即插入排序的改进(分步长的插入排序)
func ShellSort2(a []int) {
	for gap := 3; gap > 0; gap-- {
		for i := 0; i < len(a)-gap; i++ {
			for j := i + gap; j >= gap; j = j - gap {
				if a[j] < a[j-gap] {
					swap(a, j, j-gap)
				}
			}
		}
	}
}
