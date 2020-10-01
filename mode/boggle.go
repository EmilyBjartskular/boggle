package boggle

// Size defines the board size, 0 being width and 1 being height.
type Size [2]int32

// MatchScore defines the scores for matchlength.
type MatchScore map[int32]int32

// ValidMatch defines if a matched string if characters if a valid match.
type ValidMatch func(match string)

type board struct {
	size       Size
	matchScore MatchScore
	validMatch ValidMatch
}

// NewBoggleBoard returns a board for Boggle mode
func NewBoggleBoard() {

}

// NewFoggleBoard return a board for Foggle mode
func NewFoggleBoard() {

}
