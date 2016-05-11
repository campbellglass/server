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

	// Connect to Server from Local
	Conn, err := net.DialUDP(UDPTYPE, LocalAddr, ServerAddr)
	FailIf(err)
	defer Conn.Close()

	// Send data
	buffer := make([]byte, BUFSIZE)
	for i := 0; true; i += 1 {
		// Create request
		toSend := []byte(strconv.Itoa(i))

		// Send request
		fmt.Printf("Sending %s\n", string(toSend))
		Conn.Write(toSend)

		// Wait for response packet
		packetLen, err := Conn.Read(buffer)
		FailIf(err)

		// Handle response packet
		fmt.Printf("Received '%s'\n", string(buffer[:packetLen]))

		// Wait for next request
		time.Sleep(DOWNTIME)
	}
}
