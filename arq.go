// A library for performing Automatic Reqpeat Requests

package main

import (
	"fmt"
	"time"
)

const (
	// Timeout before resending the response
	TIMEOUT = 5 * time.Second
)

// Test code goes here? Move it to a separate file later
func main() {
	fmt.Println("Running ARQ code")
}

// Sends a request on a channel, then listens for a response
// If no correct response is received before timing out, resends the response
// If a correct response is received before timing out, exits successfully
func arq(responseChannel chan bool) {
	// send request
	success := false
	for success {
		select {
		case <-responseChannel:
			fmt.Println("got response")
			success = true
		case <-time.After(TIMEOUT):
			fmt.Println("got timeout")
		}
	}
	return
}
