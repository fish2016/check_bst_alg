package bst

import (
	"container/list"
	_ "fmt"
)

type Node struct {
	Key       int
	LeftNode  *Node
	RightNode *Node
	Height    int
}

type Tree struct {
	Root *Node
}

func compair(v int, v2 int) int {
	return v - v2 // v < v2时，返回负数； v > v2时， 返回正数
}

func (this *Tree) TraverseBst(compair func(int, int) int) (result bool) {

	var stack list.List
	var current *Node
	if this.Root == nil {
		return true
	}

	stack.PushBack(this.Root)
	for stack.Len() != 0 {
		elm := stack.Back()
		if elm != nil {
			current = elm.Value.(*Node)

			//fmt.Printf("current: %d\n", (current).Key)
			//fmt.Printf("current l: %d\n", (current).LeftNode.Key)
			//fmt.Printf("current r: %d\n", (current).RightNode.Key)
			stack.Remove(elm)
		}

		if current.LeftNode != nil {
			if current.Key > current.LeftNode.Key {
				stack.PushBack(current.LeftNode)
			} else {
				return false
			}

		}

		if current.RightNode != nil {
			if current.Key <= current.RightNode.Key {
				stack.PushBack(current.RightNode)
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
	if this.Root == nil {
		return true
	}

	stack.PushBack(this.Root)
	for stack.Len() != 0 {
		elm := stack.Back()
		if elm != nil {
			current = elm.Value.(*Node)
			//fmt.Printf("current: %d\n", (*current).Value)
			stack.Remove(elm)
		}

		if current.LeftNode != nil && current.Key > current.LeftNode.Key {
			stack.PushBack(current.LeftNode)
		} else {
			//return false //不是bst
		}

		if current.RightNode != nil && current.Key <= current.RightNode.Key {
			stack.PushBack(current.RightNode)
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
		current.Key = arr[i]
		//fmt.Println("create current:", current.Key)
		if (2*i + 1) < len(treeArr) {
			current.LeftNode = &treeArr[2*i+1]
			//fmt.Println("create left:", 2*i+1)
		}

		if (2*i + 2) < len(treeArr) {
			current.RightNode = &treeArr[2*i+2]
			//fmt.Println("create right:", 2*i+2)
		}

	}

	var tree = Tree{Root: &treeArr[0]}

	return &tree
}

func PostOrderTraverse(node *Node, action func(node *Node)) {
	if node == nil {
		return
	}

	//fmt.Println("[Debug]PostOrderTraverse:", node)
	PostOrderTraverse(node.LeftNode, action)
	PostOrderTraverse(node.RightNode, action)
	action(node)

}
