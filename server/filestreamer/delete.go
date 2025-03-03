package filestreamer

import (
	"context"

	fs "github.com/nanoDFS/Slave/server/filestreamer/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t Server) Delete(ctx context.Context, req *fs.DeleteReq) (*fs.DeleteRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
