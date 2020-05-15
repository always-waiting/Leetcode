package array

/*
Contents
1. 买卖股票的最佳时机II[maxProfit]
2. 两数之和 II - 输入有序数组[twoSum]	--	★
3. 多数元素[majorityElement]	--	★
4. 旋转数组[rotate]
5. 存在重复元素[containsDuplicate]
6. 存在重复元素II[containsNearbyDuplicate]
7. 缺失数字[missingNumber]
8. 移动零[moveZeroes]
9. 第三大的数[thirdMax]
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

/*
多数元素

给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1:
输入: [3,2,3]
输出: 3

示例 2:
输入: [2,2,1,1,1,2,2]
输出: 2
*/
func majorityElement(nums []int) int {
	// 摩尔
	count := 1
	ret := nums[0]
	for _, val := range nums[1:] {
		if val != ret {
			count--
			if count == 0 {
				ret = val
				count++
			}
		} else {
			count++
		}
	}
	return ret
}

/*
旋转数组

给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

示例 1:
输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]

示例 2:
输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]

说明:
尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为O(1)的原地算法。
*/
func rotate(nums []int, k int) {
	l := k % len(nums)
	if l == 0 {
		return
	}
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, l-1)
	reverse(nums, l, len(nums)-1)
	/*
		num := len(nums)
		tmp := make([]int, l)
		copy(tmp, nums[num-l:])
		for i := num - l - 1; i >= 0; i-- {
			nums[i+l] = nums[i]
		}
		for i := 0; i < l; i++ {
			nums[i] = tmp[i]
		}
	*/
}

func reverse(nums []int, start, end int) {
	for start < end {
		tmp := nums[start]
		nums[start] = nums[end]
		nums[end] = tmp
		start++
		end--
	}
}

/*
存在重复元素

给定一个整数数组，判断是否存在重复元素。
如果任意一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。

示例 1:
输入: [1,2,3,1]
输出: true

示例 2:
输入: [1,2,3,4]
输出: false

示例 3:
输入: [1,1,1,3,3,4,3,2,4,2]
输出: true
*/
func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

/*
存在重复元素II

给定一个整数数组和一个整数k，判断数组中是否存在两个不同的索引i和j，使得nums[i] = nums[j]，并且i和j的差的绝对值至多为k。

示例 1:
输入: nums = [1,2,3,1], k = 3
输出: true

示例 2:
输入: nums = [1,0,1,1], k = 1
输出: true

示例 3:
输入: nums = [1,2,3,1,2,3], k = 2
输出: false
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && j-i <= k {
				return true
			}
		}
	}
	return false
}

/*
缺失数字

给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。

示例 1:
输入: [3,0,1]
输出: 2

示例 2:
输入: [9,6,4,2,3,5,7,0,1]
输出: 8

说明:
你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?
*/
func missingNumber(nums []int) int {
	n := len(nums)
	sum := 0
	for _, val := range nums {
		sum = sum + val
	}
	exp := (n + 1) * n / 2
	return exp - sum
}

/*
移动零

给定一个数组nums，编写一个函数将所有0移动到数组的末尾，同时保持非零元素的相对顺序。

示例:
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]

说明:
必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
*/
func moveZeroes(nums []int) {
	/*
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				for j := i + 1; j < len(nums); j++ {
					if nums[j] != 0 {
						nums[i] = nums[j]
						nums[j] = 0
						break
					}
				}
			}
		}
	*/
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for j < len(nums) {
		nums[j] = 0
		j++
	}
}

/*
第三大的数

给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。

示例 1:
输入: [3, 2, 1]
输出: 1
解释: 第三大的数是 1.

示例 2:
输入: [1, 2]
输出: 2
解释: 第三大的数不存在, 所以返回最大的数 2 .

示例 3:
输入: [2, 2, 3, 1]
输出: 1
解释: 注意，要求返回第三大的数，是指第三大且唯一出现的数。
存在两个值为2的数，它们都排第二。
*/
func thirdMax(nums []int) int {
	stack := []int{}
	for _, val := range nums {
		if len(stack) == 0 {
			stack = append(stack, val)
		} else if len(stack) == 1 {
			if stack[0] < val {
				stack = append(stack, stack[0])
				stack[0] = val
			} else if stack[0] == val {
			} else {
				stack = append(stack, val)
			}
		} else if len(stack) == 2 {
			if stack[0] < val {
				stack = append(stack, stack[1])
				stack[1] = stack[0]
				stack[0] = val
			} else if stack[0] == val {
			} else if stack[1] < val {
				stack = append(stack, stack[1])
				stack[1] = val
			} else if stack[1] == val {
			} else {
				stack = append(stack, val)
			}
		} else {
			if stack[0] < val {
				stack[2] = stack[1]
				stack[1] = stack[0]
				stack[0] = val
			} else if stack[0] == val {
			} else if stack[1] < val {
				stack[2] = stack[1]
				stack[1] = val
			} else if stack[1] == val {
			} else if stack[2] < val {
				stack[2] = val
			}
		}
	}
	if len(stack) == 3 {
		return stack[2]
	}
	return stack[0]
}
