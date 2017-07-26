package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func initFileWatcher() func() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)

			}

		}
	}()

	files := [...]string{"hello.go", "hello", "file_watching.go"}

	for _, name := range files {
		err = watcher.Add(name)
		if err != nil {
			log.Fatal(err)
		}
	}

	return func() {
		err := watcher.Close()
		if err != nil {
			log.Fatal("Failed to close watcher.")
		}
	}
}
