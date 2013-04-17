tree
====

This project now requires go 1.1 to build!

A strongly typed binary search tree in Go without casts or reflection

godoc documentation here: http://godoc.org/github.com/natefinch/tree

The basic idea is simply to separate data storage from data organization.  The tree stores the organization, you store the data in a slice or other indexable data structure.

The organization tree uses a compare function similar to that of the standard library's package sort to compare using indexes into your data structure.

See https://github.com/natefinch/treesample for a working sample


License:
--------
MIT License

