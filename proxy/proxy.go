package proxy

import (
	"drasticproxy/network"
)

type IProxy interface {
	Start()
	Stop()
	Init() bool
	GetConnectedPlayers() []network.PlayerConnection
}

type DrasticProxy struct {
	IProxy
	port     int
	players  []network.PlayerConnection
	listener network.BasicListener
}

// GetConnectedPlayers implements IProxy.
// Subtle: this method shadows the method (IProxy).GetConnectedPlayers of DrasticProxy.IProxy.
func (d *DrasticProxy) GetConnectedPlayers() []network.PlayerConnection {
	return d.players
}

// Start implements IProxy.
// Subtle: this method shadows the method (IProxy).Start of DrasticProxy.IProxy.
func (d *DrasticProxy) Start() {
	d.listener.Start()
}

// Stop implements IProxy.
// Subtle: this method shadows the method (IProxy).Stop of DrasticProxy.IProxy.
func (d *DrasticProxy) Stop() {
	panic("unimplemented")
}

func (d *DrasticProxy) Init() bool {
	return false
}

func CreateProxy(port int) IProxy {
	return &DrasticProxy{
		port:     port,
		players:  make([]network.PlayerConnection, 5),
		listener: *network.CreateListener(port),
	}
}
