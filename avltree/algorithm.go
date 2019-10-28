package avltree

import (
	"errors"

	"github.com/andremissaglia/algorithms_study/internal/utils"
)

// ErrNotFound means the desired @key is not in the @items
var ErrNotFound error = errors.New("Not Found")

// AvlTree is a data structure that inserts/retrieves data with O(log n) complexity
type AvlTree interface {
	Insert(elem utils.Element)
	Remove(key string)
	Get(key string) (utils.Element, error)
	PreOrderTraversal() []utils.Element
}

type implAvl struct {
	root *node
}

type node struct {
	elem   utils.Element
	left   *node
	right  *node
	height int
}

// New returns an empty AVL Tree
func New() AvlTree {
	return &implAvl{}
}

func (tree *implAvl) Insert(elem utils.Element) {
	tree.root = tree.root.insert(&elem)

}
func (tree *implAvl) Remove(key string) {
	tree.root = tree.root.deleteNode(key)
}
func (tree *implAvl) Get(key string) (utils.Element, error) {
	node := tree.root
	for node != nil {
		if key == node.elem.Key {
			return node.elem, nil
		} else if key > node.elem.Key {
			node = node.right
		} else {
			node = node.left
		}
	}
	return utils.Element{}, ErrNotFound
}

func (tree *implAvl) PreOrderTraversal() []utils.Element {
	elements := []utils.Element{}
	if tree.root != nil {
		elements = tree.root.recursiveTraversal(elements)
	}
	return elements
}

func (n *node) insert(elem *utils.Element) *node {
	if n == nil {
		return &node{
			elem:   *elem,
			height: 1,
		}
	}
	if elem.Key < n.elem.Key {
		n.left = n.left.insert(elem)
	} else {
		n.right = n.right.insert(elem)
	}

	n = n.rebalance()
	n.updateHeight()
	return n
}

func (n *node) deleteNode(key string) *node {
	if n == nil {
		return nil
	}
	if key == n.elem.Key {
		if n.left == nil && n.right == nil {
			// Case 1: leaf
			return nil
		} else if n.left != nil && n.right == nil {
			// Case 2: 1 child (left)
			return n.left
		} else if n.left == nil && n.right != nil {
			// Case 2: 1 child (right)
			return n.right
		} else {
			// Case 3: 2 children
			successor := n.right.minValueNode()
			n.right = n.deleteNode(successor.elem.Key)
			n.elem = successor.elem
		}
	} else if key < n.elem.Key {
		n.left = n.left.deleteNode(key)
	} else {
		n.right = n.right.deleteNode(key)
	}
	n = n.rebalance()
	n.updateHeight()

	return n
}

func (n *node) recursiveTraversal(elements []utils.Element) []utils.Element {
	if n == nil {
		return elements
	}
	elements = append(elements, n.elem)
	elements = n.left.recursiveTraversal(elements)
	elements = n.right.recursiveTraversal(elements)
	return elements
}

func (n *node) leftRotate() *node {
	newRoot := n.right
	middleSubtree := n.right.left
	newRoot.left = n
	n.right = middleSubtree

	n.updateHeight()
	newRoot.updateHeight()
	return newRoot
}

func (n *node) rightRotate() *node {
	newRoot := n.left
	middleSubtree := n.left.right
	newRoot.right = n
	n.left = middleSubtree

	n.updateHeight()
	newRoot.updateHeight()
	return newRoot
}

func (n *node) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}
func (n *node) updateHeight() {
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
}

func (n *node) getBalance() int {
	return n.right.getHeight() - n.left.getHeight()
}

func (n *node) rebalance() *node {
	balance := n.getBalance()
	if balance < -1 {
		if n.left != nil && n.left.getBalance() > 0 {
			// Left-right case
			n.left = n.left.leftRotate()
		}
		n = n.rightRotate()

	} else if balance > 1 {
		if n.right != nil && n.right.getBalance() < 0 {
			// right-left case
			n.right = n.right.rightRotate()
		}
		n = n.leftRotate()

	}
	return n
}

func (n *node) minValueNode() *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
