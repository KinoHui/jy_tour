package main

import (
	"log"

	"github.com/jyhdashuaibi/jy_tour/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
