package network

import (
	"fmt"
	"net"
)

type IListener interface {
	Start()
	Stop()
}

type BasicListener struct {
	IListener
	port    int
	running bool
}

func CreateListener(port int) *BasicListener {
	return &BasicListener{
		port:    port,
		running: true,
	}
}

func (l *BasicListener) Start() {
	address, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", l.port))

	if err != nil {
		return
	}

	socket, err := net.ListenTCP("tcp", address)
	if err != nil {
		return
	}

	go func() {
		for {
			connection, err := socket.AcceptTCP()
			if err != nil {
				continue
			}
		}
	}()

}

func (l *BasicListener) Stop() {

}
