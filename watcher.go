package main

import (
	"log"

	"golang.org/x/exp/fsnotify"
)

func watch(pathes []string) error {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case ev := <-w.Event:
				log.Println("event:", ev)
			case err := <-w.Error:
				log.Println("error:", err)
			}
		}
	}()

	for _, path := range pathes {
		if err := w.Watch(path); err != nil {
			return err
		}
	}
	return nil
}
