package foggle

import (
	"reflect"

	"github.com/EmilyBjartskular/boggle/gamemode/base"
)

const (
	// Name is the name of the gamemode
	Name = "Foggle Boggle"
)

var (
	supportedSizes []base.Size = []base.Size{{4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}, {10, 10}}
	// Mode is a struct representing this gamemode
	Mode = &base.GameMode{
		Name:           Name,
		DieType:        reflect.Int,
		SupportedSizes: supportedSizes,
	}
)
