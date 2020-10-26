package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var (
	// App is the main application of boggle
	App *tview.Application

	// Pages contains all display states of App
	// observe that display states are not the
	// same as game states
	Pages *tview.Pages

	// Main menu variables
	mainFlex   *tview.Flex
	mainWidth  int
	mainHeight int

	// Mode selection menu variables
	modeFlex   *tview.Flex
	modeWidth  int
	modeHeight int

	// In Game ui variables
	gameGrid *tview.Grid

	playerNameMaxLength int = 15
)

// Display States
var (
	mainMenuPage      *tview.Flex
	modeMenuPage      *tview.Flex
	settingsMenuPage  *tview.Flex
	preGameLobbyPage  *tview.Flex
	inGamePage        *tview.Flex
	postGameLobbyPage *tview.Flex
)

func listVimNavigation(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'h':
		return tcell.NewEventKey(tcell.KeyBackspace, 0, tcell.ModNone)
	case 'j':
		return tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)
	case 'k':
		return tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)
	case 'l':
		return tcell.NewEventKey(tcell.KeyEnter, 0, tcell.ModNone)
	}
	return event
}
