package main

import (
	"fmt"

	"github.com/nanoDFS/Slave/controller/register"
	"github.com/nanoDFS/Slave/server/monitor"
)

func main() {
	registerer := register.NewRegister(":9800")
	if err := registerer.Register(); err != nil {
		fmt.Printf("%v", err)
	}

	monitor, _ := monitor.NewMonitorServerRunner(":9800")
	monitor.Listen()

	select {}
}
