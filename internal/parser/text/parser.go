package text

import (
	"context"
	"io/fs"
	"path/filepath"

	"supercollector/internal/user"
)

type UserService interface {
	WriteUsers(ctx context.Context, users []user.User) error
	GetUsers(ctx context.Context, filter user.Filter) ([]user.User, error)
}

type Parser struct {
	root        string
	batchSize   int
	UserService UserService
}

func New(root string, batchSize int) Parser {
	return Parser{
		root:      root,
		batchSize: batchSize,
	}
}

func (p *Parser) Parse(ctx context.Context) error {
	filesToRead := make([]string, 0)
	walkFn := func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			filesToRead = append(filesToRead, path)
		}

		return nil
	}

	err := filepath.Walk(p.root, walkFn)
	if err != nil {
		return err
	}

	// worker pool

	return nil
}

// processBatch read the files batch and writes content to DB.
func (p *Parser) processFile(ctx context.Context, pathToFile string) error {
	// 1) open and read
	// 2) string to user.User
	// 3) write to user service

	return nil
}
