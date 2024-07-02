package main

import (
	"github.com/nautilusgames/demo/auth/internal/server"
	"github.com/nautilusgames/demo/config"
)

func main() {
	flags := config.ParseFlags()
	server.Run(flags)
}
