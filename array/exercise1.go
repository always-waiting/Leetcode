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
	/*
		count := 0
		loopNum := len(nums)
		for i := 0; i < loopNum; i++ {
			if i == 0 {
				count++
				continue
			}
			if nums[i-1] == nums[i] {
				k := i
				for _, val := range nums[i+1:] {
					nums[k] = val
					k++
				}
				i--
				loopNum--
			} else {
				count++
			}
			//fmt.Println(nums)
		}
		return count
	*/
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
