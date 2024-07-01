package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/nautilusgames/demo/internal/server"
)

func main() {
	flags := server.ParseFlags()
	server.Run(flags)
}
