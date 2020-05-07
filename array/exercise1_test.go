package array

import (
	"testing"
)

func TestExercise1(t *testing.T) {
	{
		t.Log("删除排序数组中的重复项......")
		a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		num := removeDuplicates(a)
		if num != 5 {
			t.Errorf("检查错误")
		} else {
			t.Log("检查正确")
		}
	}
}
