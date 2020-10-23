package mode

// NewFoggleMode returns a board for Foggle Boggle
func NewFoggleMode() *GameMode {
	name := "Foggle Boggle"
	result := &GameMode{
		Name: name,
	}
	return result
}
