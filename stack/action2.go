package stack

/*
给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
在 S 上反复执行重复项删除操作，直到无法继续删除。
在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

示例：
输入："abbaca"
输出："ca"
解释：
例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。之后我们得到字符串 "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。

提示：
1 <= S.length <= 20000
S 仅由小写英文字母组成。
*/
func removeDuplicates(S string) string {
	stack := Stack{}
	for _, val := range S {
		if stack.IsEmpty() {
			stack.Push(val)
			continue
		}
		flag := true
		for !stack.IsEmpty() && stack.Top().Rune() == val {
			flag = false
			stack.Pop()
		}
		if flag {
			stack.Push(val)
		}
	}
	ret := make([]rune, 0)
	for !stack.IsEmpty() {
		ret = append(ret, stack.Top().Rune())
		stack.Pop()
	}
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}
	return string(ret)
}

/*
给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。

示例 1：
输入：S = "ab#c", T = "ad#c"
输出：true
解释：S 和 T 都会变成 “ac”。

示例 2：
输入：S = "ab##", T = "c#d#"
输出：true
解释：S 和 T 都会变成 “”。

示例 3：
输入：S = "a##c", T = "#a#c"
输出：true
解释：S 和 T 都会变成 “c”。

示例 4：
输入：S = "a#c", T = "b"
输出：false
解释：S 会变成 “c”，但 T 仍然是 “b”。

提示：
1 <= S.length <= 200
1 <= T.length <= 200
S 和 T 只含有小写字母以及字符 '#'。
*/
func backspaceCompare(S string, T string) bool {
	stack1 := Stack{}
	stack2 := Stack{}
	for _, val := range S {
		if string(val) == "#" {
			stack1.Pop()
		} else {
			stack1.Push(val)
		}
	}
	for _, val := range T {
		if string(val) == "#" {
			stack2.Pop()
		} else {
			stack2.Push(val)
		}
	}
	for !stack1.IsEmpty() && !stack2.IsEmpty() {
		if stack1.Pop().Rune() != stack2.Pop().Rune() {
			return false
		}
	}
	if !stack1.IsEmpty() || !stack2.IsEmpty() {
		return false
	}
	return true
}
