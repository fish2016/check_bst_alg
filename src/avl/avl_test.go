package avl

import (
	"bst"
	"fmt"
	"testing"
)

func display(node *bst.Node) {
	fmt.Printf("%d, ", node.Key)
}

func Test_CheckBst(t *testing.T) {
	//rightRotation(nil)
	var root *bst.Node
	arr := []int{1, 2, 3, 4, 5, 6, 110, 111, 201, 202, 203, 205}
	for _, v := range arr {
		//fmt.Println("----loop---", i)
		root = InsertAVL(root, v)
	}
	bst.PostOrderTraverse(root, display)
	fmt.Println("\n========")
}
