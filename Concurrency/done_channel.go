package main

import (
	"fmt"
	"time"
)

// the done channel is passed in as a read-only channel using the <- after the channel name
func doWork(done <- chan bool){
	for {
		select {
		// this syntax is basically like expecting data from the down channel
		case <- done:
			return
		default:
			fmt.Println("Doing work")
		}
		time.Sleep(time.Second*1)
	}
}

func main(){
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 3)
	close(done)
}