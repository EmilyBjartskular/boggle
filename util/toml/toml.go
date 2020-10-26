package toml

import (
	"fmt"
	"log"

	"github.com/pelletier/go-toml"
)

// GetProperty returns a property from a tomlfile
func GetProperty(prop string, filepath string) interface{} {
	file, err := toml.LoadFile(filepath)
	if err != nil {
		panic(err)
	}
	if !file.Has(prop) {
		log.Fatalln(fmt.Sprintf("Missing `%s` in `%s`", prop, filepath))
	}
	return file.Get(prop)
}

// GetTree returns a toml tree from a tomlfile
func GetTree(filepath string) *toml.Tree {
	tree, err := toml.LoadFile(filepath)
	if err != nil {
		panic(err)
	}
	return tree
}
