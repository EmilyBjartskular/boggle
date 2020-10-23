package mode

// NewBattleMode returns a board for Battle Boggle
func NewBattleMode() *GameMode {
	name := "Battle Boggle"
	result := &GameMode{
		Name: name,
	}
	return result
}
