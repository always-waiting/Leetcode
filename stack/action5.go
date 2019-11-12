package stack

/*
给你一个由 '('、')' 和小写字母组成的字符串 s。
你需要从字符串中删除最少数目的 '(' 或者 ')' （可以删除任意位置的括号)，使得剩下的「括号字符串」有效。
请返回任意一个合法字符串。
有效「括号字符串」应当符合以下 任意一条 要求：
空字符串或只包含小写字母的字符串
可以被写作 AB（A 连接 B）的字符串，其中 A 和 B 都是有效「括号字符串」
可以被写作 (A) 的字符串，其中 A 是一个有效的「括号字符串」

示例 1：
输入：s = "lee(t(c)o)de)"
输出："lee(t(c)o)de"
解释："lee(t(co)de)" , "lee(t(c)ode)" 也是一个可行答案。

示例 2：
输入：s = "a)b(c)d"
输出："ab(c)d"

示例 3：
输入：s = "))(("
输出：""
解释：空字符串也是有效的

示例 4：
输入：s = "(a(b(c)d)"
输出："a(b(c)d)"

提示：
1 <= s.length <= 10^5
s[i] 可能是 '('、')' 或英文小写字母
*/
func minRemoveToMakeValid(s string) string {
	removeIdx := make([]int, 0)
	rets := make([]rune, 0)
	runes := []rune("()")
	for _, val := range s {
		if val != runes[0] && val != runes[1] {
			rets = append(rets, val)
		} else if val == runes[0] {
			rets = append(rets, val)
			removeIdx = append(removeIdx, len(rets)-1)
		} else {
			if len(removeIdx) != 0 {
				removeIdx = removeIdx[:len(removeIdx)-1]
				rets = append(rets, val)
			}
		}
	}
	if len(removeIdx) != 0 {
		for i := len(removeIdx) - 1; i >= 0; i-- {
			rets = append(rets[:removeIdx[i]], rets[removeIdx[i]+1:]...)
		}
	}
	return string(rets)
}

/*
给出一个字符串 s（仅含有小写英文字母和括号）。
请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。

注意，您的结果中 不应 包含任何括号。

示例 1：
输入：s = "(abcd)"
输出："dcba"

示例 2：
输入：s = "(u(love)i)"
输出："iloveu"

示例 3：
输入：s = "(ed(et(oc))el)"
输出："leetcode"

示例 4：
输入：s = "a(bcdefghijkl(mno)p)q"
输出："apmnolkjihgfedcbq"

提示：
0 <= s.length <= 2000
s 中只有小写英文字母和括号
我们确保所有括号都是成对出现的
*/
func reverseParentheses(s string) string {
	rets := make([][]rune, 1)
	runes := []rune("()")
	idx := 0
	for _, val := range s {
		if val == runes[0] {
			idx++
			if len(rets) == idx {
				rets = append(rets, make([]rune, 0))
			} else {
				rets[idx] = []rune{}
			}
		} else if val == runes[1] {
			tmp := rets[idx]
			idx--
			if idx%2 == 1 {
				rets[idx] = append(tmp, rets[idx]...)
			} else {
				rets[idx] = append(rets[idx], tmp...)
			}
		} else {
			if idx%2 == 1 {
				rets[idx] = append([]rune{val}, rets[idx]...)
			} else {
				rets[idx] = append(rets[idx], val)
			}
		}
	}
	return string(rets[0])
}

/*
给你一份工作时间表 hours，上面记录着某一位员工每天的工作小时数。
我们认为当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」。
所谓「表现良好的时间段」，意味在这段时间内，「劳累的天数」是严格 大于「不劳累的天数」。
请你返回「表现良好时间段」的最大长度。

示例 1：
输入：hours = [9,9,6,0,6,6,9]
输出：3
解释：最长的表现良好时间段是 [9,9,6]。

提示：
1 <= hours.length <= 10000
0 <= hours[i] <= 16
*/
func longestWPI(hours []int) int {
	var h, l, total int
	for _, val := range hours {
		if val > 8 {
			h++
		} else {
			l++
		}
		if h > l {
			if total < h+l {
				total = h + l
			}
		} else {
			h = 0
			l = 0
		}
	}
	return total
}
