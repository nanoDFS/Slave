package register

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	cs_pb "github.com/nanoDFS/Slave/controller/register/proto"
	"github.com/nanoDFS/Slave/utils/config"
	"google.golang.org/grpc"
)

type RegisterChunkServer struct {
	masterAddr    string
	monitorAddr   string
	streamingAddr string
}

func NewRegister(monitorAddr string, streamingAddr string) RegisterChunkServer {
	addr := config.LoadConfig().Master.Addr
	return RegisterChunkServer{masterAddr: addr, monitorAddr: monitorAddr, streamingAddr: streamingAddr}
}

func (t RegisterChunkServer) Register() error {
	conn, err := grpc.NewClient(t.masterAddr, grpc.WithInsecure())
	if err != nil {
		log.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	client := cs_pb.NewChunkServerRegisterServiceClient(conn)
	resp, err := client.Register(context.Background(), &cs_pb.ChunkServerRegisterReq{MonitorAddress: t.monitorAddr, StreamingAddress: t.streamingAddr, Space: 10000})
	if err != nil || !resp.Success {
		return fmt.Errorf("failed to register: %v", err)
	}
	return nil
}
