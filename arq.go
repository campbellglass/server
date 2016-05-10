// A library for performing Automatically Repeated Requests

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
	success := false
	// send request and wait for correct response
	for success {
		// send request
		select {
		case <-responseChannel:
			// check for response correctness
			fmt.Println("got response")
			// if correct, indicate success
			success = true
		case <-time.After(TIMEOUT):
			// if timeout, try again
			fmt.Println("got timeout")
		}
	}
	return
}
