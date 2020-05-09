package array

/*
Contents
1. 买卖股票的最佳时机II[maxProfit]
2. 两数之和 II - 输入有序数组[twoSum]	--	★
*/

/*
买卖股票的最佳时机II

给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:
输入: [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。

示例 2:
输入: [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。

示例 3:
输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

提示：
1 <= prices.length <= 3 * 10 ^ 4
0 <= prices[i] <= 10 ^ 4
*/
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	inPrice := prices[0]
	lastPrice := prices[0]
	var ret int
	for _, val := range prices[1:] {
		if lastPrice > val {
			ret = ret + lastPrice - inPrice
			inPrice = val
		}
		lastPrice = val
	}
	ret = ret + lastPrice - inPrice
	return ret
}

/*
两数之和 II - 输入有序数组

给定一个已按照升序排列的有序数组，找到两个数使得它们相加之和等于目标数。
函数应该返回这两个下标值index1和index2，其中index1必须小于index2。

说明:
返回的下标值（index1 和 index2）不是从零开始的。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。

示例:
输入: numbers = [2, 7, 11, 15], target = 9
输出: [1,2]
解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
*/
func twoSum(numbers []int, target int) []int {
	/*
		// 速度慢!
		ret := make([]int, 0)
		for i, vali := range numbers {
			for j, valj := range numbers[i+1:] {
				if vali+valj == target {
					ret = append(ret, i+1, i+j+2)
					return ret
				}
			}
		}
		return ret
	*/
	// 双指针
	start := 0
	end := len(numbers) - 1
	for start < end {
		if numbers[start]+numbers[end] == target {
			return []int{start + 1, end + 1}
		} else if numbers[start]+numbers[end] < target {
			start++
		} else {
			end--
		}
	}
	return []int{}
}
