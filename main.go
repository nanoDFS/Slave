package main

import (
	"fmt"
	"os"

	"github.com/nanoDFS/Slave/controller/register"
	"github.com/nanoDFS/Slave/server"
	"github.com/nanoDFS/Slave/utils"
)

func main() {
	utils.InitLog()
	monitorAddr, streamingAddr := os.Getenv("MONITOR_ADDR"), os.Getenv("STREAMING_ADDR")

	registerer := register.NewRegister(monitorAddr, streamingAddr)
	if err := registerer.Register(); err != nil {
		fmt.Printf("%v", err)
	}

	monitor := server.NewChunkServerRunner(monitorAddr, streamingAddr)
	monitor.Listen()

	select {}
}
