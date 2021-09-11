package core

import "fmt"

type InitializationConnection struct {
	network  *Network
	portName string
	fullName string
	closed   bool
	value    string
}

func (c *InitializationConnection) IsEmpty() bool { /* TODO */ return true }

func (c *InitializationConnection) receive(p *Process) *Packet {

	if c.closed {
		return nil
	}
	fmt.Println(p.Name, "Receiving IIP")
	var pkt *Packet = new(Packet)
	pkt.Contents = c.value
	pkt.owner = p
	p.ownedPkts++
	c.closed = true
	fmt.Println(p.Name, "Received IIP: ", pkt.Contents)
	return pkt
}

func (c *InitializationConnection) IsClosed() bool {
	return c.closed
}

func (c *InitializationConnection) ResetClosed() {
	c.closed = false
}

func (c *InitializationConnection) GetType() string {
	return "InitializationConnection"
}
