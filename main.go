package main

import (
	"os"

	"github.com/MalteKl/modbusbeat/cmd"

	_ "github.com/MalteKl/modbusbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
