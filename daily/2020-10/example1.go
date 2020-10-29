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
7. 二叉树的前序遍历			--	https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
8. 数组中的最长山脉			--	https://leetcode-cn.com/problems/longest-mountain-in-array/
9. 视频拼接					--	https://leetcode-cn.com/problems/video-stitching/
10. 回文链表				--	https://leetcode-cn.com/problems/palindrome-linked-list/
11. 划分字母区间			--	https://leetcode-cn.com/problems/partition-labels/
12. 根到叶子节点数字之和	--	https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/
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

/*
144. 二叉树的前序遍历
给定一个二叉树，返回它的 前序 遍历。

示例:
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
*/
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := []int{root.Val}
	right := preorderTraversal(root.Right)
	left := preorderTraversal(root.Left)
	if left != nil {
		ret = append(ret, left...)
	}
	if right != nil {
		ret = append(ret, right...)
	}
	return ret
}

func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := []int{root.Val}
	stack := []*TreeNode{}
	if root.Right != nil {
		stack = append(stack, root.Right)
	}
	if root.Left != nil {
		stack = append(stack, root.Left)
	}
	for len(stack) != 0 {
		now := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		ret = append(ret, now.Val)
		if now.Right != nil {
			stack = append(stack, now.Right)
		}
		if now.Left != nil {
			stack = append(stack, now.Left)
		}
	}
	return ret
}

/*
845. 数组中的最长山脉
我们把数组A中符合下列属性的任意连续子数组B称为 “山脉”：
B.length >= 3
存在 0 < i < B.length - 1 使得 B[0] < B[1] < ... B[i-1] < B[i] > B[i+1] > ... > B[B.length - 1]
（注意：B 可以是 A 的任意子数组，包括整个数组 A。）
给出一个整数数组 A，返回最长 “山脉” 的长度。
如果不含有 “山脉” 则返回 0。

示例 1：
输入：[2,1,4,7,3,2,5]
输出：5
解释：最长的 “山脉” 是 [1,4,7,3,2]，长度为 5。
示例 2：
输入：[2,2,2]
输出：0
解释：不含 “山脉”。

提示：
0 <= A.length <= 10000
0 <= A[i] <= 10000
*/
func longestMountain(A []int) int {
	ret := 0
	var start, top, end, j int
	for j < len(A) {
		start = j
		top = j
		for i := start + 1; i < len(A); i++ {
			if A[i] > A[i-1] {
				top++
			} else {
				break
			}
		}
		end = top
		for i := top + 1; i < len(A); i++ {
			if A[i] < A[i-1] {
				end++
			} else {
				break
			}
		}
		if end <= j {
			j++
		} else {
			j = end
		}
		if end == top || start == top {
			continue
		}
		if ret < end-start+1 {
			ret = end - start + 1
		}
	}
	/*
		for j := 0; j < len(A); j++ {
			if j < top {
				continue
			} else {
				start = j
				top = j
			}
			for i := start + 1; i < len(A); i++ {
				if A[i] > A[i-1] {
					top++
				} else {
					break
				}
			}
			end = top
			for i := top + 1; i < len(A); i++ {
				if A[i] < A[i-1] {
					end++
				} else {
					break
				}
			}
			if end == top || start == top {
				continue
			}
			if ret < end-start+1 {
				ret = end - start + 1
			}
		}
	*/
	if ret < 3 {
		ret = 0
	}
	return ret
}

/*
1024. 视频拼接
你将会获得一系列视频片段，这些片段来自于一项持续时长为 T 秒的体育赛事。这些片段可能有所重叠，也可能长度不一。
视频片段 clips[i] 都用区间进行表示：开始于 clips[i][0] 并于 clips[i][1] 结束。我们甚至可以对这些片段自由地再剪辑，例如片段 [0, 7] 可以剪切成 [0, 1] + [1, 3] + [3, 7] 三部分。
我们需要将这些片段进行再剪辑，并将剪辑后的内容拼接成覆盖整个运动过程的片段（[0, T]）。返回所需片段的最小数目，如果无法完成该任务，则返回 -1 。

示例 1：
输入：clips = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]], T = 10
输出：3
解释：
我们选中 [0,2], [8,10], [1,9] 这三个片段。
然后，按下面的方案重制比赛片段：
将 [1,9] 再剪辑为 [1,2] + [2,8] + [8,9] 。
现在我们手上有 [0,2] + [2,8] + [8,10]，而这些涵盖了整场比赛 [0, 10]。
示例 2：
输入：clips = [[0,1],[1,2]], T = 5
输出：-1
解释：
我们无法只用 [0,1] 和 [1,2] 覆盖 [0,5] 的整个过程。
示例 3：
输入：clips = [[0,1],[6,8],[0,2],[5,6],[0,4],[0,3],[6,7],[1,3],[4,7],[1,4],[2,5],[2,6],[3,4],[4,5],[5,7],[6,9]], T = 9
输出：3
解释：
我们选取片段 [0,4], [4,7] 和 [6,9] 。
示例 4：
输入：clips = [[0,4],[2,8]], T = 5
输出：2
解释：
注意，你可能录制超过比赛结束时间的视频。

提示：
1 <= clips.length <= 100
0 <= clips[i][0] <= clips[i][1] <= 100
0 <= T <= 100
*/
func videoStitching(clips [][]int, T int) int {
	var now int
	picks := [][]int{}
	count := 0
	for now < T && count < len(clips) {
		clipick := []int{now, now}
		for _, clip := range clips {
			if clip[0] <= now && now < clip[1] {
				if clip[1]-now > clipick[1]-now {
					clipick = clip
				}
			}
		}
		picks = append(picks, clipick)
		now = clipick[1]
		count++
	}
	if now < T {
		return -1
	}
	return len(picks)
}

