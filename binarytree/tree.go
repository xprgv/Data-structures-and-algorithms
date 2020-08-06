package binarytree

import "fmt"

// Tree struct
type Tree struct {
	root *nodeType
}

type nodeType struct {
	value int
	left  *nodeType
	right *nodeType
}

// New - create new tree
func New() Tree {
	return Tree{
		root: nil,
	}
}

// Add new node to tree
func (tree *Tree) Add(value int) {
	if tree.root == nil {
		tree.root = &nodeType{
			value: value,
			left:  nil,
			right: nil,
		}
	} else {
		addNodeRecursively(tree.root, value)
	}
}

func addNodeRecursively(node *nodeType, value int) *nodeType {
	if node == nil {
		return &nodeType{
			value: value,
			left:  nil,
			right: nil,
		}
	} else if value < node.value {
		node.left = addNodeRecursively(node.left, value)
	} else if value >= node.value {
		node.right = addNodeRecursively(node.right, value)
	}
	return node
}

// RemoveItem from tree
func (tree *Tree) RemoveItem(value int) {
	removeItemRecursively(tree.root, value)
}

func removeItemRecursively(node *nodeType, value int) *nodeType {
	return nil
}

// IsEmpty - reture bool value
func (tree *Tree) IsEmpty() bool {
	return tree.root == nil
}

// Clear - delete all items from tree
func (tree *Tree) Clear() {
	// clear(tree.root)

	// GC will delete all unattainable nodes
	tree.root = nil
}

// Print - infix print
func (tree *Tree) Print() {
	infixPrint(tree.root)
}

// PrefixPrint ...
func (tree *Tree) PrefixPrint() {
	prefixPrint(tree.root)
}

func prefixPrint(node *nodeType) {
	if node != nil {
		fmt.Printf("%v\n", node.value)
		prefixPrint(node.left)
		prefixPrint(node.right)
	}
}

// InfixPrint ...
func (tree *Tree) InfixPrint() {
	infixPrint(tree.root)
}

func infixPrint(node *nodeType) {
	if node != nil {
		infixPrint(node.left)
		fmt.Printf("%v\n", node.value)
		infixPrint(node.right)
	}
}

// PostfixPrint ...
func (tree *Tree) PostfixPrint() {
	postfixPrint(tree.root)
}

func postfixPrint(node *nodeType) {
	if node != nil {
		postfixPrint(node.left)
		postfixPrint(node.right)
		fmt.Printf("%v\n", node.value)
	}
}
