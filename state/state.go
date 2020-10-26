package state

import (
	"net"

	"github.com/EmilyBjartskular/boggle/gamemode/base"
	"github.com/EmilyBjartskular/boggle/network"
	"github.com/rivo/tview"
)

// GameState blash
type GameState int8

// Language represents an available language
type Language struct {
	Name        string
	DiceConfigs map[string]base.Dice
	BoardSizes  []base.Size
}

const (
	// StateMainMenu is gamestate for the main menu
	StateMainMenu GameState = iota
	// StatePreGame is the state for when the user is in the pregame lobby
	StatePreGame GameState = iota
	// StateInGame is the state for when the user is in a match
	StateInGame GameState = iota
	// StatePostGame is the state for when the user is in the postgame lobby
	StatePostGame GameState = iota
)

var (
	// App is the main application of boggle
	App *tview.Application

	// Languages is a list of available languages
	Languages []*Language
)

var (
	playerName       string
	currentState     GameState = StateMainMenu
	selectedMode     *base.GameMode
	selectedLanguage *Language
	selectedSize     base.Size
	ipAddress        net.IP
)

var (
	// PlayerLobby represents connected players
	PlayerLobby *network.Lobby = network.NewLobby()
)

// SetPlayerName set the players name
func SetPlayerName(name string) {
	playerName = name
}

// GetPlayerName returns the players name
func GetPlayerName() string {
	return playerName
}

// SetCurrentState sets the current gamestate
func SetCurrentState(state GameState) {
	currentState = state
}

// GetCurrentState return the current gamestate
func GetCurrentState() GameState {
	return currentState
}

// SetSelectedMode sets which gamemode the user has selected in the mode menu
func SetSelectedMode(m *base.GameMode) {
	selectedMode = m
}

// GetSelectedMode returns the gamemode the user has selected in the mode menu
func GetSelectedMode() *base.GameMode {
	return selectedMode
}

// SetSelectedLanguage sets which language the user has selected in the mode settings menu
func SetSelectedLanguage(l *Language) {
	selectedLanguage = l
}

// GetSelectedLanguage returns the language the user has selected in the mode settings menu
func GetSelectedLanguage() *Language {
	return selectedLanguage
}

// SetSelectedSize sets which boardsize the user has selected in the mode settings menu
func SetSelectedSize(s base.Size) {
	selectedSize = s
}

// GetSelectedSize returns the boardsize the user has selected in the mode settings menu
func GetSelectedSize() base.Size {
	return selectedSize
}

// SetIPAddress sets the LAN ip
func SetIPAddress(ip net.IP) {
	ipAddress = ip
}

// GetIPAddress returns the LAN ip
func GetIPAddress() net.IP {
	return ipAddress
}
