package logic

import (
	"github.com/EmilyBjartskular/boggle/network"
	"github.com/EmilyBjartskular/boggle/network/logic"
	"github.com/EmilyBjartskular/boggle/state"
)

// CreatePreLobby creates a pregame lobby
func CreatePreLobby() {
	go logic.TCPServer()
	go logic.UDPServer()

	// Add player to lobby
	player := &network.Player{
		PlayerName: state.GetPlayerName(),
		IPAddress:  "localhost",
		Connection: nil,
	}

	state.PlayerLobby.AddPlayer(player)
}

// DestroyPreLobby stops the tcp and udp servers
func DestroyPreLobby() {
	network.TCPChannel <- "stop"
	network.UDPChannel <- "stop"
}
