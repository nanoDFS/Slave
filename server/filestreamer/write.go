package filestreamer

import (
	fs "github.com/nanoDFS/Slave/server/filestreamer/proto"
	"google.golang.org/grpc"
)

func (t Server) Write(stream grpc.ClientStreamingServer[fs.Payload, fs.WriteRes]) error {
	return nil
}
