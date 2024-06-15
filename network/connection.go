package network

type IConnection interface {
	IsAuthenticated() bool
}

type PlayerConnection struct {
	IConnection
}

func (conn *PlayerConnection) IsAuthenticated() bool {
	return true
}
