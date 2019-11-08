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

/*
有效括号字符串为空 ("")、"(" + A + ")" 或 A + B，其中 A 和 B 都是有效的括号字符串，+ 代表字符串的连接。例如，""，"()"，"(())()" 和 "(()(()))" 都是有效的括号字符串。
如果有效字符串 S 非空，且不存在将其拆分为 S = A+B 的方法，我们称其为原语（primitive），其中 A 和 B 都是非空有效括号字符串。
给出一个非空有效字符串 S，考虑将其进行原语化分解，使得：S = P_1 + P_2 + ... + P_k，其中 P_i 是有效括号字符串原语。
对 S 进行原语化分解，删除分解中每个原语字符串的最外层括号，返回 S 。

示例 1：
输入："(()())(())"
输出："()()()"
解释：
输入字符串为 "(()())(())"，原语化分解得到 "(()())" + "(())"，
删除每个部分中的最外层括号后得到 "()()" + "()" = "()()()"。

示例 2：
输入："(()())(())(()(()))"
输出："()()()()(())"
解释：
输入字符串为 "(()())(())(()(()))"，原语化分解得到 "(()())" + "(())" + "(()(()))"，
删除每隔部分中的最外层括号后得到 "()()" + "()" + "()(())" = "()()()()(())"。

示例 3：
输入："()()"
输出：""
解释：
输入字符串为 "()()"，原语化分解得到 "()" + "()"，
删除每个部分中的最外层括号后得到 "" + "" = ""。

提示：
S.length <= 10000
S[i] 为 "(" 或 ")"
S 是一个有效括号字符串
*/
func removeOuterParentheses(S string) string {
	//items := make([]string, 0)
	stack := Stack{}
	idxMap := make(map[int]bool, 0)
	for idx, runeVal := range S {
		val := string(runeVal)
		if stack.IsEmpty() {
			stack.Push(val)
			idxMap[idx] = true
		} else {
			node := stack.Top()
			cmp := node.String() + val
			if cmp != "{}" && cmp != "[]" && cmp != "()" {
				stack.Push(val)
			} else {
				stack.Pop()
				if stack.IsEmpty() {
					idxMap[idx] = true
				}
			}
		}
	}
	rets := make([]rune, 0)
	for idx, val := range S {
		if _, ok := idxMap[idx]; !ok {
			rets = append(rets, val)
		}
	}
	return string(rets)
}

func removeOuterParentheses1(S string) string {
	num := 0
	ret := make([]rune, 0)
	for _, runeVal := range S {
		strVal := string(runeVal)
		if strVal == "(" {
			if num > 0 {
				ret = append(ret, runeVal)
			}
			num++
		} else {
			num--
			if num > 0 {
				ret = append(ret, runeVal)
			}
		}

	}
	return string(ret)
}
