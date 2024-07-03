package main

import (
	"github.com/nautilusgames/demo/config"
	"github.com/nautilusgames/demo/wallet/internal/server"
)

func main() {
	flags := config.ParseFlags()
	server.Run(flags)
}
