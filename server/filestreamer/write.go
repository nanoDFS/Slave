package filestreamer

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/nanoDFS/Slave/controller/auth"
	"github.com/nanoDFS/Slave/filesystem"
	fs "github.com/nanoDFS/Slave/server/filestreamer/proto"
)

func (t Server) Write(stream fs.FileStreamingService_WriteServer) error {
	token, err := ReadMetadata(stream.Context(), "auth")
	claim, ok := auth.NewAuth().AuthorizeWrite(token)
	if err != nil || !ok {
		return fmt.Errorf("failed to authorize %v", err)
	}

	ChunkId, err := ReadMetadata(stream.Context(), "chunk_id")
	if err != nil {
		return fmt.Errorf("invalid chunk id")
	}

	fileSystem := filesystem.NewFileSystem("./")
	file, err := fileSystem.Create(filesystem.FileOpts{FileId: claim.FileId, ChunkId: ChunkId})
	if err != nil || write(stream, file) != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	log.Info("Successfully writen to ", "file", file.Name())

	defer file.Close()
	return stream.SendAndClose(&fs.WriteRes{
		Status: true,
	})
}

func write(stream fs.FileStreamingService_WriteServer, file *os.File) error {
	for {
		payload, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("failed to recieve payload: %v", err)
		}
		file.Write(payload.Data)
	}
	return nil
}
