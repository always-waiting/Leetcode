package sort

import (
	"fmt"
)

func BaseSort(a []int) []int {
	maxBit := maxBit(a)
	fmt.Println(maxBit)
	for bit := 0; bit < maxBit; bit++ {
		a = baseSort(a, bit)
	}
	return a
}

func maxBit(a []int) int {
	maxBit := 1
	for _, i := range a {
		bit := 1
		for i > 9 {
			i = i / 10
			bit++
		}
		if bit > maxBit {
			maxBit = bit
		}
	}
	return maxBit
}

func baseSort(a []int, subNum int) []int {
	ret := make([]int, len(a))
	buckets := make([]int, 10)
	for _, num := range a {
		idx := getIdx(num, subNum)
		buckets[idx]++
	}

	for i := 1; i < len(buckets); i++ {
		buckets[i] = buckets[i] + buckets[i-1]
	}
	for i := len(a) - 1; i >= 0; i-- {
		idx := getIdx(a[i], subNum)
		ret[buckets[idx]-1] = a[i]
		buckets[idx]--
	}
	return ret
}

func getIdx(num int, subNum int) int {
	var idx int
	for i := 0; i <= subNum; i++ {
		idx = num % 10
		num = num / 10
	}
	return idx
}
