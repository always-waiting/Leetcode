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

/*
给定两个没有重复元素的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。
nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出-1。

示例 1:
输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
输出: [-1,3,-1]
解释:
    对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
    对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
    对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。

示例 2:
输入: nums1 = [2,4], nums2 = [1,2,3,4].
输出: [3,-1]
解释:
    对于num1中的数字2，第二个数组中的下一个较大数字是3。
    对于num1中的数字4，第二个数组中没有下一个更大的数字，因此输出 -1。

注意:
nums1和nums2中所有元素是唯一的。
nums1和nums2 的数组大小都不超过1000。
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ret := make([]int, 0)
	for idx, val := range nums1 {
		flag := false
		for _, com := range nums2 {
			if com == val {
				flag = true
			}
			if com > val && flag {
				ret = append(ret, com)
				break
			}
		}
		if len(ret) < idx+1 {
			ret = append(ret, -1)
		}
	}
	return ret
}

func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	ret := make([]int, 0)
	total := len(nums2)
	for _, val := range nums1 {
		num := -1
		for i := total - 1; i >= 0; i-- {
			if nums2[i] == val {
				break
			}
			if nums2[i] > val {
				num = nums2[i]
			}
		}
		ret = append(ret, num)
	}
	return ret
}
