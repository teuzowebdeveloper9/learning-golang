package main

type Node struct {
	value int
	right *Node
	left  *Node
}

type BinaryTree struct {
	Root *Node
}

func insertNode(node *Node, value int) *Node {
	if node == nil {
		return &Node{value: value}
	}

	if value < node.value {
		node.left = insertNode(node.left, value)
	} else if value > node.value {
		node.right = insertNode(node.right, value)
	}

	return node
}

func searchNode(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if value == node.value {
		return true
	}
	if value < node.value {
		return searchNode(node.left, value)
	}
	return searchNode(node.right, value)
}

func (t *BinaryTree) Insert(value int) {
	t.Root = insertNode(t.Root, value)
}
