package battle

import (
	"reflect"

	"github.com/EmilyBjartskular/boggle/gamemode/base"
)

const (
	// Name is the name of the gamemode
	Name = "Battle Boggle"
)

var (
	// Mode is a struct representing this gamemode
	Mode = &base.GameMode{
		Name:    Name,
		DieType: reflect.String,
	}
)
