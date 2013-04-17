// Package tree implements a very simple binary tree without any balancing.
// This is mainly intended as a proof of concept for a strongly typed tree
// in Go without using reflection or casts.
package tree

import "errors"

// Tree holds a binary tree data organization.
//
// Note that Tree is not intended to hold data itself, it just
// maintains a structure, and data is retrieved using a node's Val as the index
// into another data structure that holds the actual values
type Tree struct {
	Head *Node
}

// Node is a node in the tree
// Val is an index into an external data structure.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Comparer is a comparison function that should return
// -1 if data[i] < data[i], 0 if equal, and 1 if greater than
type Compare func(i, j int) int8

// Insert adds the value to the tree.
//
// The value should already exist in the backing datastore at index i
// returns the node created in the tree and nil, or nil and and error if there was a problem
func (t *Tree) Insert(i int, cmp Compare) (*Node, error) {
	if t.Head == nil {
		t.Head = &Node{Val: i}

		return t.Head, nil
	}
	if cmp == nil {
		return nil, errors.New("Nil comparison")
	}
	n := t.Head
	for {
		switch cmp(i, n.Val) {
		case -1:
			if n.Left == nil {
				n.Left = &Node{Val: i}
				return n.Left, nil
			}
			n = n.Left
		case 0, 1:
			if n.Right == nil {
				n.Right = &Node{Val: i}
				return n.Right, nil
			}
			n = n.Right
		default:
			return nil, errors.New("Comparison function should only return 0, 1, or -1")
		}
	}
}

// Search returns the node of the tree with the value at index i if it exists, otherwise nil
func (t *Tree) Search(i int, cmp Compare) (*Node, error) {
	if cmp == nil {
		return nil, errors.New("Nil Comparer!")
	}
	n := t.Head
	for n != nil {
		switch cmp(i, n.Val) {
		case -1:
			n = n.Left
		case 0:
			return n, nil
		case 1:
			n = n.Right
		default:
			return nil, errors.New("Comparison function should only return 0, 1, or -1")
		}
	}
	return nil, nil
}

// Delete removes the node of the tree with the value at index i
//
// Delete returns the index of rhte item if the value was found, otherwise will return an error
func (t *Tree) Delete(i int, cmp Compare) (int, error) {
	if cmp == nil {
		return -1, errors.New("Nil Comparer!")
	}

	// parent holds the parent of the node we're deleting
	var parent *Node

	n := t.Head
	for n != nil {
		switch cmp(i, n.Val) {
		default:
			panic("Comparison function should only return 0, 1, or -1")
		case -1:
			parent = n
			n = n.Left
		case 1:
			parent = n
			n = n.Right
		case 0:
			val := n.Val
			// found the node, let's delete it
			if n.Left != nil {
				right := n.Right
				n.Val = n.Left.Val
				n.Left = n.Left.Left
				n.Right = n.Left.Right

				// reinsert the values of the right subtree
				if right != nil {
					sub := &Tree{Head: n}
					Walk(right, func(n *Node) bool {
						sub.Insert(n.Val, cmp)
						return true
					})
				}
				return val, nil
			}

			if n.Right != nil {
				// Left is nil, so we can just move Right up

				n.Val = n.Right.Val
				n.Left = n.Right.Left
				n.Right = n.Right.Right
				return val, nil
			}

			// deleting a leaf node, set parent pointer to nil

			if parent == nil {
				// deleting head
				t.Head = nil
				return val, nil
			}

			if parent.Left == n {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
			return val, nil
		}
	}
	return -1, errors.New("Item not found")
}

// Walk implements an in-order walk of a tree using recursion.
//
// function f should return false if it wants the walk to stop
// Walk returns false if f ever returns false, otherwise true
func Walk(n *Node, f func(*Node) bool) bool {
	if n == nil {
		return true
	}
	if !Walk(n.Left, f) {
		return false
	}
	if !f(n) {
		return false
	}
	return Walk(n.Right, f)
}
