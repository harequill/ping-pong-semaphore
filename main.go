package main

import (
	"fmt"
	"time"
)

func ping(out chan<- string) {
	for {
		out <- "ping"
		time.Sleep(time.Second * 2)
	}
}

func main() {
	out := make(chan string)

	go ping(out)

	for {
		fmt.Println(<-out)
	}

}
