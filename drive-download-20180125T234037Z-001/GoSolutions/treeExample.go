package main

import "fmt"

type node struct {
	id		int64
	left, right	*node
}

func insert(root *node, i int64) *node {
	if root == nil {
			return &node{i, nil, nil}
	}
	if i < root.id {
		root.left = insert(root.left, i)
	} else if i > root.id {
		root.right = insert(root.right, i)
	}
	return root
}

func treePrint(root *node) {
	if root.left != nil {
		treePrint(root.left)
	}
	fmt.Println(root.id)
	if root.right != nil {
		treePrint(root.right)
	}
}

func main() {
	root := node{100,nil,nil}
	insert(&root, 178)
	insert(&root, 5)
	insert(&root, 55)
	insert(&root, 135)
	fmt.Println(root)
	treePrint(&root)
}

