package gamemode

import (
	"github.com/EmilyBjartskular/boggle/gamemode/base"
	"github.com/EmilyBjartskular/boggle/gamemode/battle"
	"github.com/EmilyBjartskular/boggle/gamemode/foggle"
	"github.com/EmilyBjartskular/boggle/gamemode/generous"
	"github.com/EmilyBjartskular/boggle/gamemode/standard"
)

var (
	// Modes refers to implemented gamemodes
	Modes = []*base.GameMode{
		standard.Mode,
		foggle.Mode,
		generous.Mode,
		battle.Mode,
	}
)
