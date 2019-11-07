package singly

import "fmt"

// 单链表
type SinglyChain struct {
	Name string       `json:"名称"`
	Next *SinglyChain `json:"单链表下一个元素"`
}

func (this SinglyChain) Print() {
	cur := &this
	for {
		fmt.Printf("%s->", cur.Name)
		if cur.Next == nil {
			fmt.Println("nil")
			break
		}
		cur = cur.Next
	}
}

func BaseExample1() {
	a := SinglyChain{}
	pa := &a
	b := *pa // 产生一个新的对象,a和b的地址不同
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	b.Name = "b"
	b.Print()
	a.Print()
}
