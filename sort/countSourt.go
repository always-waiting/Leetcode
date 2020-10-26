package sort

// 不稳定
func CountSort(a []int, k int) []int {
	ret := make([]int, len(a))
	buckets := make([]int, k)
	for _, i := range a {
		buckets[i]++
	}
	idx := 0
	for i, count := range buckets {
		for count != 0 {
			ret[idx] = i
			idx++
			count--
		}
	}
	return ret
}

// 稳定
func CountSort1(a []int, k int) []int {
	ret := make([]int, len(a))
	buckets := make([]int, k)
	for _, i := range a {
		buckets[i]++
	}
	for i := 1; i < len(buckets); i++ {
		buckets[i] = buckets[i] + buckets[i-1]
	}
	for i := len(a) - 1; i >= 0; i-- {
		ret[buckets[i]] = i
		buckets[i]--
	}
	return ret
}
