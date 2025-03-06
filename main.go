package main

import (
	"fmt"

	"github.com/nanoDFS/Slave/controller/register"
	"github.com/nanoDFS/Slave/server"
)

func main() {

	monitorAddr, streamingAddr := ":9800", ":8000"

	registerer := register.NewRegister(monitorAddr, streamingAddr)
	if err := registerer.Register(); err != nil {
		fmt.Printf("%v", err)
	}

	monitor := server.NewChunkServerRunner(monitorAddr, streamingAddr)
	monitor.Listen()

	select {}
}
