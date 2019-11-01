package list

/*
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。
如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：
get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。


示例：
MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
linkedList.get(1);            //返回2
linkedList.deleteAtIndex(1);  //现在链表是1-> 3
linkedList.get(1);            //返回3


提示：
所有val值都在 [1, 1000] 之内。
操作次数将在  [1, 1000] 之内。
请不要使用内置的 LinkedList 库。
*/
type MyLinkedList struct {
	num  int
	Val  int
	Next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	cur := this
	for {
		if index == 0 {
			if cur == nil {
				return -1
			}
			return cur.Val
		}
		if cur == nil {
			break
		}
		cur = cur.Next
		index--
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.num == 0 {
		this.Val = val
	} else {
		node := &MyLinkedList{Val: this.Val, Next: this.Next}
		this.Val = val
		this.Next = node
	}
	this.num++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.num == 0 {
		this.Val = val
	} else {
		node := &MyLinkedList{Val: val}
		cur := this
		for {
			if cur.Next == nil {
				cur.Next = node
				break
			}
			cur = cur.Next
		}
	}
	this.num++
	return
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this.num == 0 {
		if index == 0 {
			this.Val = val
			this.num++
		}
	} else {
		node := &MyLinkedList{Val: val}
		if index == 0 {
			node.Next = this.Next
			node.Val = this.Val
			this.Val = val
		} else {
			pre := &MyLinkedList{Next: this}
			cur := this
			for {
				if index == 0 {
					pre.Next = node
					node.Next = cur
					this.num++
					break
				}
				if cur == nil {
					if index == 0 {
						pre.Next = node
						this.num++
					}
					break
				}
				pre, cur = cur, cur.Next
				index--
			}
		}
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this
	pre := &MyLinkedList{Next: this}
	for {
		if cur == nil {
			break
		}
		/*错误示例
		if index == 0 { // 当删除第0个元素时，不会有效果
			pre.Next = cur.Next
			this.num--
			break
		}
		*/
		if index == 0 { // 注意删除方式!
			if cur.Next == nil {
				pre.Next = nil
			} else {
				cur.Val = cur.Next.Val
				cur.Next = cur.Next.Next
			}
			this.num--
			break
		}
		index--
		pre, cur = cur, cur.Next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
/*
 ["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
[[],[1],[3],[1,2],[1],[1],[1]]
["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
[[],[1],[3],[1,2],[1],[0],[0]]
["MyLinkedList","addAtIndex","addAtIndex","addAtIndex","get"]
[[],[0,10],[0,20],[1,30],[0]]
*/
