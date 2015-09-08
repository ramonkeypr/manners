package test_helpers

import "net"

type Conn struct {
	net.Conn
	closeCalled bool
	localAddr   net.Addr
}

func (f *Conn) LocalAddr() net.Addr {
	return &net.IPAddr{}
}

func (c *Conn) Close() error {
	c.closeCalled = true
	return nil
}
