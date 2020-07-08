package list

import (
	"fmt"
)

func test() {
	fmt.Println("测试")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTreeByLevelLoop(in []interface{}) (root *TreeNode) {
	if len(in) == 0 {
		return root
	}
	rNum, ok := in[0].(int)
	if !ok {
		return
	}
	root = &TreeNode{Val: rNum}
	lList := []interface{}{}
	rList := []interface{}{}
	pos := 1
	Lloop := 0
	Rloop := 0
	for pos <= len(in)-1 {
		if pos+p2(Lloop) > len(in)-1 {
			lList = append(lList, in[pos:len(in)]...)
			break
		} else {
			tmp := in[pos : pos+p2(Lloop)]
			if tmp[0] == nil && Lloop == 0 {
				lList = append(lList, nil)
				pos = pos + 1
				rList = in[pos:]
				break
			}
			pos = pos + p2(Lloop)
			for _, val := range tmp {
				if _, ok := val.(int); ok {
					Lloop++
				} else {
					Lloop--
				}
				lList = append(lList, val)
			}
		}
		if pos+p2(Rloop) > len(in)-1 {
			rList = append(rList, in[pos:len(in)]...)
		} else {
			tmp := in[pos : pos+p2(Rloop)]
			if tmp[0] == nil && Rloop == 0 {
				rList = append(rList, nil)
				pos = pos + 1
				lList = append(lList, in[pos:]...)
				break
			}
			pos = pos + p2(Rloop)
			for _, val := range tmp {
				if _, ok := val.(int); ok {
					Rloop++
				} else {
					Rloop--
				}
				rList = append(rList, val)
			}
		}
	}
	//fmt.Println(lList)
	//fmt.Println(rList)
	root.Left = createTreeByLevelLoop(lList)
	root.Right = createTreeByLevelLoop(rList)
	return root
}

func p2(n int) int {
	if n < 0 {
		return 1
	}
	if n == 0 {
		return 1
	} else {
		return 2 * p2(n-1)
	}
}
