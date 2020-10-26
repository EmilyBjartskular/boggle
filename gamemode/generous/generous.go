package generous

import (
	"reflect"

	"github.com/EmilyBjartskular/boggle/gamemode/base"
)

const (
	// Name is the name of the gamemode
	Name = "Generous Boggle"
)

var (
	// Mode is a struct representing this gamemode
	Mode = &base.GameMode{
		Name:    Name,
		DieType: reflect.String,
	}
)
