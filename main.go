package main

import (
	"flag"
	"fmt"
	"os"
	"go_boilerplate/config"
	"go_boilerplate/data_adapters"
	"go_boilerplate/logger"
	"go_boilerplate/server"
)

func main() {
	environment := flag.String("e", "DEV", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment, false)
	logger.Init()
	data_adapters.Init()
	server.Init()
}