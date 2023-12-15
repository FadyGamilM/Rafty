package cluster

import (
	"fmt"
	"log"
	"net"

	"github.com/FadyGamilM/Rafty/raftycache"
)

type clusterOptions struct {
	leader bool
	addr   string
}

func NewClusterOptions(addr string, isLeader bool) *clusterOptions {
	return &clusterOptions{
		leader: isLeader,
		addr:   addr,
	}
}

type Cluster struct {
	options       *clusterOptions
	cacheInstance *raftycache.RaftyCache
}

func New(opt *clusterOptions, cache *raftycache.RaftyCache) *Cluster {
	return &Cluster{options: opt, cacheInstance: cache}
}

func (c *Cluster) Run(port string) error {
	listener, err := c.CreateListener(port)
	if err != nil {
		return fmt.Errorf("➜ error creating tcp listener on port %v : %v", port, err)
	}
	log.Printf("➜ cluster is up on port : {%v}\n", port)

	// now keep the server waiting for new connections by creating an infinite for loop that accept a connection , then the Accept method is a block method, so it will block untill we have a new connection, and for each connection we spin a go routine to handle this connection
	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("➜ error accepting a new connection : %v", err)
		}

		go c.handle(conn)
	}

}

func (c *Cluster) handle(conn net.Conn) {
	// for each go routine , we should close the connection after finish handling
	defer conn.Close()

	buffer := make([]byte, 1024*4) // buffer of 4_KB

	// instead of reading
	for {
		numOfReadBytes, err := conn.Read(buffer)
		if err != nil {
			log.Printf("➜ error reading the data into the buffer : %v", err)
			break
		}

		request := buffer[:numOfReadBytes]
		log.Println("command received .. ", string(request))
		// Send a response back to the client
		_, err = conn.Write([]byte("Command received\n"))
		if err != nil {
			log.Printf("➜ error sending response to client: %v", err)
			break
		}
	}
}

func (c *Cluster) CreateListener(addr string) (net.Listener, error) {
	return net.Listen("tcp", addr)
}
