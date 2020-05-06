package list

import (
	"fmt"
	"strings"
)

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func (this *Node) SimpleString() string {
	if this == nil {
		return ""
	}
	a := []string{}
	l := this
	for l != nil {
		a = append(a, fmt.Sprintf("%d", l.Val))
		l = l.Next
	}
	return strings.Join(a, "->")
}

func createNode1(in []interface{}) *Node {
	var head, ret, l *Node
	for _, val := range in {
		if vInt, ok := val.(int); ok {
			if ret == nil {
				ret = &Node{Val: vInt}
				head = ret
				l = ret
				continue
			}
			if head == nil {
				head = &Node{Val: vInt}
				l.Child = head
				l = head
			} else {
				l.Next = &Node{Val: vInt}
				l = l.Next
			}
		} else {
			if head != nil {
				l = head
				head = nil
			} else {
				l = l.Next
			}
		}
	}
	return ret
}
