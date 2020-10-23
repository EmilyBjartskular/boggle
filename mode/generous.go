package mode

// NewGenerousMode returns a board for Generous Boggle
func NewGenerousMode() *GameMode {
	name := "Generous Boggle"
	result := &GameMode{
		Name: name,
	}
	return result
}
