package main

import (
	"fmt"
	"time"
)

func pong(ball chan <- int, action chan <- string) {
	ball <- 2
	action <- "PONG HA GOLPEADO LA PELOTA."
}

func ping(ball chan <- int, action chan <- string) {
	ball <- 1
	action <- "PING HA GOLPEADO LA PELOTA."
}

func referee(action <- chan string) {
	for {
		fmt.Println("Action: ", <- action)
	}
}

func pingpong() {
	ball := make(chan int)
	action := make(chan string)

	go referee(action)
	go ping(ball, action)

	for {
		switch <- ball {
			case 1:
				go pong(ball, action)

			case 2:
				go ping(ball, action)
		}
	}

}

func main() {
	go pingpong()

	time.Sleep(10 * time.Second)
}