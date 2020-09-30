package Self

/*
尝试构建平衡二叉搜索树
*/

type AVLNode struct {
	Height int
	Parent *AVLNode
	Val    int
	Left   *AVLNode
	Right  *AVLNode
}

func NewAVLNode(val int) *AVLNode {
	return &AVLNode{Val: val, Height: 1}
}

func PreorderTraversal(root *AVLNode) []int {
	if root == nil {
		return nil
	}
	ret := []int{root.Val}
	if root.Left != nil {
		left := PreorderTraversal(root.Left)
		ret = append(ret, left...)
	}
	if root.Right != nil {
		right := PreorderTraversal(root.Right)
		ret = append(ret, right...)
	}
	return ret
}

func NewAVLTree(preorder []int, inorder []int) *AVLNode {
	if len(preorder) != len(inorder) {
		panic("输入数组长度不等")
	}
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &AVLNode{Val: rootVal}
	leftNum := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			break
		}
		leftNum++
	}
	preorderLeft := preorder[1 : leftNum+1]
	preorderRight := preorder[leftNum+1:]
	root.Left = NewAVLTree(preorderLeft, inorder[0:leftNum])
	root.Right = NewAVLTree(preorderRight, inorder[leftNum+1:])
	if root.Left != nil {
		root.Left.Parent = root
		root.Height = root.Left.Height
	}
	if root.Right != nil {
		root.Right.Parent = root
		if root.Height < root.Right.Height {
			root.Height = root.Right.Height
		}
	}
	root.Height++
	return root
}

func BalanceFactor(root *AVLNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	} else if root.Left == nil {
		return root.Right.Height
	} else if root.Right == nil {
		return root.Left.Height
	} else {
		return root.Left.Height - root.Right.Height
	}
}

func IsBalanced(root *AVLNode) bool {
	bf := BalanceFactor(root)
	return bf < 2 && bf > -2
}

func InsertOnly(root *AVLNode, val int) *AVLNode {
	if root.Val < val {
		if root.Right == nil {
			root.Right = &AVLNode{Val: val, Parent: root, Height: 1}
		} else {
			InsertOnly(root.Right, val)
		}
	} else {
		if root.Left == nil {
			root.Left = &AVLNode{Val: val, Parent: root, Height: 1}
		} else {
			InsertOnly(root.Left, val)
		}
	}
	if root.Left != nil {
		root.Height = root.Left.Height
	}
	if root.Right != nil && root.Right.Height > root.Height {
		root.Height = root.Right.Height
	}
	root.Height++
	return root
}

// 认为root已经是最小失衡树了
func LL_Roate(root *AVLNode) *AVLNode {
	if IsBalanced(root) {
		return root
	}
	newRight := root
	newRoot := root.Left
	newRight.Parent = newRoot
	newRight.Left = newRoot.Right
	newRoot.Parent = nil
	newRoot.Right = newRight
	newRight.Height = max(newRight.Left.Height, newRight.Right.Height) + 1
	newRoot.Height = max(newRoot.Right.Height, newRoot.Left.Height) + 1
	return newRoot
}

func RR_Roate(root *AVLNode) *AVLNode {
	if IsBalanced(root) {
		return root
	}
	newLeft := root
	newRoot := root.Right
	newLeft.Parent = newRoot
	newLeft.Right = newRoot.Left
	newRoot.Parent = nil
	newRoot.Left = newLeft
	newLeft.Height = max(newLeft.Left.Height, newLeft.Right.Height) + 1
	newRoot.Height = max(newRoot.Right.Height, newRoot.Left.Height) + 1
	return newRoot
}

func LR_Roate(root *AVLNode) *AVLNode {
	RR_Roate(root.Left)
	return LL_Roate(root)
}

func RL_Roate(root *AVLNode) *AVLNode {
	LL_Roate(root.Right)
	return RR_Roate(root)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func InsertAVLTree(root *AVLNode, val int) *AVLNode {
	if root.Val < val {
		if root.Right == nil {
			root.Right = &AVLNode{Val: val, Parent: root, Height: 1}
		} else {
			InsertAVLTree(root.Right, val)
		}
	} else {
		if root.Left == nil {
			root.Left = &AVLNode{Val: val, Parent: root, Height: 1}
		} else {
			InsertAVLTree(root.Left, val)
		}
	}
	if root.Left != nil {
		root.Height = root.Left.Height
	}
	if root.Right != nil && root.Right.Height > root.Height {
		root.Height = root.Right.Height
	}
	root.Height++
	if !IsBalanced(root) {
		rootBF := BalanceFactor(root)
		leftBF := BalanceFactor(root.Left)
		rightBF := BalanceFactor(root.Right)
		if rootBF > 0 {
			if leftBF > 0 { //LL型
				root = LL_Roate(root)
			} else { // LR型
				root = LR_Roate(root)
			}
		} else {
			if rightBF > 0 { //RL型
				root = RL_Roate(root)
			} else { //RR型
				root = RR_Roate(root)
			}
		}
	}
	return root
}
