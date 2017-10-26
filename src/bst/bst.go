package bst

import (
	"container/list"
	"fmt"
)

type Node struct {
	key    int
	pLeft  *Node
	pRight *Node
}

type Tree struct {
	root *Node
}

func compair(v int, v2 int) int {
	return v - v2 // v < v2时，返回负数； v > v2时， 返回正数
}

func (this *Tree) TraverseBst(compair func(int, int) int) (result bool) {

	var stack list.List
	var current *Node
	if this.root == nil {
		return true
	}

	stack.PushBack(this.root)
	for stack.Len() != 0 {
		elm := stack.Back()
		if elm != nil {
			current = elm.Value.(*Node)

			fmt.Printf("current: %d\n", (current).key)
			//fmt.Printf("current l: %d\n", (current).pLeft.key)
			//fmt.Printf("current r: %d\n", (current).pRight.key)
			stack.Remove(elm)
		}

		if current.pLeft != nil {
			if current.key > current.pLeft.key {
				stack.PushBack(current.pLeft)
			} else {
				return false
			}

		}

		if current.pRight != nil {
			if current.key <= current.pRight.key {
				stack.PushBack(current.pRight)
			} else {
				return false
			}
		}
	}

	return true
}

func (this *Tree) TraverseBstAndAction(action func(int)) (result bool) {

	var stack list.List
	var current *Node
	if this.root == nil {
		return true
	}

	stack.PushBack(this.root)
	for stack.Len() != 0 {
		elm := stack.Back()
		if elm != nil {
			current = elm.Value.(*Node)
			//fmt.Printf("current: %d\n", (*current).Value)
			stack.Remove(elm)
		}

		if current.pLeft != nil && current.key > current.pLeft.key {
			stack.PushBack(current.pLeft)
		} else {
			//return false //不是bst
		}

		if current.pRight != nil && current.key <= current.pRight.key {
			stack.PushBack(current.pRight)
		} else {
			return false //不是bst
		}

	}

	return true
}

func (this *Tree) CheckBst() bool {

	return this.TraverseBst(compair)
}

func CreateBinTree(arr []int) *Tree {

	treeArr := make([]Node, len(arr))

	for i, _ := range treeArr {
		current := &treeArr[i]
		current.key = arr[i]
		fmt.Println("create current:", current.key)
		if (2*i + 1) < len(treeArr) {
			current.pLeft = &treeArr[2*i+1]
			fmt.Println("create left:", 2*i+1)
		}

		if (2*i + 2) < len(treeArr) {
			current.pRight = &treeArr[2*i+2]
			fmt.Println("create right:", 2*i+2)
		}

	}

	var tree = Tree{root: &treeArr[0]}

	return &tree
}

func PostOrderTraverse(node *Node, action func(node *Node)) {
	if node == nil {
		return
	}

	PostOrderTraverse(node.pLeft, action)
	PostOrderTraverse(node.pRight, action)
	action(node)

}
