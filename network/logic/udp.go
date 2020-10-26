package logic

import (
	"io/ioutil"
	"log"
	"net"
	"runtime"

	"github.com/EmilyBjartskular/boggle/network"
)

const (
	udpHost = "0.0.0.0"
	udpPort = 53243
	udpType = "udp4"

	conReq = "Looking for Servers!"
)

var (
	udpListeners int = 0
)

// UDPServer manages udp packages
func UDPServer() {
	addr := net.UDPAddr{
		Port: udpPort,
		IP:   net.ParseIP(udpHost),
	}
	conn, err := net.ListenUDP(udpType, &addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	maxListeners := runtime.NumCPU()

	for {
		select {
		case <-network.UDPChannel:
			ioutil.WriteFile("udplog", []byte("Stop me."), 0644)
			return
		default:
			for udpListeners < maxListeners {
				go udpListen(conn)
			}
		}
	}
}

func udpListen(conn *net.UDPConn) {
	udpListeners++
	buffer := make([]byte, 1024)
	n, remoteAddr, err := 0, new(net.UDPAddr), error(nil)
	for err == nil {
		n, remoteAddr, err = conn.ReadFromUDP(buffer)
		// you might copy out the contents of the packet here, to
		// `var r myapp.Request`, say, and `go handleRequest(r)` (or
		// send it down a channel) to free up the listening
		// goroutine. you do *need* to copy then, though,
		// because you've only made one buffer per listen().
		go udpHandleRequest(remoteAddr, string(buffer[:n]))
	}
	udpListeners--
}

func udpHandleRequest(addr *net.UDPAddr, msg string) {

}
