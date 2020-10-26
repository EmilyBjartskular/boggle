package ui

import (
	"fmt"
	"reflect"

	"github.com/EmilyBjartskular/boggle/consts"
	"github.com/EmilyBjartskular/boggle/gamemode"
	"github.com/EmilyBjartskular/boggle/state"
	"github.com/EmilyBjartskular/boggle/state/logic"
	"github.com/EmilyBjartskular/boggle/ui/api"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func init() {
	App = tview.NewApplication()
	Pages = tview.NewPages()

	mainMenuPage = tview.NewFlex()
	modeMenuPage = tview.NewFlex()
	settingsMenuPage = tview.NewFlex()
	preGameLobbyPage = tview.NewFlex()
	inGamePage = tview.NewFlex()
	postGameLobbyPage = tview.NewFlex()

	initMainMenu()
	initModeMenu()

	Pages.AddPage("mainMenuPage", mainMenuPage, true, true)
	Pages.AddPage("modeMenuPage", modeMenuPage, true, false)
	Pages.AddPage("modeSettingsPage", settingsMenuPage, true, false)
	Pages.AddPage("preGameLobbyPage", preGameLobbyPage, true, false)

	Pages.SwitchToPage("mainMenuPage")
}

// Initialize main menu page
func initMainMenu() {
	// Get logo
	logoWidth, logoHeight, logoBox := LogoArt(consts.GameName)

	// Menu setup
	mainMenu := tview.NewList().
		AddItem("Create Lobby", "", '1', func() {
			Pages.SwitchToPage("modeMenuPage")
		}).
		AddItem("Join Lobby", "", '2', func() {

		}).
		AddItem("Quit", "", 'q', func() {
			App.Stop()
		}).
		ShowSecondaryText(false)
	mainMenu.SetInputCapture(listVimNavigation)

	mainFrame := tview.NewFrame(mainMenu).
		AddText("Main Menu", true, tview.AlignCenter, tcell.ColorWhite)

	// Flexbox setup
	mainFlex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(logoBox, logoHeight, 1, false).
		AddItem(mainFrame, 0, 1, true)

	// logoHeight + (1 * number of items in mainMenu) + (1 * number of primitives in mainFlex)
	// + (1 + 1 * number of textrows in mainFrame)
	mainHeight = logoHeight + 3 + 2 + 2
	mainWidth = logoWidth

	// Endresult flexbox for the main menu page
	resFlex := Center(mainWidth, mainHeight, mainFlex)

	mainMenuPage.AddItem(resFlex, 0, 1, true)
}

// Initialize mode menu page
func initModeMenu() {
	// Get logo
	logoWidth, logoHeight, logoBox := LogoArt(consts.GameName)

	modeWidth = 0
	// Create Mode selection menu
	modeMenu := tview.NewList().ShowSecondaryText(false)
	modeMenu.SetInputCapture(listVimNavigation)

	modeMenu.SetSelectedFunc(func(index int, _, _ string, _ rune) {
		if index == len(gamemode.Modes) {
			return
		}
		state.SetSelectedMode(gamemode.Modes[index])
		createModeSettingsMenu()
		Pages.SwitchToPage("modeSettingsPage")
	})

	for _, v := range gamemode.Modes {
		modeMenu.AddItem(v.Name, "", 0, nil)
		if len(v.Name) > modeWidth {
			modeWidth = len(v.Name)
		}
	}
	modeMenu.AddItem("Cancel", "", 'c', func() {
		Pages.SwitchToPage("mainMenuPage")
	})
	modeFrame := tview.NewFrame(modeMenu)
	modeFrame.AddText("Select Game Mode", true, tview.AlignCenter, tcell.ColorWhite)

	if logoWidth > modeWidth {
		modeWidth = logoWidth
	}
	// logoHeight + 1 * number of items in mainMenu + 1 * number of primitives in mainFlex
	modeHeight = logoHeight + (len(gamemode.Modes) + 1) + 2 + 2

	modeFlex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(logoBox, logoHeight, 1, false).
		AddItem(modeFrame, 0, 1, true)

	resFlex := Center(modeWidth, modeHeight, modeFlex)

	modeMenuPage.AddItem(resFlex, 0, 1, true)
}

// Create a menu for setting modespecific settings
func createModeSettingsMenu() {
	m := state.GetSelectedMode()
	// Get logo
	logoWidth, logoHeight, logoBox := LogoArt(m.Name)

	settingsForm := tview.NewForm()

	nameField := tview.NewInputField().
		SetFieldWidth(playerNameMaxLength).
		SetAcceptanceFunc(func(textToCheck string, lastChar rune) bool {
			if len(textToCheck) == playerNameMaxLength {
				return false
			}
			return true
		}).
		SetLabel("Player Name").
		SetChangedFunc(func(text string) {
			state.SetPlayerName(text)
		})

	settingsForm.AddFormItem(nameField)

	// Different form items depending on the gamemodes die type
	switch m.DieType {
	case reflect.String:
		// Get list of languages
		initial := 0
		var langs []string
		for i, lang := range state.Languages {
			langs = append(langs, lang.Name)
			// Shhh, don't think too much about it ;)
			if lang.Name == "English" {
				initial = i
			}
		}

		state.SetSelectedLanguage(state.Languages[initial])

		// Get list of sizes
		getSizes := func() []string {
			var sizes []string
			for _, s := range state.GetSelectedLanguage().BoardSizes {
				sizes = append(sizes, s.String())
			}
			return sizes
		}

		langDropDown := tview.NewDropDown()
		sizeDropDown := tview.NewDropDown()

		sizeDropDownFunc := func(option string, optionIndex int) {
			state.SetSelectedSize(state.GetSelectedLanguage().BoardSizes[optionIndex])
		}

		sizeDropDown.SetLabel("Boardsize").
			SetOptions(getSizes(), sizeDropDownFunc).
			SetCurrentOption(0)

		langDropDown.SetLabel("Language").
			SetOptions(langs, func(option string, optionIndex int) {
				state.SetSelectedLanguage(state.Languages[optionIndex])
				sizeDropDown.SetOptions(getSizes(), sizeDropDownFunc).
					SetCurrentOption(0)
			}).
			SetCurrentOption(initial)

		settingsForm.
			AddFormItem(langDropDown).
			AddFormItem(sizeDropDown)

	case reflect.Int:
		// Get list of sizes
		getSizes := func() []string {
			var sizes []string
			for _, s := range m.SupportedSizes {
				sizes = append(sizes, s.String())
			}
			return sizes
		}

		sizeDropDown := tview.NewDropDown()

		sizeDropDownFunc := func(option string, optionIndex int) {
			state.SetSelectedSize(m.SupportedSizes[optionIndex])
		}

		sizeDropDown.SetLabel("Boardsize").
			SetOptions(getSizes(), sizeDropDownFunc).
			SetCurrentOption(0)

		settingsForm.AddFormItem(sizeDropDown)
	}

	settingsForm.
		AddButton("Create Lobby", func() {
			logic.CreatePreLobby()
			CreatePreGameInterface()
			Pages.SwitchToPage("preGameLobbyPage")
		}).
		AddButton("Cancel", func() {
			Pages.SwitchToPage("modeMenuPage")
			settingsMenuPage.Clear()
		})

		// TODO: This snippet cause settingsForm to not show for some reason
		/*
			settingsForm.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				if event.Rune() == 'j' {
					return tcell.NewEventKey(tcell.KeyTab, 0, tcell.ModNone)
				}
				if event.Rune() == 'k' {
					return tcell.NewEventKey(tcell.KeyTab, 0, tcell.ModShift)
				}
				if event.Rune() == 'l' {
					return tcell.NewEventKey(tcell.KeyEnter, 0, tcell.ModNone)
				}
				return event
			})
		*/

	settingsFrame := tview.NewFrame(settingsForm).
		AddText(fmt.Sprintf("%s settings", m.Name), true, tview.AlignCenter, tcell.ColorWhite)

	frameFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(logoBox, logoHeight, 1, false).
		AddItem(settingsFrame, 0, 1, true)

	settingsWidth := logoWidth
	settingsHeight := logoHeight + 1 + 2 + 4 + 2 + 4

	resFlex := Center(settingsWidth, settingsHeight, frameFlex)

	settingsMenuPage.AddItem(resFlex, 0, 1, true)
}

// CreatePreGameInterface creates the ui for the pregame lobby
func CreatePreGameInterface() {
	m := state.GetSelectedMode()
	// Get logo
	logoWidth, logoHeight, logoBox := LogoArt(m.Name)
	logoFlex := Center(logoWidth, logoHeight, logoBox)

	// Show connected players
	connTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() {
			App.Draw()
		})
	connFrame := tview.NewFrame(connTextView).
		AddText("Connected Players", true, tview.AlignCenter, tcell.ColorGreen)

	api.MaxNameLength = playerNameMaxLength
	api.ConnectedTextView = connTextView
	api.UpdatePlayers()

	connFlex := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(connFrame, playerNameMaxLength*2+5, 1, false).
		AddItem(nil, 0, 1, false)

	// Show lobby info
	formatString := "[red]%10v [yellow]%20v[white]\n"
	infoString := ""
	infoString += fmt.Sprintf(formatString, "Game mode", state.GetSelectedMode().Name)

	if m.DieType == reflect.String {
		infoString += fmt.Sprintf(formatString, "Language", state.GetSelectedLanguage().Name)
	}

	infoString += fmt.Sprintf(formatString, "Boardsize", state.GetSelectedSize().String())

	infoText := tview.NewTextView().
		SetDynamicColors(true)
	fmt.Fprint(infoText, infoString)
	infoFrame := tview.NewFrame(infoText).
		AddText("Lobby Information", true, tview.AlignLeft, tcell.ColorGreen)

	// Show IP for manual connection
	ipString := "IP Address for manual connection\n"
	ipString += fmt.Sprintf("[red]IP Address: [yellow]%15s[white]\n", state.GetIPAddress().String())

	ipText := tview.NewTextView().
		SetDynamicColors(true)
	fmt.Fprint(ipText, ipString)
	ipFrame := tview.NewFrame(ipText).
		AddText("Connection Information", true, tview.AlignLeft, tcell.ColorGreen)

	// Show buttons
	buttonForm := tview.NewForm().
		AddButton("Start Game", func() {
			logic.CreateGame()
			CreateInGameInterface()
			Pages.SwitchToPage("inGamePage")
		}).
		AddButton("Cancel", func() {
			logic.DestroyPreLobby()
			Pages.SwitchToPage("mainMenuPage")
			settingsMenuPage.Clear()
			preGameLobbyPage.Clear()
		})

	buttFormWidth := 10 + 6 + 2*4 + 1
	buttFormHeight := 3

	buttonFlex := Center(buttFormWidth, buttFormHeight, buttonForm)

	uiGrid := tview.NewGrid().
		SetRows(-1, -3, -1).
		SetColumns(-1, -3, -1)

	uiGrid.AddItem(logoFlex, 0, 1, 1, 1, 0, 0, false)
	uiGrid.AddItem(infoFrame, 1, 0, 1, 1, 0, 0, false)
	uiGrid.AddItem(connFlex, 1, 1, 1, 1, 0, 0, false)
	uiGrid.AddItem(ipFrame, 1, 2, 1, 1, 0, 0, false)
	uiGrid.AddItem(buttonFlex, 2, 1, 1, 1, 0, 0, true)

	preGameLobbyPage.AddItem(uiGrid, 0, 1, true)
}

// CreateInGameInterface creates the ui for when in game
func CreateInGameInterface() {
	m := state.GetSelectedMode()
	// Get logo
	logoWidth, logoHeight, logoBox := LogoArt(m.Name)
	logoFlex := Center(logoWidth, logoHeight, logoBox)

	gameGrid := tview.NewGrid().
		SetRows(-1, -3, -1).
		SetColumns(-1, -3, -1).
		SetBorders(true).
		AddItem(logoFlex, 0, 1, 1, 1, 0, 0, false)

	inGamePage.AddItem(gameGrid, 0, 1, true)
}

// CreatePostGameInterface creates the ui for the postgame lobby
func CreatePostGameInterface() {

}
