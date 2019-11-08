package stack

/*
设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。

示例:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
*/

type MinStack struct {
	min  *int
	node *stackNode
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	node := stackNode{val: x}
	node.next = this.node
	this.node = &node
	if this.min == nil {
		this.min = &x
	} else {
		if *(this.min) > x {
			this.min = &x
		}
	}
}

func (this *MinStack) Pop() {
	if this.node == nil {
		panic("minstack is empty")
	}
	node := this.node
	this.node = node.next
	if node.Int() == *(this.min) {
		var val *int
		cur := this.node
		for cur != nil {
			num := cur.Int()
			if val == nil || *val > num {
				val = &num
			}
			cur = cur.next
		}
		this.min = val
	}
}

func (this *MinStack) Top() int {
	if this.node == nil {
		panic("minstack is empty")
	}
	return this.node.Int()
}

func (this *MinStack) GetMin() int {
	return *this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
