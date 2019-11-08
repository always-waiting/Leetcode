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
