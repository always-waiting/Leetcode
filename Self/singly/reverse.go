package singly

// 链表反转
func (this *SinglyChain) Reverse() (new SinglyChain) {
	new = SinglyChain{Name: this.Name, Next: nil}
	for {
		if this.Next == nil {
			break
		}
		tmp := new
		new = SinglyChain{Name: this.Next.Name, Next: &tmp}
		// new = Chain{Name: this.Next.Name, Next: &new} 错误！
		this = this.Next
	}
	return

}

func SinglyChainExample1() {
	a := SinglyChain{Name: "a"}
	b := SinglyChain{Name: "b"}
	c := SinglyChain{Name: "c"}
	d := SinglyChain{Name: "d"}
	a.Next = &b
	b.Next = &c
	c.Next = &d
	a.Print()
	a.Reverse().Print()
	d.Print()
	d.Reverse().Print()
	c.Print()
	c.Reverse().Print()
}
