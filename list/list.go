package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(data []int) *ListNode {
	var ret, loop *ListNode
	for _, val := range data {
		tmp := &ListNode{Val: val}
		if ret == nil {
			ret = tmp
			loop = tmp
			continue
		} else {
			loop.Next = tmp
			loop = tmp
		}
	}
	return ret
}

func newCycleList(data []int, pos int) *ListNode {
	if data == nil {
		return nil
	}
	if pos < 0 || pos >= len(data) {
		return newListNode(data)
	}
	var ret, loop, cycNode *ListNode
	for idx, val := range data {
		tmp := &ListNode{Val: val}
		if idx == pos {
			cycNode = tmp
		}
		if ret == nil {
			ret = tmp
			loop = tmp
			continue
		} else {
			loop.Next = tmp
			loop = tmp
		}
	}
	loop.Next = cycNode
	return ret
}

func newIntersectionList(dataA, dataB []int, interVal, sa, sb int) (*ListNode, *ListNode) {
	var retA, retB, loop, interNode *ListNode
	for idx, val := range dataA {
		tmp := &ListNode{Val: val}
		if retA == nil {
			retA = tmp
			loop = tmp
			continue
		} else {
			loop.Next = tmp
			loop = tmp
		}
		if idx == sa && val == interVal {
			interNode = loop
		}
	}
	for idx, val := range dataB {
		tmp := &ListNode{Val: val}
		if retB == nil {
			retB = tmp
			loop = tmp
			continue
		}
		if idx == sb && interNode != nil {
			loop.Next = interNode
			break
		} else {
			loop.Next = tmp
			loop = tmp
		}
	}
	return retA, retB
}

/*
会修改列表本身,this变为队列最后一个元素
*/
func (this *ListNode) reverseList1() *ListNode {
	if this == nil || this.Next == nil {
		return this
	}
	loop := this.Next
	this.Next = nil
	for {
		if loop == nil {
			break
		}
		tmp := loop.Next
		loop.Next = this
		this = loop
		loop = tmp
	}
	return this
}

/*
会修改列表本身,让本身反转,因此不用返回反转变量
循环过多
*/
func (this *ListNode) reverseList2() {
	if this == nil || this.Next == nil {
		return
	}
	total := 1
	loop := this
	for {
		if loop.Next == nil {
			break
		}
		loop = loop.Next
		total++
	}
	loopNum := total / 2
	for i := 0; i < loopNum; i++ {
		idx := 1
		var a, b *ListNode
		loop := this
		for {
			if loop == nil {
				break
			}
			if idx == i+1 {
				a = loop
			}
			if idx == total-i {
				b = loop
			}
			loop = loop.Next
			idx++
		}
		tmp := a.Val
		a.Val = b.Val
		b.Val = tmp
	}
}

/*
效果同reverseList2
减少循环，但是耗内存
*/
func (this *ListNode) reverseList3() {
	if this == nil || this.Next == nil {
		return
	}
	total := 1
	loop := this
	list := make([]*ListNode, 0)
	for {
		list = append(list, loop)
		loop = loop.Next
		if loop.Next == nil {
			total++
			list = append(list, loop)
			break
		}
		total++
	}
	loopNum := total / 2
	for i := 0; i < loopNum; i++ {
		idx1 := i
		idx2 := total - i - 1
		tmp := list[idx1].Val
		list[idx1].Val = list[idx2].Val
		list[idx2].Val = tmp
	}
}
