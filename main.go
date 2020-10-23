package main

import (
	"log"

	"github.com/EmilyBjartskular/boggle/locale/wl/enus"
	"github.com/EmilyBjartskular/boggle/mode"
	"github.com/EmilyBjartskular/boggle/ui"
)

func main() {
	if 1 == 0 {
		log.Println(enus.CollinsScrabbleWords2019)
	}
	//lst := make(chan *btree.Tree)
	//go sdk.CreateTree(CollinsScrabbleWords2019, lst)

	//tree := <-lst
	//log.Println(tree.Get("AAA"))

	mode.Init()
	ui.Init()
}
