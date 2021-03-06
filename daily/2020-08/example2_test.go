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

func Test_repeatedSubstringPattern(t *testing.T) {
	t.Log(repeatedSubstringPattern("aabaab"))
}

func Test_findSubsequences(t *testing.T) {
	t.Log(findSubsequences([]int{4, 6, 7, 7}))
}

func Test_letterCombinations(t *testing.T) {
	t.Log(letterCombinations("23"))
}

func Test_findItinerary(t *testing.T) {
	t.Log(findItinerary([][]string{
		//[]string{"JFK", "KUL"}, []string{"JFK", "NRT"}, []string{"NRT", "JFK"},
		[]string{"JFK", "SFO"}, []string{"JFK", "ATL"}, []string{"SFO", "ATL"},
		[]string{"ATL", "JFK"}, []string{"ATL", "SFO"},
	}))
}

func Test_canVisitAllRooms(t *testing.T) {
	t.Log(canVisitAllRooms([][]int{
		[]int{1},
		[]int{2},
		[]int{3},
		[]int{},
	}))
	t.Log(canVisitAllRooms([][]int{
		[]int{1, 3},
		[]int{3, 0, 1},
		[]int{2},
		[]int{0},
	}))
}
