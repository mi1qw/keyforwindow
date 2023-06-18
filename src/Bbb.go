package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {
	for n := 0; ; n++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%v - %s - %d \n",
			robotgo.GetPid(),
			robotgo.GetTitle(), n)
	}
}
