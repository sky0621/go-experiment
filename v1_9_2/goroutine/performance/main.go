package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	//goroutineNum = 100 * 10000
	goroutineNum = 50 * 10000
	//goroutineNum = 10 * 10000
	//goroutineNum = 1 * 10000
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	for i := 0; i < goroutineNum; i++ {
		wg.Add(1)
		go func(ctx context.Context, wg *sync.WaitGroup, i int) {
			fmt.Printf("[No.%d]Start\n", i)
			time.Sleep(5 * time.Second)
			fmt.Printf("[No.%d]End\n", i)
			wg.Done()
			ctx.Done()
		}(ctx, &wg, i)
	}

	wg.Wait()
}
