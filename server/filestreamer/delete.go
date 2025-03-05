package filestreamer

import (
	"context"
	"fmt"

	"github.com/nanoDFS/Slave/controller/auth"
	"github.com/nanoDFS/Slave/filesystem"
	fs "github.com/nanoDFS/Slave/server/filestreamer/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t Server) Delete(ctx context.Context, req *fs.DeleteReq) (*fs.DeleteRes, error) {
	token, err := ReadMetadata(ctx, "auth")
	claim, ok := auth.NewAuth().AuthorizeWrite(token)
	if err != nil || !ok {
		return &fs.DeleteRes{Status: false}, fmt.Errorf("failed to authorize")
	}

	ChunkId, err := ReadMetadata(ctx, "chunk_id")
	if err != nil {
		return &fs.DeleteRes{Status: false}, fmt.Errorf("invalid chunk id")
	}

	fileSystem := filesystem.NewFileSystem("./")
	err = fileSystem.Delete(filesystem.FileOpts{FileId: claim.FileId, ChunkId: ChunkId})
	if err != nil {
		return &fs.DeleteRes{Status: false}, fmt.Errorf("failed to delete file")
	}
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
