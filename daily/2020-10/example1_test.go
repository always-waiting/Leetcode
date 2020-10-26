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
