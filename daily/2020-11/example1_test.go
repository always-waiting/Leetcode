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

func Test_nextPermutation(t *testing.T) {
	nums := []int{1, 3, 2}
	nextPermutation(nums)
	t.Log(nums)
}

func Test_oddEvenList(t *testing.T) {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next = &ListNode{Val: 5}
	root.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	t.Log(root.Line())
	oddEvenList(root)
	t.Log(root.Line())
}

func Test_reconstructQueue(t *testing.T) {
	t.Log(reconstructQueue([][]int{
		[]int{7, 0}, []int{4, 4}, []int{7, 1},
		[]int{5, 0}, []int{6, 1}, []int{5, 2},
	}))
}

func Test_moveZeroes(t *testing.T) {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
	t.Log(a)
}
