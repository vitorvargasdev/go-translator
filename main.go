package main

import (
	"gotranslator/src/cli"
	"log"
	"os"
)

func main() {
	cli := cli.Prepare()

	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
