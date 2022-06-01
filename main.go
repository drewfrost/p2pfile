package main

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
)

func main() {
	node, err := libp2p.New()
	if err != nil {
		panic(err)
	}

	// print the node's listening addresses
	fmt.Println("Listen addresses:", node.Addrs())

	// shut the node down
	if err := node.Close(); err != nil {
		panic(err)
	}
	fmt.Println("closed")
}
