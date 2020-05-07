package stack

import (
	"reflect"
	"testing"
)

func TestExercise1(t *testing.T) {
	{
		t.Log("滑动窗口的最大值......")
		nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
		k := 3
		ret := maxSlidingWindow(nums, k)
		expected := []int{3, 3, 5, 5, 6, 7}
		if reflect.DeepEqual(ret, expected) {
			t.Log("检查正确")
		} else {
			t.Errorf("检查错误")
		}
	}
}
