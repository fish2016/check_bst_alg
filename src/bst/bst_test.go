package bst

import (
	"testing"
)

func Test_CheckBst(t *testing.T) {
	arr := []int{0, 1, 2}
	tree := CreateBinTree(arr)
	result := tree.CheckBst()

	if result != false {
		t.Error("fail: case 1")
	}

	arr2 := []int{1, 0, 2}
	tree = CreateBinTree(arr2)
	result = tree.CheckBst()

	if result != true {
		t.Error("fail: case 2")
	}

	arr3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	tree = CreateBinTree(arr3)
	result = tree.CheckBst()

	if result != false {
		t.Error("fail: case 3")
	}

	arr4 := []int{100, 50, 150, 40, 60, 110, 160, 30, 50, 35, 70}
	tree = CreateBinTree(arr4)
	result = tree.CheckBst()

	if result != true {
		t.Error("fail: case 4")
	}
}
