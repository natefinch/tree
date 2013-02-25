package tree

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

// Returns
func New() *Tree {
	return &Tree{Val: -1}
}

// Insert adds the value to the tree
// i is the key of the item in your datastructure
// cmp is a comparison function that should return
// -1 if <val> < datastructure[i]
//  0 if <val> == datastructure[i]
//  1 if <val> > datastructure[i]
func (t *Tree) Insert(i int, cmp func(int) int8) {
	if t.Val == -1 {
		t.Val = i
		return
	}
	for {
		switch cmp(t.Val) {
		case -1:
			if t.Left == nil {
				t.Left = &Tree{Val: i}
				return
			}
			t = t.Left
		case 0, 1:
			if t.Right == nil {
				t.Right = &Tree{Val: i}
				return
			}
			t = t.Right
		default:
			panic("Comparison function should only return 0, 1, or -1")
		}
	}
	panic("impossible")
}

// Search returns the index of the item if it is in the tree or -1 if it is not
// cmp is a comparison function that should return
// -1 if <val> < datastructure[i]
//  0 if <val> == datastructure[i]
//  1 if <val> > datastructure[i]
func (t *Tree) Search(cmp func(int) int8) int {
	t = t.search(cmp)
	if t == nil {
		return -1
	}
	return t.Val
}

// Find returns the index of the item if it is in the tree or -1 if it is not
// cmp is a comparison function that should return
// -1 if <val> < datastructure[i]
//  0 if <val> == datastructure[i]
//  1 if <val> > datastructure[i]
func (t *Tree) search(cmp func(int) int8) *Tree {
	if cmp == nil {
		panic("Nil comparison function!")
	}

	for t != nil {
		switch cmp(t.Val) {
		case -1:
			t = t.Left
		case 0:
			return t
		case 1:
			t = t.Right
		default:
			panic("Comparison function should only return 0, 1, or -1")
		}
	}
	return nil
}

// in-order walk of a tree using recursion
func Walk(t *Tree, f func(*Tree)) {
	if t.Left != nil {
		Walk(t.Left, f)
	}
	f(t)
	if t.Right != nil {
		Walk(t.Right, f)
	}
}
