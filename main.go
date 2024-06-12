package main

import (
	"io/fs"
	"log"
	"path/filepath"
	"time"

	"supercollector/internal/config"
)

func main() {
	start := time.Now()
	conf := config.Read()
	log.Print(conf)

	root := conf.PathRoot
	ch := make(chan string)

	go goroutine1Print(ch)
	go goroutine1Print(ch)
	go goroutine1Print(ch)

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			ch <- path
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	close(ch)

	log.Print("DURATION = ", start.Sub(time.Now()))
}

func goroutine1Print(ch chan string) {
	for path := range ch {
		log.Print(path)
	}
}
