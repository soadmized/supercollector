package text

import (
	"bufio"
	"context"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"

	"supercollector/internal/user"
)

const (
	chanBuffer   = 1000
	workersCount = 10
)

type UserService interface {
	WriteUsers(ctx context.Context, users []user.User) error
	GetUsers(ctx context.Context, filter user.Filter) ([]user.User, error)
}

type Parser struct {
	root        string
	UserService UserService
	rows        chan string
	filePaths   chan string
}

func New(root string) Parser {
	return Parser{
		root:      root,
		rows:      make(chan string, chanBuffer),
		filePaths: make(chan string, chanBuffer),
	}
}

func (p *Parser) Parse(ctx context.Context) error {
	err := filepath.Walk(p.root, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			p.filePaths <- path
		}

		return nil
	})
	if err != nil {
		return err
	}

	close(p.filePaths)

	wg := sync.WaitGroup{}
	wg.Add(workersCount)

	for w := 0; w < workersCount; w++ {
		go func() {
			defer wg.Done()

			p.processFiles(ctx)
		}()
	}

	wg.Wait()

	close(p.rows)

	return nil
}

func (p *Parser) Process(ctx context.Context) error {
	for row := range p.rows {

		// do something with rows
		log.Print(row)
	}

	return nil
}

func (p *Parser) processFiles(ctx context.Context) {
	for path := range p.filePaths {
		if err := p.processFile(ctx, path); err != nil {
			log.Print("reading file", err)
		}
	}
}

func (p *Parser) processFile(_ context.Context, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()
		p.rows <- row
	}

	return nil
}
