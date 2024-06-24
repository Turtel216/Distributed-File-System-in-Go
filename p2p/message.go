package p2p

import "net"

// RPC holds represents any arbitrary data that is being send over
// each transport between two nodes in the networks
type RPC struct {
	From    net.Addr
	Payload []byte
}
