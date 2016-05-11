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

	// Listen on the port
	Server, err := net.Listen(CONNTYPE, myIP+":"+strconv.Itoa(PORT)) // might need to drop the IP prependiction
	FailIf(err)

	fmt.Printf("Running a server on %s\n", Server.Addr().String())

	for {
		fmt.Printf("Waiting for next connection\n")

		// Accept connection
		ClientConn, err := Server.Accept()
		FailIf(err)
		fmt.Printf("Connected to %s\n", ClientConn.RemoteAddr().String())

		fmt.Println("Waiting for packets")
		// read in input
		buffer := make([]byte, BUFSIZE)
		for {

			// read next incoming packet
			packetLen, err := ClientConn.Read(buffer)
			if err != nil {
				break
			}

			// get packet
			packet := buffer[:packetLen]

			// make response
			toSend := []byte("Thank you for " + string(packet))

			// respond to packet
			fmt.Printf("Got a packet: '%s'\n", string(packet))
			_, err = ClientConn.Write(toSend)
			if err != nil {
				break
			}
		}
		fmt.Printf("Closing connection\n")
		ClientConn.Close()
	}
}
