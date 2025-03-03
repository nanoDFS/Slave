package health

import (
	"net"

	"github.com/nanoDFS/p2p/p2p/transport"
)

type HealthServer struct {
	listenAddr net.Addr
	server     *transport.TCPTransport
	quitChan   chan struct{}
}

func NewHealthServer(addr string) (*HealthServer, error) {
	server, err := transport.NewTCPTransport(addr)
	if err != nil {
		return nil, err
	}

	return &HealthServer{
		listenAddr: server.ListenAddr,
		server:     server,
		quitChan:   make(chan struct{}),
	}, nil
}

func (t *HealthServer) Listen() error {
	return t.server.Listen()
}

func (t *HealthServer) Stop() {
	t.server.Stop()
}
