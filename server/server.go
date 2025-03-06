package server

import (
	fs "github.com/nanoDFS/Slave/server/filestreamer"
	ms "github.com/nanoDFS/Slave/server/monitor"
)

type ChunkServer struct {
	MS *ms.MonitorServer
	FS *fs.FileStreamingServer
}

func NewChunkServerRunner(monitorAddr string, streamingAddr string) *ChunkServer {
	monitoringServer, _ := ms.NewMonitorServerRunner(monitorAddr)
	streamingServer, _ := fs.NewFileStreamingServerRunner(streamingAddr)

	return &ChunkServer{
		MS: monitoringServer,
		FS: streamingServer,
	}
}

func (t *ChunkServer) Listen() error {
	err := t.MS.Listen()
	if err != nil {
		return err
	}
	err = t.FS.Listen()
	if err != nil {
		return err
	}
	return nil
}

func (t *ChunkServer) Stop() {
	t.MS.Stop()
	t.FS.Stop()
}
