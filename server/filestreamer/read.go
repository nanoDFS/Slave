package filestreamer

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/nanoDFS/Slave/controller/auth"
	"github.com/nanoDFS/Slave/filesystem"
	fs "github.com/nanoDFS/Slave/server/filestreamer/proto"
)

func (t Server) Read(req *fs.ReadReq, stream fs.FileStreamingService_ReadServer) error {
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
	file, err := fileSystem.Open(filesystem.FileOpts{FileId: claim.FileId, ChunkId: ChunkId})
	if err != nil || read(stream, file) != nil {
		return fmt.Errorf("failed to read from file: %v", err)
	}

	log.Info("Successfully read from ", "file", file.Name())

	defer file.Close()
	return nil
}

func read(stream fs.FileStreamingService_ReadServer, file *os.File) error {
	buff := make([]byte, 1024)
	for {
		n, err := file.Read(buff)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("failed to send payload: %v", err)
		}

		if err := stream.Send(&fs.Payload{Data: buff[:n]}); err != nil {
			return fmt.Errorf("failed to send chunk: %v", err)
		}
	}
	return nil
}
