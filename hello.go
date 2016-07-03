package main

import "log"
import "github.com/fsnotify/fsnotify"


func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer wather.Close()

	done := make(chan tool)
	go func(){
		for{
			select {
			case event := <- watcher.Events:
				log.Println("Event:", event)
				if event.Op&fsnotify.Write === fsnotify.Write {
					log Println("modified file: ", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("Error: ", err)
			}
		}
		}()
		err = watcher.Add("/tmp/foo")
		if err != nil {
			log.Fatal(err)
		}
		<-done
	}
