package state

type GameState int8

const (
	StateMainMenu GameState = iota
	StateInGame   GameState = iota
	StatePostGame GameState = iota
)
