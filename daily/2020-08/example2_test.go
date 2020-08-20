package main

import (
	"testing"
)

func Test_outerTrees(t *testing.T) {
	t.Log(outerTrees([][]int{
		[]int{1, 1}, []int{2, 2}, []int{2, 0},
		[]int{2, 4}, []int{3, 3}, []int{4, 2},
	}))
	t.Log(outerTrees([][]int{
		[]int{1, 2}, []int{2, 2}, []int{4, 2},
	}))
	t.Log(outerTrees([][]int{
		[]int{1, 2}, []int{2, 2}, []int{4, 2},
		[]int{5, 2}, []int{6, 2}, []int{7, 2},
	}))
}

func Test_updateBoard(t *testing.T) {
	board := [][]byte{
		[]byte{'E', 'E', 'E', 'E', 'E'},
		[]byte{'E', 'E', 'M', 'E', 'E'},
		[]byte{'E', 'E', 'E', 'E', 'E'},
		[]byte{'E', 'E', 'E', 'E', 'E'},
	}
	t.Log(updateBoard(board, []int{3, 0}))
}
