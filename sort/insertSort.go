package sort

/*
样本小且基本有序时效率高
*/
func InsertSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				swap(a, j, j-1)
			}
		}
	}
}
