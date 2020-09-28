package sdk

import (
	// ok
	"github.com/emirpasic/gods/trees/btree"
)

// CreateTree ok
func CreateTree(list []string, c chan *btree.Tree) {
	tree := btree.NewWithStringComparator(3)
	for _, v := range list {
		tree.Put(v, "")
	}
	c <- tree
}
