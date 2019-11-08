package stack

type StrArrStack struct {
	items []string
}

// 向栈中添加元素
func (this *StrArrStack) Push(item string) {
	if this.items == nil {
		this.items = make([]string, 0)
	}
	this.items = append(this.items, item)
}

// 从栈中获取元素
func (this *StrArrStack) Pop() string {
	if this.items == nil || len(this.items) == 0 {
		panic("stack is empty")
	}
	len := len(this.items)
	ret := this.items[len-1]
	this.items = this.items[:len-1]
	return ret
}

// 返回栈顶元素值
func (this *StrArrStack) Top() string {
	if this.items == nil || len(this.items) == 0 {
		panic("stack is empty")
	}
	return this.items[len(this.items)-1]
}

func (this StrArrStack) IsEmpty() bool {
	if this.items == nil || len(this.items) == 0 {
		return true
	}
	return false
}

// 用链表实现栈
type stackNode struct {
	val  string
	next *stackNode
}

type StrListStack struct {
	node *stackNode
}

func (this *StrListStack) Push(item string) {
	node := stackNode{val: item}
	node.next = this.node
	this.node = &node
}

func (this *StrListStack) Pop() string {
	if this.node == nil {
		panic("stack is empty")
	}
	node := this.node
	this.node = node.next
	return node.val
}

func (this *StrListStack) Top() string {
	if this.node == nil {
		panic("stack is empty")
	}
	return this.node.val
}

func (this StrListStack) IsEmpty() bool {
	if this.node == nil {
		return true
	}
	return false
}
