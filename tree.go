// Package tree implements a very simple binary tree without any balancing.
// This is mainly intended as a proof of concept for a strongly typed tree
// in Go without using reflection or casts.
package tree

// Tree holds a binary tree data organization.
//
// Note that Tree is not intended to hold data itself, it just
// maintains a structure, and data is retrieved using a node's Val as the index
// into another data structure that holds the actual values
type Tree struct {
	Head *Node
	Cmp  Comparer
}

// Node is a node in the tree
// Val is an index into an external data structure.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Comparer implements a comparison function that should return
// -1 if data[i] < data[i], 0 if equal, and 1 if greater than
type Comparer interface {
	Compare(i, j int) int8
}

// New Returns a newly initialized Tree.
func New(cmp Comparer) *Tree {
	return &Tree{Cmp: cmp}
}

// Insert adds the value to the tree.
//
// The value should already exist in the backing datastore at index i
func (t *Tree) Insert(i int) {
	if t.Head == nil {
		t.Head = &Node{Val: i}
		return
	}
	if t.Cmp == nil {
		panic("Nil Comparer during Tree Insert")
	}
	n := t.Head
	for {
		switch t.Cmp.Compare(i, n.Val) {
		case -1:
			if n.Left == nil {
				n.Left = &Node{Val: i}
				return
			}
			n = n.Left
		case 0, 1:
			if n.Right == nil {
				n.Right = &Node{Val: i}
				return
			}
			n = n.Right
		default:
			panic("Comparison function should only return 0, 1, or -1")
		}
	}
	panic("impossible")
}

// Search returns the node of the tree with the value at index i if it exists
func (t *Tree) Search(i int) *Node {
	if t.Cmp == nil {
		panic("Nil Comparer during search!")
	}
	n := t.Head
	for n != nil {
		switch t.Cmp.Compare(i, n.Val) {
		case -1:
			n = n.Left
		case 0:
			return n
		case 1:
			n = n.Right
		default:
			panic("Comparison function should only return 0, 1, or -1")
		}
	}
	return nil
}

// Walk implements an in-order walk of a tree using recursion.
func Walk(n *Node, f func(*Node) bool) bool {
	if n.Left != nil {
		if !Walk(n.Left, f) {
			return false
		}
	}
	if !f(n) {
		return false
	}
	if n.Right != nil {
		return Walk(n.Right, f)
	}
	return true
}
