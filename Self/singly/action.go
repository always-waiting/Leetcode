package singly

import (
	"fmt"
	//"unsafe"
)

func (this *SinglyChain) Search(name string) (node *SinglyChain, pos int) {
	if this.Name == name {
		return this, 0
	}
	for {
		if this.Next == nil {
			break
		}
		pos = pos + 1
		if this.Next.Name == name {
			node = this.Next
			break
		}
		this = this.Next
	}
	if node == nil {
		pos = -1
	}
	return
}

func SearchExample1() {
	line := TailCreate([]string{"a", "b", "c", "d"})
	a, pos := line.Search("a")
	fmt.Printf("查到到a在: %d\n", pos)
	a.Print()
	c, pos := line.Search("c")
	fmt.Printf("查到到c在: %d\n", pos)
	c.Print()
	none, pos := line.Search("m")
	fmt.Printf("查到到m在: %d\n", pos)
	fmt.Println(none)
}

func (this *SinglyChain) Replace(id int, name string) {
	if id == 0 {
		this.Name = name
	}
	for {
		id = id - 1
		if this.Next == nil {
			break
		}
		this = this.Next
		if id == 0 {
			this.Name = name
		}
	}
}

func ReplaceExample1() {
	line := TailCreate([]string{"a", "b", "c", "d"})
	line.Replace(2, "f")
	line.Print()
	line.Replace(10, "f")
	line.Print()
}

func (this *SinglyChain) Insert(id int, name string) {
	if id == 0 { //放到最前端
		tmp := *this
		this.Name = name
		this.Next = &tmp
		return
	}

	for {
		id = id - 1
		if id > 0 && this.Next == nil {
			add := SinglyChain{}
			this.Next = &add
		}
		if id == 0 {
			add := SinglyChain{Name: name}
			if this.Next == nil {
				this.Next = &add
			} else {
				tmp := this.Next
				this.Next = &add
				add.Next = tmp
			}
			break
		}
		this = this.Next
	}
}

func InsertExample1() {
	line := TailCreate([]string{"a", "b", "c", "d"})
	line.Insert(0, "f")
	line.Print()
	line.Insert(2, "f")
	line.Print()
	line.Insert(5, "f")
	line.Print()
	line.Insert(8, "f")
	line.Print()
}

func (this *SinglyChain) Delete(id int) {
	if id == 0 {
		if this.Next != nil {
			this.Name = this.Next.Name
			this.Next = this.Next.Next
		} else {
			fmt.Println("这里有问题，无法删除长度为1的链表,可以用一个空元素作为头，使得长度从1开始")
			//a := (*SinglyChain)(unsafe.Pointer(uintptr(unsafe.Pointer(this))))
			fmt.Printf("%p\n", this)
			fmt.Printf("%p\n", &this)
			//fmt.Printf("%s\n", a.Name)
			//a = nil
		}
		return
	}
	for {
		if this.Next == nil {
			break
		}
		if id == 0 {
			this.Name = this.Next.Name
			this.Next = this.Next.Next
			break
		}
		id = id - 1
		this = this.Next
	}
}

func DeleteExample1() {
	line := TailCreate([]string{"a", "b", "c", "d"})
	line.Delete(0)
	line.Print()
	line.Insert(0, "a")
	line.Print()
	line.Delete(2)
	line.Print()
	line.Insert(2, "c")
	line.Print()
	line.Delete(10)
	line.Print()
}

func DeleteExample2() {
	line := TailCreate([]string{"a"})
	line.Delete(0)
	line.Print()
	fmt.Printf("%p\n", line)
	fmt.Printf("%p\n", &line)
	line = nil
	fmt.Printf("%v", line)

}
