package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	fmt.Printf("Running a client\n")
	// Get our IP address
	myIP, err := WhatsMyIP()
	FailIf(err)

	fmt.Printf("Connecting to %s:%d\n", myIP, PORT)

	// Bind to server
	ServerAddr, err := net.ResolveUDPAddr(UDPTYPE, myIP+":"+strconv.Itoa(PORT))
	FailIf(err)

	// Bind locally
	LocalAddr, err := net.ResolveUDPAddr(UDPTYPE, ":0")
	FailIf(err)

	// Connect
	Conn, err := net.DialUDP(UDPTYPE, LocalAddr, ServerAddr)
	FailIf(err)
	defer Conn.Close()

	// Send data
	i := 0
	for {
		toSend := []byte(strconv.Itoa(i))
		fmt.Printf("Sending %s\n", string(toSend))
		Conn.Write(toSend)
		time.Sleep(DOWNTIME)
		i += 1
	}

}
