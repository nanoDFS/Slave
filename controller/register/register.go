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
	masterAddr string
	healthAddr string
}

func NewRegister(healthAddr string) RegisterChunkServer {
	addr := config.LoadConfig().Master.Addr
	return RegisterChunkServer{masterAddr: addr, healthAddr: healthAddr}
}

func (t RegisterChunkServer) Register() error {
	conn, err := grpc.NewClient(t.masterAddr, grpc.WithInsecure())
	if err != nil {
		log.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	client := cs_pb.NewChunkServerRegisterServiceClient(conn)
	resp, err := client.Register(context.Background(), &cs_pb.ChunkServerRegisterReq{Address: t.healthAddr, Space: 10000})
	if err != nil || !resp.Status {
		return fmt.Errorf("failed to register: %v", err)
	}
	return nil
}
