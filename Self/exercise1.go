package Self

import (
	"fmt"
)

/*
1. 合并k个有序数组 -- mergeKArray
2. 最长不重复字串
3. 字符串的前缀树
4. 目标和
5. 在每个树行中找最大值		--	https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/submissions/
*/

/*
创建一个大小为N的数组保存最后的结果
数组本身已经从小到大排好序，所以我们只需创建一个大小为k的最小堆，堆中初始元素为k个数组中的每个数组的第一个元素，
每次从堆中取出最小元素（堆顶元素），并将其存入输出数组中，将堆顶元素所在行的下一元素加入堆，重新排列出堆顶元素，
时间复杂度为logk，总共N个元素，所以总体时间复杂度是Nlogk
*/

type Ele struct {
	val      int
	idx      int
	arrayIdx int
}

type Stack []Ele

func (this *Stack) Add(e Ele) {
	if len(*this) == 0 {
		*this = append(*this, e)
	} else {
		insertIdx := -1
		for i := 0; i < len(*this); i++ {
			if (*this)[i].val > e.val {
				insertIdx = i
				break
			}
		}
		if insertIdx == 0 {
			tmp := append([]Ele{e}, *this...)
			*this = tmp
		} else if insertIdx == -1 {
			*this = append(*this, e)
		} else {
			*this = append(*this, Ele{})
			for i := len(*this) - 1; i > insertIdx; i-- {
				(*this)[i] = (*this)[i-1]
			}
			(*this)[insertIdx] = e
		}
	}
}

func (this *Stack) Pop() (e Ele) {
	e = (*this)[0]
	if len(*this) == 1 {
		*this = make([]Ele, 0)
	} else {
		*this = (*this)[1:]
	}
	return
}

func mergeKArray(input [][]int, k int) (ret []int) {
	stack := Stack{}
	ret = make([]int, 0)
	for i, arr := range input {
		e := Ele{
			val:      arr[0],
			idx:      0,
			arrayIdx: i,
		}
		stack.Add(e)
	}
	for len(stack) != 0 {
		val := stack.Pop()
		ret = append(ret, val.val)
		if val.idx < len(input[val.arrayIdx])-1 {
			newE := Ele{
				val:      input[val.arrayIdx][val.idx+1],
				arrayIdx: val.arrayIdx,
				idx:      val.idx + 1,
			}
			stack.Add(newE)
		}

	}
	return
}

/*
查找一个字串的最大不重复字串
*/

func longestSubStr(s string) (ret string) {
	if len(s) == 0 {
		return ""
	}
	seenMap := map[rune]int{}
	var start, end int
	for i, val := range s {
		if idx, ok := seenMap[val]; ok {
			if end-start > len(ret) {
				ret = s[start:end]
			}
			seenMap[val] = i
			if start <= idx {
				start = idx + 1
			}
		} else {
			seenMap[val] = i
		}
		end++
	}
	if end-start > len(ret) {
		ret = s[start:end]
	}
	return ret
}

/*
字符串的前缀树
*/
type TrieNode struct {
	val  string
	next []*TrieNode
}

func newTrieNode() *TrieNode {
	ret := &TrieNode{}
	ret.next = make([]*TrieNode, 0)
	return ret
}

func (this *TrieNode) Add(s string) {
	if len(s) == 0 {
		return
	}
	s0 := string(s[0])
	var has bool
	for _, n := range this.next {
		if n.val == s0 {
			has = true
			n.Add(s[1:])
		}
	}
	if !has {
		n := &TrieNode{val: s0, next: []*TrieNode{}}
		n.Add(s[1:])
		this.next = append(this.next, n)
	}
}

func (this *TrieNode) print() {
	fmt.Println(this.val)
	for _, n := range this.next {
		n.print()
	}
}

func (this *TrieNode) search(s string) (ret []string) {
	ret = make([]string, 0)
	if len(s) == 0 {
		if len(this.next) == 0 {
			return []string{""}
		}
		for _, n := range this.next {
			tmp := n.search(s)
			for _, tmpStr := range tmp {
				ret = append(ret, n.val+tmpStr)
			}
		}
		return ret
	}
	s0 := string(s[0])
	for _, n := range this.next {
		if n.val != s0 {
			continue
		}
		tmp := n.search(s[1:])
		for _, tmpStr := range tmp {
			ret = append(ret, s0+tmpStr)
		}
	}
	return ret
}

/*
4. 给出一个数组和一个target，返回所有和为target的组合

示例1:
输入: [7,2,3,4,5] target = 9
输出: [7,2], [2,3,4], [4,5]
*/
func getSumList(in []int, target int) (ret [][]int) {
	ret = make([][]int, 0)
	for i := 0; i < len(in); i++ {
		if in[i] == target {
			ret = append(ret, []int{in[i]})
		}
	}
	for i := 0; i < len(in)-1; i++ {
		tmp := target - in[i]
		tmpRes := getSumList(in[i+1:], tmp)
		if tmpRes != nil {
			for _, val := range tmpRes {
				ret = append(ret, append([]int{in[i]}, val...))
			}
		}
	}
	if len(ret) == 0 {
		ret = nil
	}
	return
}

/*
5. 返回一个二叉树每层的最大值
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	levelVals := getLevels(root)
	ret := []int{}
	for i := 0; i < len(levelVals); i++ {
		max := getMax(levelVals[i])
		ret = append(ret, max)
	}
	return ret
}

func getMax(in []int) int {
	a := in[0]
	for i := 1; i < len(in); i++ {
		if a < in[i] {
			a = in[i]
		}
	}
	return a
}

func getLevels(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := [][]int{[]int{root.Val}}
	left := getLevels(root.Left)
	right := getLevels(root.Right)
	for i := 0; i < len(left); i++ {
		tmp := left[i]
		if i < len(right) {
			tmp = append(tmp, right[i]...)
		}
		ret = append(ret, tmp)
	}
	if len(left) < len(right) {
		for i := len(left); i < len(right); i++ {
			ret = append(ret, right[i])
		}
	}
	return ret
}
