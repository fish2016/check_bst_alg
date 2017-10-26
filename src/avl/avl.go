package avl

import (
	"bst"
)

func MAX(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/*
type AVLNode struct {
	key       int
	leftNode  *AVLNode
	rightNode *AVLNode
	height    int
}

type AVLTree struct {
	root *AVLNode
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
func LL_Rotation(root *AVLNode) *AVLNode {
	left := root.leftNode

	root.leftNode = left.rightNode

	left.rightNode = root

	//更新height
	root.height = MAX(root.leftNode.height, root.rightNode.height) + 1
	left.height = MAX(left.leftNode.height, left.rightNode.height) + 1

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
func RR_Rotation(root *AVLNode) *AVLNode {
	right := root.rightNode

	root.rightNode = right.leftNode

	right.leftNode = root

	root.height = MAX(root.leftNode.height, root.rightNode.height) + 1
	right.height = MAX(right.leftNode.height, right.rightNode.height) + 1

	return right

}

/*
左子树不平衡，左子树的右子树更深

					a
				b		c
			d	  e
			        f


*/

func LR_Rotation(root *AVLNode) *AVLNode {
	//先将左子树施以RR_Rotation
	root.leftNode = RR_Rotation(root.leftNode)
	return LL_Rotation(root)

}

func RL_Rotation(root *AVLNode) *AVLNode {
	//先将右子树施以LL_Rotation
	root.rightNode = RR_Rotation(root.rightNode)
	return RR_Rotation(root)

}

func (this *AVLTree) insert(key int) {

	//按照bst规则插入tree
	/*
		var current *AVLNode = nil
		var parent *AVLNode = nil
		height := 0
		node := &AVLNode{key, nil, nil, 0}
		if this.root == nil {
			node.height = 1
			this.root = node

			return
		} else if key < node.key {

		}

			for current != nil {
				height++
				if current.key > node.key {
					parent = current
					current = current.leftNode
					if current == nil {
						node.height = height
						parent.leftNode = node
					}
				} else {
					parent = current
					current = current.rightNode
					if current == nil {
						node.height = height
						parent.leftNode = node
					}
				}

			}
	*/
	//if this.root.leftNode.height-this.root.rightNode.height == 2 {
	//this.root
	//}

}

func action(node *bst.Tree) {

}

func Bst2Avl(tree *bst.Tree) {
	bst.PostOrderTraverse(tree.root, action)

}
