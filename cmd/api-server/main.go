package main

import (
	"flag"

	"github.com/maaw77/crmsrvg/internal/crm"
)

func main() {
	var pathConfig string

	flag.StringVar(&pathConfig, "config", "", "the path to the configuration file")
	flag.Parse()
	crm.Run(pathConfig)

}
