package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type contextKey string

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()

	//contextKeys
	userID := contextKey("userID")
	fname := contextKey("fname")

	ctx = context.WithValue(ctx, userID, 777)
	ctx = context.WithValue(ctx, fname, "Bond")

	results, err := dbAccess(ctx, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context, key contextKey) (int, error) {

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	//making a bidirectional channel
	ch := make(chan int)

	//goroutine & annonymous func
	go func() {
		// ridiculous long running task
		uid := ctx.Value(key).(int) //assertion
		time.Sleep(10 * time.Second)

		//check to make sure we're not running in vain
		// if ctx.Done() has
		if ctx.Err() != nil {
			return
		}

		ch <- uid
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch: //this will not run because there is a timeout
		return i, nil
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
