package Self

import "testing"

func TestMergeKArray(t *testing.T) {
	input := [][]int{
		[]int{1, 3, 5, 7},
		[]int{2, 4, 6},
		[]int{0, 8, 9, 10, 11},
	}
	ret := mergeKArray(input, 3)
	t.Log(ret)
}

func Test_longestSubStr(t *testing.T) {
	t.Log(longestSubStr("aaabcdab"))
}

func Test_TrieNode(t *testing.T) {
	root := newTrieNode()
	root.Add("and")
	root.Add("as")
	root.Add("at")
	root.Add("cn")
	root.Add("com")
	root.Add("codcdme")
	t.Logf("%v", root.search("cod"))
}
