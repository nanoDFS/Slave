package filestreamer

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

func ReadMetadata(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("no metadata found")
	}
	token, ok := md[key]
	if !ok || len(token) <= 0 {
		return "", fmt.Errorf("no auth token found")
	}
	return token[0], nil
}
