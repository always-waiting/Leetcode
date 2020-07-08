package stack

/*
Contents:
1. 栈的最小值	--	★★
2. 滑动窗口的最大值[Not Finished]	--	★★★★
*/

/*
栈的最小值
请设计一个栈，除了常规栈支持的pop与push函数以外，还支持min函数，
该函数返回栈元素中的最小值。执行push、pop和min操作的时间复杂度必须为O(1)。

示例：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
*/
type MinStack struct { // 双栈使得pop和push能够以O(1)进行处理
	stack    []int
	minFuzhu []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minFuzhu: make([]int, 0),
	}

}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 || x <= this.minFuzhu[len(this.minFuzhu)-1] {
		this.minFuzhu = append(this.minFuzhu, x)
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	if this.stack[len(this.stack)-1] <= this.minFuzhu[len(this.minFuzhu)-1] {
		this.minFuzhu = this.minFuzhu[0 : len(this.minFuzhu)-1]
	}
	this.stack = this.stack[0 : len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]

}

func (this *MinStack) GetMin() int {
	return this.minFuzhu[len(this.minFuzhu)-1]
}

/*
滑动窗口的最大值

给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

示例:
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:
  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

提示：
你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
注意：本题与主站 239 题相同：https://leetcode-cn.com/problems/sliding-window-maximum/
*/
func maxSlidingWindow(nums []int, k int) []int {
	if k == 0 || nums == nil {
		return nil
	}
	if k == 1 {
		return nums
	}
	ret := make([]int, len(nums)-k+1)
	return ret
}
