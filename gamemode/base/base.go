package base

import (
	"fmt"
	"reflect"
)

// Size defines the board size, index 0 being width and 1 being height.
type Size [2]int32

// Match defines a sequence of characters
// the player has given as input during a game session
type Match string

// MatchScore defines the scoresteps for match lengths
type MatchScore map[uint32]int32

// Die defines a die with 6 sides of a generic type
type Die [6]interface{}

// Dice represents the die face distribution
type Dice []Die

// IsValidMatch defines if a matched string if characters if a valid match.
type IsValidMatch func(match Match) bool

// GameMode define a boggle gamemode
type GameMode struct {
	Name        string
	DieType     reflect.Kind
	CreateBoard BoardConstructor

	// Currently only used for gamemodes
	// where DieType is set to reflect.Int
	SupportedSizes []Size
}

// GameModeConstructor defines how constructors for a GameMode should look.
type GameModeConstructor func() *GameMode

// Board represents gamemode in an active game
type Board struct {
	Size         Size
	MatchScore   MatchScore
	Dice         Dice
	IsValidMatch IsValidMatch
}

// BoardConstructor blash
type BoardConstructor func() *Board

// String returns a string representation of Size
func (s Size) String() string {
	return fmt.Sprintf("%dx%d", s[0], s[1])
}
