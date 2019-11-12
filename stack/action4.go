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
		this.stack.Push(val.Right)
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
