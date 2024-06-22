package p2p

import "net"

// Message holds represents any arbitrary data that is being send over
// each transport between two nodes in the networks
type Message struct {
	From    net.Addr
	Payload []byte
}
