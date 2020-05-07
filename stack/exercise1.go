package stack

/*
Contents:
1. 栈的最小值	--	★★
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
/*
type MinStack struct {
	data []int
	min  *int
}

func Constructor() MinStack {
	return MinStack{data: make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if this.min == nil || *this.min > x {
		this.min = &x
	}
}

func (this *MinStack) Pop() {
	i := this.data[len(this.data)-1]
	this.data = this.data[0 : len(this.data)-1]
	if len(this.data) == 0 {
		this.min = nil
		return
	}
	if i == *this.min {
		this.min = &this.data[0]
		for _, val := range this.data {
			if *this.min > val {
				a := val
				this.min = &a
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return *this.min
}
*/
type MinStack struct {
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
