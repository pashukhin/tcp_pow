package server

import "io"

type ServerHandler func(rw io.ReadWriter)

type Server interface {
	Start()
	Stop()
}
