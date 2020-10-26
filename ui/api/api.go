package api

import (
	"fmt"
	"strings"

	"github.com/EmilyBjartskular/boggle/state"
	"github.com/rivo/tview"
)

var (
	// ConnectedTextView Shows connected players
	ConnectedTextView *tview.TextView
	// MaxNameLength max player name length
	MaxNameLength int
)

// UpdatePlayers updates Connected Players ui in PreGameLobby
func UpdatePlayers() {
	text := ""
	fmtStr := fmt.Sprintf("[yellow]%%%ds   %%%ds[white]\n", MaxNameLength, MaxNameLength)
	text += fmt.Sprintf(strings.ReplaceAll(fmtStr, "yellow", "white"), "Player Name", "Address")
	for _, p := range state.PlayerLobby.Players {
		text += fmt.Sprintf(fmtStr, p.PlayerName, p.IPAddress)
	}
	fmt.Fprint(ConnectedTextView, text)
}
