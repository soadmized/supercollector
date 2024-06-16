package main

import (
	"context"
	"log"
	"os"

	"supercollector/cmd"
	"supercollector/internal/config"
)

func main() {
	ctx := context.Background()
	conf := config.Read()
	args := os.Args[1:]

	err := cmd.Run(ctx, conf, args)
	if err != nil {
		log.Fatal(err)
	}
}
