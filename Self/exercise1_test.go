package Self

import "testing"

func TestMergeKArray(t *testing.T) {
	input := [][]int{
		[]int{1, 3, 5, 7},
		[]int{2, 4, 6},
		[]int{0, 8, 9, 10, 11},
	}
	ret := mergeKArray(input, 3)
	t.Log(ret)
}
