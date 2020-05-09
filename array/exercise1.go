package array

import (
	"fmt"
)

func fun() {
	fmt.Println("测试")
}

/*
Contents:
1. 两数之和[twoSum]
2. 删除排序数组中的重复项[removeDuplicates]	--	★
3. 移除元素[removeElement]
4. 搜索插入位置[searchInsert]
5. 最大子序和[maxSumArray]	--	★★★
6. 加一[plusOne]
7. 合并两个有序数组[merge]	--	★
8. 杨辉三角[generate]
9. 杨辉三角II[getRow]
10.买卖股票的最佳时机[maxProfit]	--	★
*/

/*
两数之和

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/
func twoSum(nums []int, target int) []int {
	ret := make([]int, 0)
	if len(nums) == 0 {
		return ret
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ret = append(ret, i, j)
			}
		}
	}
	return ret
}

/*
删除排序数组中的重复项

给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用O(1)额外空间的条件下完成。

示例 1:
给定数组 nums = [1,1,2],
函数应该返回新的长度 2, 并且原数组nums的前两个元素被修改为 1, 2。
你不需要考虑数组中超出新长度后面的元素。

示例 2:
给定 nums = [0,0,1,1,1,2,2,3,3,4],
函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
你不需要考虑数组中超出新长度后面的元素。

说明:
为什么返回数值是整数，但输出的答案是数组呢?
请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
你可以想象内部操作如下:
// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);
// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
*/
func removeDuplicates(nums []int) int {
	// 双指针!!
	if nums == nil || len(nums) == 1 {
		return 1
	}
	i := 0
	j := 1
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			i++
			nums[i] = nums[j]
			j++
		}
	}
	return i + 1
}

/*
移除元素

给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例 1:
给定 nums = [3,2,2,3], val = 3,
函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
你不需要考虑数组中超出新长度后面的元素。

示例 2:
给定 nums = [0,1,2,2,3,0,4,2], val = 2,
函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
注意这五个元素可为任意顺序。
你不需要考虑数组中超出新长度后面的元素。

说明:
为什么返回数值是整数，但输出的答案是数组呢?
请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
你可以想象内部操作如下:
// nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
int len = removeElement(nums, val);
// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
*/
func removeElement(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	i := 0
	j := 1
	for j < len(nums) {
		if nums[j] == val {
			j++
		} else {
			i++
			nums[i] = nums[j]
			j++
		}
	}
	if nums[0] == val {
		for i := 1; i < len(nums); i++ {
			nums[i-1] = nums[i]
		}
	} else {
		i++
	}
	return i
}

/*
搜索插入位置

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
你可以假设数组中无重复元素。

示例 1:
输入: [1,3,5,6], 5
输出: 2

示例 2:
输入: [1,3,5,6], 2
输出: 1

示例 3:
输入: [1,3,5,6], 7
输出: 4

示例 4:
输入: [1,3,5,6], 0
输出: 0
*/
func searchInsert(nums []int, target int) int {
	if nums[0] >= target {
		return 0
	}
	j := 1
	for j < len(nums) {
		if nums[j] == target {
			return j
		}
		if nums[j-1] < target && nums[j] > target {
			return j
		}
		j++
	}
	return j
}

/*
最大子序和

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:
输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:
如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
*/
func maxSubArray(nums []int) int { // Kadane算法
	maxEndHere := nums[0]
	maxSoFar := nums[0]
	for _, val := range nums[1:] {
		tmp := maxEndHere + val
		if val > tmp {
			maxEndHere = val
		} else {
			maxEndHere = tmp
		}
		if maxSoFar < maxEndHere {
			maxSoFar = maxEndHere
		}
	}
	return maxSoFar
}

/*
加一

给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:
输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。

示例 2:
输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
*/
func plusOne(digits []int) []int {
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		var tmp int
		if i == len(digits)-1 {
			tmp = digits[i] + 1 + carry
		} else {
			tmp = digits[i] + carry
		}
		a := tmp % 10
		digits[i] = a
		carry = tmp / 10
		if carry == 0 {
			break
		}
	}
	if carry != 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}

/*
合并两个有序数组

给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

说明:
初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

示例:
输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3
输出: [1,2,2,3,5,6]
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	last := m + n - 1
	for n != 0 {
		if m == 0 {
			nums1[last] = nums2[n-1]
			n--
			last--
			continue
		}
		if nums1[m-1] > nums2[n-1] {
			nums1[last] = nums1[m-1]
			m--
		} else {
			nums1[last] = nums2[n-1]
			n--
		}
		last--
	}
}

/*
杨辉三角

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:
输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/
func generate(numRows int) [][]int {
	ret := make([][]int, numRows)
	level := 0
	for level != numRows {
		if level == 0 {
			ret[level] = []int{1}
			level++
			continue
		}
		ret[level] = make([]int, level+1)
		for i := 0; i < level+1; i++ {
			if i == 0 || i == level {
				ret[level][i] = 1
			} else {
				ret[level][i] = ret[level-1][i-1] + ret[level-1][i]
			}
		}
		level++
	}
	return ret
}

/*
杨辉三角II
在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 3
输出: [1,3,3,1]
进阶：

你可以优化你的算法到 O(k) 空间复杂度吗？
*/
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	ret := make([]int, rowIndex+1)
	var tmp int
	for rowIndex >= 0 {
		for i := 0; i < len(ret)-rowIndex; i++ {
			if i == 0 || i == len(ret)-rowIndex-1 {
				ret[i] = 1
				tmp = ret[i]
			} else {
				a := tmp + ret[i]
				tmp = ret[i]
				ret[i] = a
			}
		}
		rowIndex--
	}
	return ret
}

/*
买卖股票的最佳时机

给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
注意：你不能在买入股票前卖出股票。

示例 1:
输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。

示例 2:
输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
[2,1,2,1,0,1,2]
*/
func maxProfit(prices []int) int {
	/*
		动态规划思想!!!!!
		1. 记录【今天之前买入的最小值】
		2. 计算【今天之前最小值买入，今天卖出的获利】，也即【今天卖出的最大获利】
		3. 比较【每天的最大获利】，取最大值即可
	*/
	if len(prices) <= 1 {
		return 0
	}
	inPrice := prices[0]
	ret := 0
	for _, val := range prices[1:] {
		if inPrice > val {
			inPrice = val
		} else {
			a := val - inPrice
			if ret < a {
				ret = a
			}
		}
	}
	return ret
}
