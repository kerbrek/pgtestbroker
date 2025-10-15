package proxy

import (
	"context"
	"net"

	"github.com/kerbrek/pgtestbroker/backend"
	"github.com/kerbrek/pgtestbroker/message"
)

type AuthPhase int

const (
	PhaseStartup AuthPhase = iota
	PhaseGSS
	PhaseSASLInit
	PhaseSASL
	PhaseOK
)

type Ctx struct {
	ClientConn     net.Conn
	ServerConn     net.Conn
	ConnInfo       backend.ConnInfo
	RowDescription *message.RowDescription
	AuthPhase      AuthPhase
	Context        context.Context
	Cancel         context.CancelFunc
}
