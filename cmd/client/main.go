package main

import (
	"time"

	"github.com/FadyGamilM/Rafty/client"
)

func main() {
	done := make(chan bool)
	go client.ConnectToRaftyCluster(done, ":3000")

	// Wait for a signal to stop the client
	time.Sleep(10 * time.Second)
	done <- true
}
