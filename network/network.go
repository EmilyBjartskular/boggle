package network

import (
	"net"
)

// Player represents a connected player
type Player struct {
	ID         int
	PlayerName string
	IPAddress  string
	Connection *net.TCPConn
}

// Lobby stores connected players
type Lobby struct {
	idCounter int
	Players   map[int]*Player
}

// NewLobby returns a new Lobby object
func NewLobby() *Lobby {
	idCounter := 0
	players := make(map[int]*Player)

	return &Lobby{
		idCounter: idCounter,
		Players:   players,
	}
}

// AddPlayer adds a player to the lobby
func (l *Lobby) AddPlayer(p *Player) {
	p.ID = l.idCounter
	l.Players[l.idCounter] = p
	l.idCounter++
}

// RemovePlayer removes a player from the lobby
func (l *Lobby) RemovePlayer(id int) {
	delete(l.Players, id)
}

// SendMessage sends a message to all players
func (l *Lobby) SendMessage(msg string) {

}

var (
	// TCPChannel is used for ommunication with the tcpserver go routine
	TCPChannel chan string = make(chan string)
	// UDPChannel is used for ommunication with the udpserver go routine
	UDPChannel chan string = make(chan string)
)

// GetManualIP returns the preferred outbout LAN ip.
// If none exists, returns loopback address
func GetManualIP() net.IP {
	conn, err := net.Dial("udp", "1.1.1.1:80")
	if err != nil {
		return net.IPv4(127, 0, 0, 1)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}
