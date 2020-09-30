package Self

import (
	"reflect"
	"testing"
)

func TestCreateAVLTree(t *testing.T) {
	preorder := []int{66, 60, 77, 75, 88}
	inorder := []int{60, 66, 75, 77, 88}
	avlTree := NewAVLTree(preorder, inorder)
	npreorder := PreorderTraversal(avlTree)
	if !reflect.DeepEqual(npreorder, preorder) {
		t.Errorf("前序遍历返回不同: expect(%v), got(%v)", preorder, npreorder)
	}
	if IsBalanced(avlTree) {
		t.Log("创建的是平衡树")
	} else {
		t.Errorf("创建的是非平衡树")
	}
	/*
		InsertOnly(avlTree, 99)
		if IsBalanced(avlTree) {
			t.Errorf("插入后树平衡")
		} else {
			t.Log("插入后树不平衡")
		}
	*/
	avlTree = InsertAVLTree(avlTree, 99)
	if IsBalanced(avlTree) {
		t.Errorf("插入后树平衡")
		t.Log(PreorderTraversal(avlTree))
	} else {
		t.Log("插入后树不平衡")
	}
}
