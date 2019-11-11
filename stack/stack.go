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
type stackStrNode struct {
	val  string
	next *stackStrNode
}

type StrListStack struct {
	node *stackStrNode
}

func (this *StrListStack) Push(item string) {
	node := stackStrNode{val: item}
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

// 用interface{}表示值
type stackNode struct {
	val  interface{}
	next *stackNode
}

func (this *stackNode) Int() int {
	val, ok := this.val.(int)
	if !ok {
		panic("不能转化为int")
	}
	return val
}

func (this *stackNode) String() string {
	val, ok := this.val.(string)
	if !ok {
		panic("不能转换为string")
	}
	return val
}

func (this *stackNode) Rune() rune {
	val, ok := this.val.(rune)
	if !ok {
		panic("不能转换为rune")
	}
	return val
}

type Stack struct {
	node *stackNode
}

func (this *Stack) Push(item interface{}) {
	node := stackNode{val: item}
	node.next = this.node
	this.node = &node
}

func (this *Stack) Pop() *stackNode {
	node := this.node
	if this.node != nil {
		this.node = this.node.next
	}
	return node
}

func (this *Stack) Top() *stackNode {
	return this.node
}

func (this *Stack) IsEmpty() bool {
	if this.node == nil {
		return true
	}
	return false
}

func (this *Stack) Len() int {
	cur := this.node
	len := 0
	for cur != nil {
		len++
		cur = cur.next
	}
	return len
}
