package main

import (
	"errors"
	"log"
	"net"
	"os"
	"time"
)

const (
	PORT     = 24601           // The port to run the server on
	IP       = "127.0.0.1"     // The IP to run the server on
	UDPTYPE  = "udp"           // The type of udp to use
	BUFSIZE  = 1024            // The size of the buffer for incoming packets
	DOWNTIME = 1 * time.Second // The amount of time to wait between client sendings
)

// Fails fatally if an error is present
func FailIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Modified from https://gist.github.com/jniltinho/9787946
func WhatsMyIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("No IP found")
}
