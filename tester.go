package main

import (
	"fmt"
	"os"
	"plugin"

	"github.com/fsnotify/fsnotify"
)

type Machine interface {
	DoSomething(name string)
}

func main() {
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()

	watcher.Add("./libFolder")

	//
	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fmt.Println(event.Op)
				fmt.Println(event.Name)
				loadLib(event.Name)

				// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	<-done
}

func loadLib(path string) {
	p, err := plugin.Open("./" + path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	symbol, err := p.Lookup("Machine")

	fmt.Println("Lookup successfull")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var machine Machine
	machine, e := symbol.(Machine)
	if !e {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	machine.DoSomething("john")
}
