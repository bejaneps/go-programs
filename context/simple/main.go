package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// context will be canceled after n seconds,
	// cancel func should also be called as well
	// for releasing allocated resources
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // even if timer cancel won't get called, we release resources

	// will never get executed because we are canceling context
	sleepAndTalk(ctx, 5*time.Second, "look at me")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println("context cancelled before message printed")
		//log.Println(ctx.Err())
	}
}
