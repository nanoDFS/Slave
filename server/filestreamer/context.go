package filestreamer

import (
	"fmt"

	fs "github.com/nanoDFS/Slave/server/filestreamer/proto"
	"google.golang.org/grpc/metadata"
)

func ReadMetadata(stream fs.FileStreamingService_WriteServer, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return "", fmt.Errorf("no metadata found")
	}
	token, ok := md[key]
	if !ok || len(token) <= 0 {
		return "", fmt.Errorf("no auth token found")
	}
	return token[0], nil
}
