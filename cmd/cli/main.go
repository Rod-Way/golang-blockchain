package main

import (
	"golang-blockchain/internal/cli"
	"golang-blockchain/internal/core"
)

func main() {
	bc := core.NewBlockchain()
	defer bc.DB.Close()

	cli := cli.CLI{bc}
	cli.Run()
}
