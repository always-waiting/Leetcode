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

/*
给定一个由 '(' 和 ')' 括号组成的字符串 S，我们需要添加最少的括号（ '(' 或是 ')'，可以在任何位置），以使得到的括号字符串有效。
从形式上讲，只有满足下面几点之一，括号字符串才是有效的：
它是一个空字符串，或者
它可以被写成 AB （A 与 B 连接）, 其中 A 和 B 都是有效字符串，或者
它可以被写作 (A)，其中 A 是有效字符串。
给定一个括号字符串，返回为使结果字符串有效而必须添加的最少括号数。

示例 1：
输入："())"
输出：1

示例 2：
输入："((("
输出：3

示例 3：
输入："()"
输出：0

示例 4：
输入："()))(("
输出：4

提示：
S.length <= 1000
S 只包含 '(' 和 ')' 字符。
*/
func minAddToMakeValid(S string) int {
	items := make([]rune, 0)
	for _, val := range S {
		len := len(items)
		if len == 0 {
			items = append(items, val)
			continue
		}
		last := items[len-1]
		str := string([]rune{last, val})
		if str == "()" {
			items = items[:len-1]
		} else {
			items = append(items, val)
		}

	}
	return len(items)
}

/*
给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。
注意:
num 的长度小于 10002 且 ≥ k。
num 不会包含任何前导零。

示例 1 :
输入: num = "1432219", k = 3
输出: "1219"
解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。

示例 2 :
输入: num = "10200", k = 1
输出: "200"
解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。

示例 3 :
输入: num = "10", k = 2
输出: "0"
解释: 从原数字移除所有的数字，剩余为空就是0。
*/
func removeKdigits(num string, k int) string {
	items := make([]rune, 0)
	for _, val := range num {
		len := len(items)
		for len > 0 && items[len-1] > val && k > 0 {
			items = items[:len-1]
			k--
			len--
		}
		if val != 48 || len != 0 {
			items = append(items, val)
		}
	}
	if k > 0 && len(items)-k >= 0 {
		items = items[:len(items)-k]
	}
	if len(items) == 0 {
		return "0"
	}
	return string(items)
}
