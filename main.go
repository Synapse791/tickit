package main

import (
	"log"

	"github.com/Synapse791/tickit/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
