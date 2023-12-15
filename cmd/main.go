package main

import (
	"github.com/FadyGamilM/Rafty/cluster"
	"github.com/FadyGamilM/Rafty/raftycache"
)

func main() {
	rafty := raftycache.New()
	opts := cluster.NewClusterOptions(":3000", true)
	leader := cluster.New(opts, rafty)
	leader.Run(":3000")

}
