package cmd

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"supercollector/internal/config"
	"supercollector/internal/parser/text"
)

func runParser(ctx context.Context, conf config.Config) error {
	// parses text files
	parser := text.New(conf.PathRoot)

	eg := errgroup.Group{}

	eg.Go(func() error {
		err := parser.Parse(ctx)
		if err != nil {
			return fmt.Errorf("parse files at %s: %w", conf.PathRoot, err)
		}

		return nil
	})

	eg.Go(func() error {
		err := parser.Process(ctx)
		if err != nil {
			return fmt.Errorf("process rows at %s: %w", conf.PathRoot, err)
		}

		return nil
	})

	err := eg.Wait()
	if err != nil {
		return err
	}

	return nil
}
