package test

import (
	"os"
	"testing"

	fs "github.com/nanoDFS/Slave/filesystem"
)

func TestCreateAndDelete(t *testing.T) {
	fileSystem := fs.NewFileSystem("./test_root")
	os.RemoveAll("./test_root")
	os.Mkdir("./test_root", 0755)
	opts := fs.FileOpts{FileId: "some-random-file-id", ChunkId: "0"}
	file, err := fileSystem.Create(opts)
	if err != nil {
		t.Errorf("failed to create file: %v", err)
	} else {
		file.Close()
	}
	err = fileSystem.Delete(opts)
	if err != nil {
		t.Errorf("failed to delete: %v", err)
	}
}
