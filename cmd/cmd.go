package cmd

import (
	"context"

	"github.com/cockroachdb/errors"

	"supercollector/internal/config"
)

const (
	parserCommand = "parse"
	apiCommand    = "api"
)

func Run(ctx context.Context, conf config.Config, args []string) error {
	if len(args) == 0 {
		return errors.Errorf("no command given. possible commands: \n%s\n%s", parserCommand, apiCommand)
	}

	cmd := args[0]

	err := processCommand(ctx, cmd, conf)
	if err != nil {
		return err
	}

	return nil
}

func processCommand(ctx context.Context, cmd string, conf config.Config) error {
	switch cmd {
	case parserCommand:
		return runParser(ctx, conf)
	case apiCommand:
		return runAPI(ctx, conf)
	}

	return errors.Errorf("wrong command given. possible commands: \n%s\n%s", parserCommand, apiCommand)
}
