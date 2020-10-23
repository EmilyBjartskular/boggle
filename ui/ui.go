package ui

import (
	"log"

	"github.com/EmilyBjartskular/boggle/mode"
	"github.com/common-nighthawk/go-figure"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var (
	// App is the main application of boggle
	App *tview.Application
	// MainFlex is the main flexbox
	MainFlex *tview.Flex
	// MainFrame is the top-level menu's frame
	MainFrame *tview.Frame
	// MainMenu is the top-level menu
	MainMenu *tview.List
	// ModeFrame is the mode selection menu frame
	ModeFrame *tview.Frame
	// ModeMenu is the mode selection menu
	ModeMenu *tview.List

	boggleHeader string
)

func init() {
	App = tview.NewApplication()
	MainFlex = tview.NewFlex()
	boggleHeader = figure.NewFigure("Boggle", "doom", true).String()
}

// Init initializes the display
func Init() {

	// Create main menu
	MainMenu = tview.NewList().
		AddItem("New game", "", '1', displayNewGame).
		AddItem("Join game", "", '2', displayJoinGame).
		AddItem("Quit", "", 'q', func() {
			App.Stop()
		}).
		ShowSecondaryText(false)

	MainFrame = tview.NewFrame(MainMenu)
	MainFrame.SetBorder(true)
	MainFrame.AddText("test\ntest", true, tview.AlignCenter, tcell.ColorGreen)
	MainFrame.AddText("test\ntest", true, tview.AlignCenter, tcell.ColorGreen)
	log.Println(boggleHeader)

	// Create new game menu
	ModeMenu = tview.NewList().ShowSecondaryText(false)
	for _, v := range mode.Modes {
		ModeMenu.AddItem(v.Name, "", 0, displayGameMode)
	}

	MainFlex.AddItem(MainFrame, 0, 1, true).
		AddItem(ModeMenu, 0, 1, false)

	if err := App.SetRoot(MainFlex, true).SetFocus(MainFlex).Run(); err != nil {
		panic(err)
	}
}

func displayNewGame() {

}

func displayJoinGame() {

}

func displayGameMode() {

}

func displayGame() {

}
