package filestreamer

import (
	fs "github.com/nanoDFS/Slave/server/filestreamer/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t Server) Read(req *fs.ReadReq, stream grpc.ServerStreamingServer[fs.Payload]) error {
	return status.Errorf(codes.Unimplemented, "method Read not implemented")
}
