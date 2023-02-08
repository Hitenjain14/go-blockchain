package main

import "github.com/hitenjain14/go-blockchain/network"

func main() {

	trLocal := network.NewLocalTransport(network.NetAddr("local"))

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}
