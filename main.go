package main

import (
	"github.com/EmilyBjartskular/boggle/sdk"
	"github.com/emirpasic/gods/trees/btree"
)

//go:generate go run scripts/includetxt.go

func main() {
	lst := make(chan *btree.Tree)
	go sdk.CreateTree(CollinsScrabbleWords2019, lst)

	tree := <-lst
}
