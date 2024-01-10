package main

import (
	"fmt"
	"strconv"
	"time"
)

var quitChannels = make(map[string]chan bool)

func main() {
	for i := 0; i < 5; i++ {
		id := strconv.Itoa(i)
		startWorker(id)
	}
	time.Sleep(5 * time.Second)
	quit := quitChannels["0"]
	quit <- true
	fmt.Println("Notified worker 0 to quit")
	select {}
}

func startWorker(id string) {
	quit := make(chan bool)
	go worker(id, quit)
	quitChannels[id] = quit
}

func worker(id string, quit chan bool) {
	for {
		select {
		case <-quit:
			fmt.Println("Quitting worker ", id)
			return
		default:
			time.Sleep(10 * time.Second)
		}
	}
}
