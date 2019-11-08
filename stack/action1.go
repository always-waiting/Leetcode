package stack

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:
输入: "()"
输出: true

示例 2:
输入: "()[]{}"
输出: true

示例 3:
输入: "(]"
输出: false

示例 4:
输入: "([)]"
输出: false

示例 5:
输入: "{[]}"
输出: true
*/

func isValid(s string) bool {
	stack := StrArrStack{}
	for _, val := range s {
		if string(val) == " " {
			continue
		}
		if stack.IsEmpty() {
			stack.Push(string(val))
		} else {
			cmp := stack.Top() + string(val)
			if cmp != "()" && cmp != "{}" && cmp != "[]" {
				stack.Push(string(val))
			} else {
				stack.Pop()
			}
		}
	}
	if stack.IsEmpty() {
		return true
	}
	return false
}
