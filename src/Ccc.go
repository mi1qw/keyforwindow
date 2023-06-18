package main

import (
	"fmt"
)

//пример многопоточки

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			select {
			case ch1 <- i:
				fmt.Println("Sent to ch1:", i)
			case ch2 <- i:
				fmt.Println("Sent to ch2:", i)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		select {
		case val := <-ch1:
			fmt.Println("Received from ch1:", val)
		case val := <-ch2:
			fmt.Println("Received from ch2:", val)
		}
	}

	request := make(chan string)
	response := make(chan string)

	go func() {
		for {
			req := <-request
			response <- "Received request: " + req
		}
	}()

	for i := 0; i < 4; i++ {
		select {
		case request <- "message":
			fmt.Println("Sent message")
		case res := <-response:
			fmt.Println(res)
		}
	}
}