/*
234. 回文链表
请判断一个链表是否为回文链表。
示例 1:
输入: 1->2
输出: false
示例 2:
输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/
func isPalindrome(head *ListNode) bool {
	vals := []int{}
	now := head
	for now != nil {
		vals = append(vals, now.Val)
		now = now.Next
	}
	for i := 0; i < len(vals); i++ {
		if vals[i] != vals[len(vals)-1-i] {
			return false
		}
	}
	return true
}

/*
763. 划分字母区间
字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。

示例：
输入：S = "ababcbacadefegdehijhklij"
输出：[9,7,8]
解释：
划分结果为 "ababcbaca", "defegde", "hijhklij"。
每个字母最多出现在一个片段中。
像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。

提示：
S的长度在[1, 500]之间。
S只包含小写字母 'a' 到 'z' 。
*/
func partitionLabels(S string) []int {
	if len(S) == 0 {
		return nil
	}
	startVal := S[0]
	endIdx := 0
	for i := 1; i < len(S); i++ {
		if S[i] == startVal {
			endIdx = i
		}
	}
	for i := 1; i < endIdx; i++ {
		for j := endIdx + 1; j < len(S); j++ {
			if S[i] == S[j] {
				endIdx = j
			}
		}
	}
	ret := []int{endIdx + 1}
	if endIdx < len(S) {
		nS := S[endIdx+1:]
		next := partitionLabels(nS)
		if next != nil {
			ret = append(ret, next...)
		}
	}
	return ret
}

func partitionLabels1(S string) []int {
	visit := map[byte]int{}
	for i := 0; i < len(S); i++ {
		visit[S[i]] = i
	}
	var start, end int
	ret := []int{}
	for i := 0; i < len(S); i++ {
		if end < visit[S[i]] {
			end = visit[S[i]]
		}
		if i == end {
			ret = append(ret, end-start+1)
			start = end + 1
			end = start
		}
	}
	return ret
}

/*
1207. 独一无二的出现次数
给你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。
如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。

示例 1：
输入：arr = [1,2,2,1,1,3]
输出：true
解释：在该数组中，1 出现了 3 次，2 出现了 2 次，3 只出现了 1 次。没有两个数的出现次数相同。
示例 2：
输入：arr = [1,2]
输出：false
示例 3：
输入：arr = [-3,0,1,-3,1,1,1,-3,10,0]
输出：true

提示：
1 <= arr.length <= 1000
-1000 <= arr[i] <= 1000
*/
func uniqueOccurrences(arr []int) bool {
	count := map[int]int{}
	for _, val := range arr {
		count[val]++
	}
	uniq := map[int]bool{}
	for _, num := range count {
		if _, ok := uniq[num]; !ok {
			uniq[num] = true
		} else {
			return false
		}

	}
	return true
}

/*
129. 求根到叶子节点数字之和
给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。
例如，从根到叶子节点路径 1->2->3 代表数字 123。
计算从根到叶子节点生成的所有数字之和。
说明: 叶子节点是指没有子节点的节点。
示例 1:
输入: [1,2,3]
    1
   / \
  2   3
输出: 25
解释:
从根到叶子节点路径 1->2 代表数字 12.
从根到叶子节点路径 1->3 代表数字 13.
因此，数字总和 = 12 + 13 = 25.
示例 2:
输入: [4,9,0,5,1]
    4
   / \
  9   0
 / \
5   1
输出: 1026
解释:
从根到叶子节点路径 4->9->5 代表数字 495.
从根到叶子节点路径 4->9->1 代表数字 491.
从根到叶子节点路径 4->0 代表数字 40.
因此，数字总和 = 495 + 491 + 40 = 1026.
*/
func sumNumbers(root *TreeNode) int {
	var getNumbers func(*TreeNode, int) []int
	getNumbers = func(r *TreeNode, pre int) []int {
		if r == nil {
			return nil
		}
		ret := []int{}
		if r.Left != nil {
			lRet := getNumbers(r.Left, pre*10+r.Val)
			ret = append(ret, lRet...)
		}
		if r.Right != nil {
			rRet := getNumbers(r.Right, pre*10+r.Val)
			ret = append(ret, rRet...)
		}
		if len(ret) == 0 {
			ret = append(ret, pre*10+r.Val)
		}
		return ret
	}
	nums := getNumbers(root, 0)
	fmt.Println(nums)
	sum := 0
	for _, val := range nums {
		sum = sum + val
	}
	return sum
}
