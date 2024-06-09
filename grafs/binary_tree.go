package main

import (
	"fmt"
)

func main() {
	var err error
	tree := &BinaryTree{}
	tree.Insert(7)
	tree.Insert(3)
	tree.Insert(9)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(8)
	tree.Insert(10)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(6)

	// tree.root.Print("")

	err = tree.Delete(7)
	if err != nil {
		fmt.Println(err, "7")
	}

	tree.root.Print("")
}

type BinaryTree struct {
	root *BinaryTreeNode
}

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func (t *BinaryTree) Insert(value int) error {
	if t.root == nil {
		t.root = &BinaryTreeNode{value, nil, nil}
		return nil
	}

	currentNode := t.root
	parentNode := t.root

	for {
		parentNode = currentNode

		if value < currentNode.value {
			currentNode = currentNode.left
			if currentNode == nil {
				parentNode.left = &BinaryTreeNode{value, nil, nil}
				return nil
			}
		} else {
			currentNode = currentNode.right
			if currentNode == nil {
				parentNode.right = &BinaryTreeNode{value, nil, nil}
				return nil
			}
		}
	}
}

func (n *BinaryTree) Delete(value int) error {
	if n.root == nil {
		return fmt.Errorf("tree is empty")
	}

	currentNode, parentNode, isLeft, err := n.root.Find(value)

	if err != nil {
		return err
	}

	// value is leaf
	if currentNode.left == nil && currentNode.right == nil {
		if n.root == currentNode {
			n.root = nil
		} else if isLeft {
			parentNode.left = nil
		} else {
			parentNode.right = nil
		}
		return nil
	}

	// value has one child
	// if left child is nil
	if currentNode.left == nil {
		if n.root == currentNode {
			n.root = currentNode.right
		} else if isLeft {
			parentNode.left = currentNode.right
		} else {
			parentNode.right = currentNode.right
		}
		return nil
	} else if currentNode.right == nil {
		if n.root == currentNode {
			n.root = currentNode.left
		} else if isLeft {
			parentNode.left = currentNode.left
		} else {
			parentNode.right = currentNode.left
		}
		return nil
	} else {
		// value has two children
		// find successor
		successor := currentNode.right
		successorParent := currentNode
		for {
			if successor.left == nil {
				break
			}
			successorParent = successor
			successor = successor.left
		}
		// replace currentNode value with successor value
		currentNode.value = successor.value
		// delete successor
		if successor == currentNode.right {
			// successor is right child of currentNode
			successorParent.right = nil
		} else {
			// successor is left child of currentNode
			successorParent.left = successor.right
		}
	}

	return nil
}

func (n *BinaryTreeNode) Find(value int) (current *BinaryTreeNode, parrent *BinaryTreeNode, isLeft bool, err error) {
	if n == nil {
		err = fmt.Errorf("value not found")
		return
	}

	current = n
	parrent = n
	isLeft = true

	for {
		if current.value == value {
			return
		} else if value < current.value {
			isLeft = true
			parrent = current
			current = current.left
		} else {
			isLeft = false
			parrent = current
			current = current.right
		}

		if current == nil {
			err = fmt.Errorf("value not found")
			return
		}
	}
}

func (root *BinaryTreeNode) Print(level string) {
	if root != nil {
		// print root
		fmt.Println(level, root.value)
		// print left tree
		root.left.Print(level + " ")
		// print right tree
		root.right.Print(level + " ")
	}
}
