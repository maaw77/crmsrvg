package main

import (
	"flag"

	"github.com/maaw77/crmsrvg/internal/crm"
)

func main() {
	var pathConfig string

	flag.StringVar(&pathConfig, "config", "./config/config.yaml", "the path to the configuration file")
	flag.Parse()

	crm.Run(pathConfig)

}
