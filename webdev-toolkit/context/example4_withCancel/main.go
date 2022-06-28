package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	//WithCancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //make sure all paths cancel the context to avoid context leak

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel() //call to cancel! signals ctx.Done()
			break
		}
	}

	time.Sleep(1 * time.Minute)
}

//retuns a receiver chan
func gen(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				return //avoid leaking of this goroutine when ctx is done.
			case ch <- n:
				//putting in the channel
				n++
			}
		}
	}()

	return ch
}
