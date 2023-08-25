package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()
	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(ctx context.Context) {
		defer wg.Done()
		gracefulShutdown(ctx)
	}(ctxTimeout)

	wg.Wait()
	fmt.Println("finishing main goroutine")
}

func gracefulShutdown(ctx context.Context) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("context timeout exceeded")
			return
		case <-signalChan:
			fmt.Println("finishing the process because it was interrupted")
			return
		case <-time.After(2 * time.Second):
			fmt.Println("ay!")
		}
	}
}
