package util

import (
	"time"

	"github.com/emirpasic/gods/trees/btree"
)

// TreeSearchTime Returns total time taken to find all strings of list in tree
func TreeSearchTime(list []string, tree *btree.Tree) (tot int64) {
	tot = 0
	for _, v := range list {
		start := time.Now()
		_, ok := tree.Get(v)
		elapsed := time.Since(start)
		if !ok {
			panic("Did not find [" + v + "]")
		}
		tot += elapsed.Nanoseconds()
	}
	return
}

// ListSearchTime Return total time taken find all strings from list in itself
func ListSearchTime(list []string) (tot int64) {
	tot = 0
	for _, v := range list {
		start := time.Now()
		ok := false
		for i := range list {
			if list[i] == v {
				ok = true
				break
			}
		}
		elapsed := time.Since(start)
		if !ok {
			panic("Did not find [" + v + "]")
		}
		tot += elapsed.Nanoseconds()
	}
	return
}
