package filesystem

import (
	"fmt"
	"os"
	"path"

	"github.com/nanoDFS/Slave/utils/crypto"
)

type FileOpts struct {
	FileId  string
	ChunkId string
}

func NewFile(fileId string, chunkId string) FileOpts {
	return FileOpts{
		fileId, chunkId,
	}
}

func (t FileOpts) ID() string {
	return fmt.Sprintf("%s-%s", t.FileId, t.ChunkId)
}

type FileSystem struct {
	root string
}

func NewFileSystem(root string) FileSystem {
	return FileSystem{root: root}
}

func (t FileSystem) fullPath(filePath string) string {
	return path.Join(t.root, fmt.Sprintf("%s.bin", filePath))
}

func (t FileSystem) generatePath(file FileOpts) string {
	hashKey := crypto.HashSHA256(file.ID())
	return t.fullPath(hashKey)
}

func (t FileSystem) Exists(fileOpts FileOpts) bool {
	_, err := os.Stat(t.generatePath(fileOpts))
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (t FileSystem) Create(fileOpts FileOpts) (*os.File, error) {
	filePath := t.generatePath(fileOpts)
	if t.Exists(fileOpts) {
		return nil, fmt.Errorf("file already exists")
	}
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (t FileSystem) Open(fileOpts FileOpts) (*os.File, error) {
	filePath := t.generatePath(fileOpts)
	if !t.Exists(fileOpts) {
		return nil, fmt.Errorf("file doesn't exists")
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (t FileSystem) Delete(fileOpts FileOpts) error {
	filePath := t.generatePath(fileOpts)
	if !t.Exists(fileOpts) {
		return fmt.Errorf("file doesn't exists")
	}
	return os.Remove(filePath)
}
