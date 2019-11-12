package stack

/*
根据逆波兰表示法，求表达式的值。
有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

说明：
整数除法只保留整数部分。
给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。

示例 1：
输入: ["2", "1", "+", "3", "*"]
输出: 9
解释: ((2 + 1) * 3) = 9

示例 2：
输入: ["4", "13", "5", "/", "+"]
输出: 6
解释: (4 + (13 / 5)) = 6

示例 3：
输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
输出: 22
解释:
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
*/
func evalRPN(tokens []string) int {
	stack := Stack{}
	for _, val := range tokens {
		var num int
		if val == "+" {
			num1 := stack.Pop().Int()
			num2 := stack.Pop().Int()
			num = num1 + num2
		} else if val == "-" {
			num1 := stack.Pop().Int()
			num2 := stack.Pop().Int()
			num = num2 - num1
		} else if val == "*" {
			num1 := stack.Pop().Int()
			num2 := stack.Pop().Int()
			num = num1 * num2
		} else if val == "/" {
			num1 := stack.Pop().Int()
			num2 := stack.Pop().Int()
			num = num2 / num1
		} else {
			num, _ = strconv.Atoi(val)
		}
		stack.Push(num)
	}
	return stack.Pop().Int()
}

/*
实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。
调用 next() 将返回二叉搜索树中的下一个最小的数。

示例：
	7
   / \
  3  15
    /  \
   9   20

BSTIterator iterator = new BSTIterator(root);
iterator.next();    // 返回 3
iterator.next();    // 返回 7
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 9
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 15
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 20
iterator.hasNext(); // 返回 false


提示：
next() 和 hasNext() 操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 中至少存在一个下一个最小的数。
*/

type BSTIterator struct {
	stack Stack
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{stack: Stack{}}
	cur := root
	for cur != nil {
		iter.stack.Push(cur)
		cur = cur.Left
	}
	return iter
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	node := this.stack.Pop()
	val := node.val.(*TreeNode)
	if val.Right != nil {
		cur := val.Right
		for cur != nil {
			this.stack.Push(cur)
			cur = cur.Left
		}
	}
	return val.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.stack.node != nil {
		return true
	}
	return false
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

/*
 给你一个字符串 s，「k 倍重复项删除操作」将会从 s 中选择 k 个相邻且相等的字母，并删除它们，使被删去的字符串的左侧和右侧连在一起。
你需要对 s 重复进行无限次这样的删除操作，直到无法继续为止。
在执行完所有删除操作后，返回最终得到的字符串。
本题答案保证唯一。

示例 1：
输入：s = "abcd", k = 2
输出："abcd"
解释：没有要删除的内容。

示例 2：
输入：s = "deeedbbcccbdaa", k = 3
输出："aa"
解释：
先删除 "eee" 和 "ccc"，得到 "ddbbbdaa"
再删除 "bbb"，得到 "dddaa"
最后删除 "ddd"，得到 "aa"

示例 3：
输入：s = "pbbcggttciiippooaais", k = 2
输出："ps"

提示：
1 <= s.length <= 10^5
2 <= k <= 10^4
s 中只含有小写英文字母。
*/
func removeDuplicates(s string, k int) string {
	stackS := Stack{}
	stackC := Stack{}
	var count int
	for _, val := range s {
		if stackS.IsEmpty() {
			stackS.Push(val)
			continue
		}
		if stackS.Top().Rune() == val {
			count++
		} else {
			stackC.Push(count)
			count = 0
		}
		stackS.Push(val)
		if count == k-1 {
			for i := 0; i < k; i++ {
				stackS.Pop()
			}
			tmp := stackC.Pop()
			if tmp != nil {
				count = tmp.Int()
			} else {
				count = 0
			}
		}
	}
	ret := ""
	for !stackS.IsEmpty() {
		val := stackS.Pop().Rune()
		ret = string(val) + ret
	}
	return ret
}
