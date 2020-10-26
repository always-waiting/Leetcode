package main

import (
	"fmt"
	"math"
)

/*
1. 环形链表		--	https://leetcode-cn.com/problems/linked-list-cycle/
2. 反转字符串	--	https://leetcode-cn.com/problems/reverse-string/
3. 颜色分类		--	https://leetcode-cn.com/problems/sort-colors/
4. 环形链表IIi	--	https://leetcode-cn.com/problems/linked-list-cycle-ii/
5. 两两交换链表中的节点		--	https://leetcode-cn.com/problems/swap-nodes-in-pairs/
6. 有多少小于当前数字的数字	--	https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/
*/
func test() {
	fmt.Println("testing")
}

/*
环形链表
给定一个链表，判断链表中是否有环。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
如果链表中存在环，则返回 true 。 否则，返回 false 。

进阶：
你能用 O(1)（即，常量）内存解决此问题吗？
示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。
示例 2：
输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。
示例 3：
输入：head = [1], pos = -1
输出：false
解释：链表中没有环。

提示：
链表中节点的数目范围是 [0, 104]
-105 <= Node.val <= 105
pos 为 -1 或者链表中的一个 有效索引 。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var fast, slow *ListNode
	fast, slow = head, head
	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

/*
反转字符串
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。
示例 1：
输入：["h","e","l","l","o"]
输出：["o","l","l","e","h"]
示例 2：
输入：["H","a","n","n","a","h"]
输出：["h","a","n","n","a","H"]
*/
func reverseString(s []byte) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}

/*
颜色分类
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
注意:
不能使用代码库中的排序函数来解决这道题。
示例:
输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]

进阶：
一个直观的解决方案是使用计数排序的两趟扫描算法。
首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
你能想出一个仅使用常数空间的一趟扫描算法吗？
*/
func sortColorsOld(nums []int) {
	count := map[int]int{}
	for _, val := range nums {
		count[val]++
	}
	idx := 0
	for _, val := range []int{0, 1, 2} {
		if n, ok := count[val]; !ok {
			continue
		} else {
			for n > 0 {
				nums[idx] = val
				idx++
				n--
			}
		}
	}
}

func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for ; i <= p2 && nums[i] == 2; p2-- {
			nums[i], nums[p2] = nums[p2], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}

/*
环形链表 II
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
说明：不允许修改给定的链表。
示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。
示例 2：
输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。
示例 3：
输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。

进阶：
你是否可以不用额外空间解决此题？
*/
func detectCycle(head *ListNode) *ListNode {
	var fast, slow *ListNode
	fast, slow = head, head
	for fast != nil && slow != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			slow = head
			for slow != nil && fast != nil {
				if slow == fast {
					return slow
				}
				slow = slow.Next
				fast = fast.Next
			}
		}
	}
	return nil
}

/*
530. 二叉搜索树的最小绝对差
给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。
示例：
输入：
   1
    \
     3
    /
   2
输出：
1
解释：
最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
提示：
树中至少有 2 个节点。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	preOrder := getInorder(root)
	minRet := preOrder[1] - preOrder[0]
	for i := 1; i < len(preOrder)-1; i++ {
		tmp := preOrder[i+1] - preOrder[i]
		if tmp < minRet {
			minRet = tmp
		}
	}
	return minRet
}

func getInorder(root *TreeNode) []int {
	var ret []int
	if root.Left != nil {
		ret = getInorder(root.Left)
	}
	ret = append(ret, root.Val)
	if root.Right != nil {
		tmp := getInorder(root.Right)
		ret = append(ret, tmp...)
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMinimumDifferenceWooooo(root *TreeNode) int {
	ans, pre := math.MaxInt64, -1
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

/*
24. 两两交换链表中的节点
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
示例 1：
输入：head = [1,2,3,4]
输出：[2,1,4,3]
示例 2：
输入：head = []
输出：[]
示例 3：
输入：head = [1]
输出：[1]
提示：
链表中节点的数目在范围 [0, 100] 内
0 <= Node.val <= 100
*/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var a, b, pre, ret *ListNode
	a, b = head, head.Next
	for a != nil && b != nil {
		tmp := b.Next
		b.Next = a
		if pre != nil {
			pre.Next = b
		} else {
			ret = b
		}
		pre = a
		a = tmp
		if a != nil {
			b = a.Next
		}
	}
	pre.Next = a
	return ret
}

func swapPairsBetter(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}

/*
1365. 有多少小于当前数字的数字
给你一个数组nums，对于其中每个元素nums[i]，请你统计数组中比它小的所有数字的数目。
换而言之，对于每个 nums[i] 你必须计算出有效的j的数量，其中j满足 j != i 且 nums[j] < nums[i] 。
以数组形式返回答案。

示例 1：
输入：nums = [8,1,2,2,3]
输出：[4,0,1,1,3]
解释：
对于 nums[0]=8 存在四个比它小的数字：（1，2，2 和 3）。
对于 nums[1]=1 不存在比它小的数字。
对于 nums[2]=2 存在一个比它小的数字：（1）。
对于 nums[3]=2 存在一个比它小的数字：（1）。
对于 nums[4]=3 存在三个比它小的数字：（1，2 和 2）。

示例 2：
输入：nums = [6,5,4,8]
输出：[2,1,0,3]

示例 3：
输入：nums = [7,7,7,7]
输出：[0,0,0,0]

提示：
2 <= nums.length <= 500
0 <= nums[i] <= 100
*/

func smallerNumbersThanCurrent(nums []int) []int {
	maxNum := maxArr(nums)
	buckets := make([]int, maxNum+1)
	for _, num := range nums {
		buckets[num]++
	}
	cumBuckets := make([]int, maxNum+1)
	for i := 1; i < len(buckets); i++ {
		cumBuckets[i] = cumBuckets[i-1] + buckets[i-1]
	}
	ret := make([]int, len(nums))
	for i, num := range nums {
		ret[i] = cumBuckets[num]
	}
	return ret
}

func maxArr(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}
