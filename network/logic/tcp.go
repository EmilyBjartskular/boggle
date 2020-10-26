package logic

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/EmilyBjartskular/boggle/network"
	"github.com/EmilyBjartskular/boggle/state"
	"github.com/EmilyBjartskular/boggle/ui/api"
)

const (
	tcpHost = "0.0.0.0"
	tcpPort = 53242
	tcpType = "tcp"

	// TCP Command to add player
	plaStr = "pla:"
)

// TCPServer manages tcp connections
func TCPServer() {
	addr := net.TCPAddr{
		Port: tcpPort,
		IP:   net.ParseIP(tcpHost),
	}
	conn, err := net.ListenTCP(tcpType, &addr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	go tcpListen(conn)

	for {
		select {
		case <-network.TCPChannel:
			return
		default:
		}
	}
}

func tcpListen(conn *net.TCPListener) {
	for {
		c, err := conn.AcceptTCP()
		if err != nil {
			return
		}

		go preHandleConnection(c)
	}
}

// Check if the newly established tcp connection
// wants to add a player. If not, close the connection and return
func preHandleConnection(conn *net.TCPConn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		conn.Close()
		return
	}

	msg := string(buffer[:len(buffer)-1])

	if !strings.HasPrefix(msg, plaStr) {
		conn.Close()
		return
	}

	playerName := msg[:len(plaStr)]

	player := &network.Player{
		PlayerName: playerName,
		IPAddress:  conn.RemoteAddr().String(),
		Connection: conn,
	}

	state.PlayerLobby.AddPlayer(player)

	// update ui
	api.UpdatePlayers()

	go handleConnection(player)
}

func handleConnection(player *network.Player) {
	conn := player.Connection
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		conn.Close()
		state.PlayerLobby.RemovePlayer(player.ID)

		// update ui
		api.UpdatePlayers()
		return
	}

	msg := string(buffer[:len(buffer)-1])

	// Handle input
	if strings.HasPrefix(msg, "") {
	}

	log.Println("Client message:", msg)

	handleConnection(player)
}
