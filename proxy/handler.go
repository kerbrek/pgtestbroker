package proxy

import (
	"github.com/kerbrek/pgtestbroker/message"
)

type MessageHandler func(*Ctx, []byte) (message.Reader, error)

type MessageHandlerRegister interface {
	GetHandler(byte) MessageHandler
}
