package mode

// Size defines the board size, index 0 being width and 1 being height.
type Size [2]int32

// Match defines a sequence of characters
// the player has given as input during a game session
type Match string

// MatchScore defines the scoresteps for match lengths
type MatchScore map[uint32]int32

// Die defines a die with 6 sides of a generic type
type Die [6]interface{}

// DiceList represents the die face distribution
type DiceList []Die

// IsValidMatch defines if a matched string if characters if a valid match.
type IsValidMatch func(match Match) bool

// GameMode define a boggle gamemode
type GameMode struct {
	Name       string
	Size       Size
	MatchScore MatchScore
	DiceList   DiceList

	ValidMatch IsValidMatch
}

// GameModeConstructor defines how constructors for a GameMode should look.
type GameModeConstructor func() *GameMode

var (
	// Modes refers to implemented gamemodes
	Modes []*GameMode
)

// Init initializes all gamemodes
func Init() {
	Modes = []*GameMode{
		NewStandardMode(),
		NewFoggleMode(),
		NewGenerousMode(),
		NewBattleMode(),
	}
}
