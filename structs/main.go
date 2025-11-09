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

func MaxSumSubarray(nums []int, k int) int {
	if k <= 0 || k > len(nums) {
		return 0
	}

	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}

	maxSum := windowSum

	for i := k; i < len(nums); i++ {
		windowSum += nums[i]
		windowSum -= nums[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}
