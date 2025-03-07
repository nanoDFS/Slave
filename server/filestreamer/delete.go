package filestreamer

import (
	"context"
	"fmt"

	"github.com/nanoDFS/Slave/controller/auth"
	"github.com/nanoDFS/Slave/filesystem"
	fs "github.com/nanoDFS/Slave/server/filestreamer/proto"
)

func (t Server) Delete(ctx context.Context, req *fs.DeleteReq) (*fs.DeleteRes, error) {
	token, err := ReadMetadata(ctx, "auth")
	claim, ok := auth.NewAuth().AuthorizeDelete(token)

	if err != nil || !ok {
		return nil, fmt.Errorf("failed to authorize %v", err)
	}

	ChunkId, err := ReadMetadata(ctx, "chunk_id")
	if err != nil {
		return nil, fmt.Errorf("invalid chunk id")
	}

	fileSystem := filesystem.NewFileSystem("./test_root")
	err = fileSystem.Delete(filesystem.FileOpts{FileId: claim.FileId, ChunkId: ChunkId})
	if err != nil {
		return nil, err
	}
	return &fs.DeleteRes{Status: true}, nil
}
