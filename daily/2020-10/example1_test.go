package main

import (
	"testing"
)

func Test_reverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	t.Log(string(s))
}

func Test_sortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	t.Log(nums)
}

func Test_swapPairs(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	ret := swapPairs(head)
	t.Log(
		ret.Val,
		ret.Next.Val,
		ret.Next.Next.Val,
		ret.Next.Next.Next.Val,
	)
}

func Test_smallerNumbersThanCurrent(t *testing.T) {
	t.Log(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
}

func Test_longestMountain(t *testing.T) {
	t.Log(longestMountain([]int{2, 1, 4, 7, 3, 2, 5}))
	t.Log(longestMountain([]int{2, 2, 2, 2, 2}))
	t.Log(longestMountain([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
	t.Log(longestMountain([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}))
}

func Test_videoStitching(t *testing.T) {
	t.Log(
		videoStitching([][]int{
			[]int{0, 2}, []int{4, 6}, []int{8, 10},
			[]int{1, 9}, []int{1, 5}, []int{5, 9},
		}, 10),
	)
	t.Log(
		videoStitching([][]int{
			[]int{0, 1}, []int{1, 2},
		}, 5),
	)
	t.Log(
		videoStitching([][]int{
			[]int{0, 4}, []int{2, 8},
		}, 5),
	)
	t.Log(
		videoStitching([][]int{
			[]int{0, 1}, []int{6, 8}, []int{0, 2},
			[]int{5, 6}, []int{0, 4}, []int{0, 3},
			[]int{6, 7}, []int{1, 3}, []int{4, 7},
			[]int{1, 4}, []int{2, 5}, []int{2, 6},
			[]int{3, 4}, []int{4, 5}, []int{5, 7}, []int{6, 9}}, 9),
	)
}

func Test_partitionLabels(t *testing.T) {
	t.Log(partitionLabels("ababcbacadefegdehijhklij"))
}

func Test_uniqueOccurrences(t *testing.T) {
	t.Log(uniqueOccurrences([]int{1, 2}))
}

func Test_sumNumbers(t *testing.T) {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 1}

	t.Log(sumNumbers(root))
}
