package standard

import (
	"reflect"

	"github.com/EmilyBjartskular/boggle/gamemode/base"
)

const (
	// Name is the name of the gamemode
	Name = "Standard Boggle"
)

var (
	// Mode is a struct representing this gamemode
	Mode = &base.GameMode{
		Name:        Name,
		DieType:     reflect.String,
		CreateBoard: NewBoard,
	}
)

func validMatch(match base.Match) bool {

	return true
}

// NewBoard returns a gameboard
func NewBoard() *base.Board {

	result := &base.Board{
		IsValidMatch: validMatch,
	}
	return result
}
