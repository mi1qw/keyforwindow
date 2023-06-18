package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task was stopped")
			return
		default:
			fmt.Println("Task is running")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go longRunningTask(ctx)

	time.Sleep(5 * time.Second)
	cancel()

	time.Sleep(2 * time.Second)
	fmt.Println("Main function was stopped")
}
