package sort

/*
// 第一次写的,同样正确
func BubbleSort(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		start := 0
		for j := 1; j <= i; j++ {
			if a[start] > a[j] {
				swap(a, start, j)
			}
			start = j
		}
	}
}
*/

func BubbleSort(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				swap(a, j, j+1)
			}
		}
	}
}
