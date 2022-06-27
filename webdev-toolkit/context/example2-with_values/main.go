package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
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

	// warning not to use built-in string for the key; create my own type
	userID := contextKey("userID")
	fName := contextKey("fName")

	ctx = context.WithValue(ctx, userID, 777)
	ctx = context.WithValue(ctx, fName, "Bond")

	results := dbAccess(ctx, userID)

	fmt.Println(w, results)

}

func dbAccess(ctx context.Context, key contextKey) int {
	uid := ctx.Value(key).(int) //assert

	return uid
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
