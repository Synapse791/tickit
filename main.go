package main

import (
	"log"

	"github.com/itmecho/tickit/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
