package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	// Get our IP address
	myIP, err := WhatsMyIP()
	FailIf(err)

	// Bind to a port
	ServerAddr, err := net.ResolveUDPAddr(UDPTYPE, myIP+":"+strconv.Itoa(PORT))
	FailIf(err)

	// Listen on the port
	ServerConn, err := net.ListenUDP(UDPTYPE, ServerAddr)
	FailIf(err)

	fmt.Printf("Running a server on %s\n", ServerAddr.String())
	fmt.Println("Waiting for packets")
	// read in input
	buffer := make([]byte, BUFSIZE)
	for {
		fmt.Println("Waiting for next packet")
		// read input
		packetLen, retAddr, err := ServerConn.ReadFromUDP(buffer)
		_ = retAddr // don't respond yet
		FailIf(err)

		// get packet
		packet := buffer[:packetLen]

		fmt.Printf("Got a packet: '%s'\n", string(packet))
	}

}
