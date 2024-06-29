package cmd

import (
	"context"
	"log"

	"supercollector/internal/builder"
	"supercollector/internal/config"
)

func runAPI(ctx context.Context, conf config.Config) error {
	go func() {
		err := builder.NewGRPCServer(ctx)
		if err != nil {
			log.Print(err)
		}
	}()

	err := builder.NewRESTServer(ctx)
	if err != nil {
		log.Print(err)
	}

	return nil
}
