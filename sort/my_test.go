package sort

import (
	"reflect"
	"testing"
)

func Test_SelectSort(t *testing.T) {
	a := []int{9, 4, 5, 6, 1, 3, 7, 8, 2}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	SelectSort(a)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("选择排序返回错误, got(%v), expect(%v)", a, b)
	}
}

func Test_BubbleSort(t *testing.T) {
	a := []int{9, 4, 5, 6, 1, 3, 7, 8, 2}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	BubbleSort(a)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("冒泡排序返回错误, got(%v), expect(%v)", a, b)
	}
}

func Test_InsertSort(t *testing.T) {
	a := []int{9, 6, 1, 3, 5}
	b := []int{1, 3, 5, 6, 9}
	InsertSort(a)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("插入排序返回错误, got(%v), expect(%v)", a, b)
	}
}

func Test_ShellSort(t *testing.T) {
	a := []int{9, 6, 11, 3, 5, 12, 8, 7, 10, 15, 14, 4, 1, 13, 2}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	ShellSort(a)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("希尔排序返回错误, got(%v), expect(%v)", a, b)
	}
}

func Test_MergeSort(t *testing.T) {
	a := []int{1, 4, 6, 7, 10, 2, 3, 5, 8, 9}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a = MergeSort(a)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("归并排序返回错误, got(%v), expect(%v)", a, b)
	}
}

func Test_QuickSort(t *testing.T) {
	a := []int{1, 4, 6, 7, 10, 2, 3, 5, 8, 9}
	QuickSort2(a, 0, len(a)-1)
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("快速排序返回错误, got(%v), expect(%v)", a, b)
	}
}

func Test_DualPivotQuickSort(t *testing.T) {
	a := []int{12, 1, 9, 6, 7, 10, 2, 3, 5, 8, 4, 11}
	t.Log(a)
	DualPivotQuickSort(a, 0, len(a)-1)
	t.Log(a)
}

func Test_CountSort(t *testing.T) {
	a := []int{1, 4, 0, 6, 7, 10, 2, 3, 5, 8, 9}
	b := CountSort(a, 11)
	t.Log(a)
	t.Log(b)
}

func Test_CountSort1(t *testing.T) {
	a := []int{1, 4, 0, 6, 7, 10, 2, 3, 5, 8, 9}
	b := CountSort(a, 11)
	t.Log(a)
	t.Log(b)
}

func Test_BaseSort(t *testing.T) {
	a := []int{421, 240, 115, 532, 305, 430, 124}
	t.Log(a)
	a = BaseSort(a)
	t.Log(a)

}
