package stack

/*
使用栈实现队列的下列操作：
push(x) -- 将一个元素放入队列的尾部。
pop() -- 从队列首部移除元素。
peek() -- 返回队列首部的元素。
empty() -- 返回队列是否为空。

示例:
MyQueue queue = new MyQueue();
queue.push(1);
queue.push(2);
queue.peek();  // 返回 1
queue.pop();   // 返回 1
queue.empty(); // 返回 false

说明:
你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。
*/
type MyQueue struct {
	node *stackNode
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	node := stackNode{val: x}
	node.next = this.node
	this.node = &node
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	cur := this.node
	var pre *stackNode
	for cur != nil {
		if cur.next == nil {
			if pre == nil {
				this.node = nil
			} else {
				pre.next = nil
			}
		}
		pre, cur = cur, cur.next
	}
	return pre.Int()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	cur := this.node
	var pre *stackNode
	for cur != nil {
		pre, cur = cur, cur.next
	}
	return pre.Int()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if this.node == nil {
		return true
	}
	return false
}
