package filestreamer

import (
	"net"

	"github.com/charmbracelet/log"
	fst "github.com/nanoDFS/Slave/server/filestreamer/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	fst.UnimplementedFileStreamingServiceServer
}

type FileStreamingServer struct {
	Addr     net.Addr
	listener *net.Listener
	server   *grpc.Server
}

func NewFileStreamingServerRunner(addr string) (*FileStreamingServer, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()
	fst.RegisterFileStreamingServiceServer(s, Server{})
	reflection.Register(s)
	return &FileStreamingServer{
		Addr:     listener.Addr(),
		listener: &listener,
		server:   s,
	}, nil
}

func (t *FileStreamingServer) Listen() error {
	go func() {
		log.Infof("started file chunk streaming service, listening on port: %s", t.Addr)
		if err := t.server.Serve(*t.listener); err != nil {
			log.Fatalf("failed to listen on port %s", t.Addr)
		}
	}()
	return nil
}

func (t *FileStreamingServer) Stop() {
	t.server.Stop()
}
