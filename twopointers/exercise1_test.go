package twopoints

import (
	"testing"
)

func TestSubarraysWithKDistinct(t *testing.T) {
	a := []int{1, 2, 1, 2, 3}
	k := 2
	ret := subarraysWithKDistinct(a, k)
	if ret != 7 {
		t.Errorf("Wrong answer: got(%d), expected(%d)", ret, 7)
	}
}
