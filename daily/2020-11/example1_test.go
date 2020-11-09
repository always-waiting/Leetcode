package main

import (
	"testing"
)

func Test_validMountainArray(t *testing.T) {
	t.Log(validMountainArray([]int{14, 82, 89, 84, 79, 70, 70, 68, 67, 66, 63, 60, 58, 54, 44, 43, 32, 28, 26, 25, 22, 15, 13, 12, 10, 8, 7, 5, 4, 3}))
}

func Test_insert(t *testing.T) {
	t.Log(insert(
		[][]int{[]int{1, 2}, []int{3, 5}, []int{6, 7}, []int{8, 10}, []int{12, 16}},
		[]int{4, 8},
	))
	t.Log(insert(
		[][]int{[]int{1, 2}, []int{5, 9}},
		[]int{7, 8},
	))
	t.Log(insert(
		[][]int{},
		[]int{5, 6},
	))
	t.Log(insert(
		[][]int{[]int{1, 5}},
		[]int{6, 8},
	))
}

func Test_ladderLength(t *testing.T) {
	t.Log(ladderLength("hit", "cog", []string{"hot", "dot", "lot", "log", "cog"}))
	t.Log(ladderLength("abc", "def", []string{"aec", "mec", "aef", "def"}))
}

func Test_kClosest(t *testing.T) {
	t.Log(kClosest(
		[][]int{
			[]int{1, 3}, []int{2, -2},
		},
		1,
	))
}
