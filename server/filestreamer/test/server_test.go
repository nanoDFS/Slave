package test

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/nanoDFS/Slave/controller/auth/acl"
	fs "github.com/nanoDFS/Slave/server/filestreamer"
	fs_pb "github.com/nanoDFS/Slave/server/filestreamer/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestWrite(t *testing.T) {
	port := ":8080"
	serverRunner, err := fs.NewFileStreamingServerRunner(port)
	if err != nil {
		t.Error(err)
	}
	serverRunner.Listen()

	conn, err := grpc.NewClient(port, grpc.WithInsecure())
	if err != nil {
		log.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	token, err := acl.NewJWT().Generate(&acl.Claims{UserId: "random", FileId: "some-random", Mode: acl.Write, Size: 10278})

	md := metadata.Pairs(
		"auth", string(token),
		"chunk_id", "0",
	)

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	client := fs_pb.NewFileStreamingServiceClient(conn)
	stream, _ := client.Write(ctx)

	reader := bytes.NewReader([]byte("This is one more file "))
	buff := make([]byte, 1024)
	for {
		n, _ := reader.Read(buff)
		if err != nil && err.Error() != "EOF" {
			log.Fatalf("failed to read data: %v", err)
		}
		if n == 0 {
			break
		}
		if err := stream.Send(&fs_pb.Payload{Data: buff[:n]}); err != nil {
			log.Fatalf("failed to send payload: %v", err)
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to receive response: %v", err)
	}
	log.Printf("Stream closed successfully: %v\n", resp)
}

func TestRead(t *testing.T) {
	port := ":8080"
	serverRunner, err := fs.NewFileStreamingServerRunner(port)
	if err != nil {
		t.Error(err)
	}
	serverRunner.Listen()

	conn, err := grpc.NewClient(port, grpc.WithInsecure())
	if err != nil {
		log.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	token, _ := acl.NewJWT().Generate(&acl.Claims{UserId: "random", FileId: "some-random", Mode: acl.Read, Size: 10278})

	md := metadata.Pairs(
		"auth", string(token),
		"chunk_id", "0",
	)

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	client := fs_pb.NewFileStreamingServiceClient(conn)
	stream, _ := client.Read(ctx, &fs_pb.ReadReq{})

	var buff bytes.Buffer
	for {
		payload, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("failed to read data: %v", err)
		}
		buff.Write(payload.Data)
	}
	log.Info("sucessfully recieved file")
	log.Info(buff.String())
}
