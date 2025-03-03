package monitor

import (
	hs "github.com/nanoDFS/Slave/server/monitor/health"
)

type MonitorServer struct {
	HS *hs.HealthServer
}

func NewMonitorServerRunner(addr string) (*MonitorServer, error) {
	healthServer, _ := hs.NewHealthServer(addr)
	return &MonitorServer{
		HS: healthServer,
	}, nil
}

func (t *MonitorServer) Listen() error {
	return t.HS.Listen()
}

func (t *MonitorServer) Stop() {
	t.HS.Stop()
}
