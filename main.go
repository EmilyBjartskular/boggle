package main

import (
	"github.com/EmilyBjartskular/boggle/ui"
)

func main() {
	//lst := make(chan *btree.Tree)
	//go sdk.CreateTree(CollinsScrabbleWords2019, lst)

	//tree := <-lst
	//log.Println(tree.Get("AAA"))

	if err := ui.App.SetRoot(ui.Pages, true).Run(); err != nil {
		panic(err)
	}
}
