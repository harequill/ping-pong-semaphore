package main

import (
	"fmt"
	"time"
)

func ping(out chan<- string, pingTurn <-chan struct{}, pongTurn chan<- struct{}) {
	//struct for semaphores is the convention in go, cause takes 0 bytes.
	for {
		<-pingTurn
		out <- "ping"
		time.Sleep(time.Second * 2)
		pongTurn <- struct{}{}
	}
}

func pong(out chan<- string, pongTurn <-chan struct{}, pingTurn chan<- struct{}) {
	for {
		<-pongTurn
		out <- "pong"
		time.Sleep(time.Second * 2)
		pingTurn <- struct{}{}
	}
}

func main() {
	out := make(chan string)
	pingTurn := make(chan struct{})
	pongTurn := make(chan struct{})

	go ping(out, pingTurn, pongTurn)
	go pong(out, pongTurn, pingTurn)

	pingTurn <- struct{}{} //starts the semaphore with "ping"

	for {
		fmt.Println(<-out)
	}
}
