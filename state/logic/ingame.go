package logic

import "github.com/EmilyBjartskular/boggle/network"

// CreateGame creates the game lobby
func CreateGame() {
	network.UDPChannel <- "stop"
}
