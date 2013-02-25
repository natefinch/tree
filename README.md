tree
====

A statically typed binary search tree in Go without casts or reflection

godoc documentation here: http://godoc.org/github.com/natefinch/tree

The basic idea is simply to separate data storage from data organization.  The tree stores the organization, you store the data in a slice or other indexable data structure.

See https://github.com/natefinch/treesample for a working sample

Example Code:
--------


	package main

	import (
		"fmt"
		"github.com/natefinch/tree"
	)

	// StringData is just a slice that implements a comparison function
	type StringData []string

	// Cmp returns a closure that will compare the given string with
	// the string located at an index of the underlying slice
	func (s StringData) Cmp(val string) func(int) int8 {

		// this is just a standard string compare based on the runes
		return func(idx int) int8 {
			other := []rune(s[idx])
			for i, r := range val {
				if i > len(other) {
					return 1
				}
				c := r - other[i]
				if c < 0 {
					return -1
				}
				if c > 0 {
					return 1
				}
			}
			return 0
		}
	}


	func main() {

		// Create a StringTree from the given slice data
		s := StringData([]string{"banana", "apple", "domino", "cat", "zebra", "monkey", "hippo"})
		t := tree.New()

		// insert all of the data in s into the tree
		// note, this doesn't copy any of the data
		for i, v := range s {
			t.Insert(i, s.Cmp(v))
		}

		// add new a new item
		pine := "pineapple"
		
		// add to backing data
		s = append(s, pine)
	
		// insert into tree
		t.Insert(len(s)-1, s.Cmp(pine))

		// let's do a binary search for the item we just inserted
		t.Search(s.Cmp(pine))  // returns 7
		
		// search for something that's not in there:
		t.Search(s.Cmp("nope")) // returns -1
	}

License:
--------
MIT License
