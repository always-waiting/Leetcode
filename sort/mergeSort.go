package sort

func MergeSort(a []int) []int {
	if len(a) == 1 {
		return a
	} else if len(a) == 2 {
		if a[0] > a[1] {
			swap(a, 0, 1)
		}
		return a
	}
	halfIdx := len(a) / 2
	m := MergeSort(a[0:halfIdx])
	n := MergeSort(a[halfIdx:])
	idxM := 0
	idxN := 0
	b := make([]int, 0)
	for idxM < len(m) || idxN < len(n) {
		if idxM >= len(m) {
			b = append(b, n[idxN])
			idxN++
		} else if idxN >= len(n) {
			b = append(b, m[idxM])
			idxM++
		} else {
			if m[idxM] <= n[idxN] {
				b = append(b, m[idxM])
				idxM++
			} else {
				b = append(b, n[idxN])
				idxN++
			}
		}
	}
	return b
}

// 更优美的写法
func MergeSort2(a []int) []int {
	if len(a) == 1 {
		return a
	}
	ret := make([]int, len(a))
	mid := len(a) / 2
	left := MergeSort2(a[0:mid])
	right := MergeSort2(a[mid:])
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			ret[k] = left[i]
			i++
		} else {
			ret[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		ret[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		ret[k] = right[j]
		j++
		k++
	}
	return ret
}
