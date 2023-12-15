package client

import (
	"fmt"
	"log"
	"net"
	"time"
)

func ConnectToRaftyCluster(done <-chan bool, addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("error trying to connect to rafty cluster : %v \n", err)
		return
	}
	idx := 0
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			_, err := conn.Write([]byte(fmt.Sprintf("SET ID %v", idx)))
			if err != nil {
				log.Printf("error trying to send command to rafty cluster : %v \n", err)
				return
			}
			idx++
		}
	}
}
