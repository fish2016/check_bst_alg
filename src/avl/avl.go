package avl

import (
	"bst"
)

//type bst.Node bst.Node
//type AVLTree bst.Tree

func MaxHeight(left *bst.Node, right *bst.Node) int {
	leftHeight := 0
	rightHeight := 0

	if left != nil {
		leftHeight = left.Height
	}

	if right != nil {
		rightHeight = right.Height
	}

	return MAX(leftHeight, rightHeight)
}

func MAX(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func GetHeight(node *bst.Node) int {
	if node == nil {
		return 0
	}

	return node.Height
}

/*
type bst.Node struct {
	key       int
	LeftNode  *bst.Node
	RightNode *bst.Node
	Height    int
}

type AVLTree struct {
	root *bst.Node
}
*/
/*
			a
		b		c
	d
e
        ||
		 V

		b
	d	   a
e				c

*/
//向右旋转

func LL_Rotation(root *bst.Node) *bst.Node {
	left := root.LeftNode

	root.LeftNode = left.RightNode

	left.RightNode = root

	//更新height
	root.Height = MAX(GetHeight(root.LeftNode), GetHeight(root.RightNode)) + 1
	left.Height = MAX(GetHeight(left.LeftNode), GetHeight(left.RightNode)) + 1

	return left

}

/*

		a
	b		c
				d
					e

		||
		V
		c
	a		d
b				e
*/
//向左旋转

func RR_Rotation(root *bst.Node) *bst.Node {
	right := root.RightNode

	root.RightNode = right.LeftNode

	right.LeftNode = root

	root.Height = MAX(GetHeight(root.LeftNode), GetHeight(root.RightNode)) + 1
	right.Height = MAX(GetHeight(right.LeftNode), GetHeight(right.RightNode)) + 1

	return right

}

/*
左子树不平衡，左子树的右子树更深

					a
				b		c
			d	  e
			        f


*/

func LR_Rotation(root *bst.Node) *bst.Node {
	//先将左子树施以RR_Rotation
	root.LeftNode = RR_Rotation(root.LeftNode)
	return LL_Rotation(root)

}

func RL_Rotation(root *bst.Node) *bst.Node {
	//先将右子树施以LL_Rotation
	root.RightNode = RR_Rotation(root.RightNode)
	return RR_Rotation(root)

}

func InsertAVL(root *bst.Node, key int) *bst.Node {

	if root == nil {
		root = &bst.Node{key, nil, nil, 0}
	}
	//fmt.Println("insert: %d, %d", root.Key, key)
	if root.Key > key {
		root.LeftNode = InsertAVL(root.LeftNode, key)
		if GetHeight(root.LeftNode)-GetHeight(root.RightNode) == 2 {
			if key < root.LeftNode.Key { //LL
				root = LL_Rotation(root)
			} else { //LR
				root = LR_Rotation(root)
			}
		}

	} else if root.Key < key {
		root.RightNode = InsertAVL(root.RightNode, key)
		if GetHeight(root.LeftNode)-GetHeight(root.RightNode) == -2 {
			if key > root.RightNode.Key { //RR
				root = RR_Rotation(root)
			} else { //RL
				root = RL_Rotation(root)
			}

		}
	}

	root.Height = MAX(GetHeight(root.LeftNode), GetHeight(root.RightNode)) + 1
	//fmt.Println("[Debug2]root.Key=", root)
	return root

}

//}
/*
func action(node *bst.Node) {

	if node.LeftNode-node.RightNode == 2 {
        node.LetNode.
	} else if node.LeftNode-node.RightNode == -2 {

	}

	node.Height = MaxHeight(node.LeftNode, node.RightNode)

}

func Bst2Avl(tree *bst.Tree) {
	bst.PostOrderTraverse(tree.Root, action)

}
*/
