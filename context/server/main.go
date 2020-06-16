package main

import "fmt"
import "log"
import "net/http"
import "time"

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:5050", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("handler started")
	defer log.Println("handler ended")

	ctx := r.Context()
	select {
	// send a response after 3 seconds
	case <-time.After(3 * time.Second):
		fmt.Fprintln(w, "hello John")
		// if client cancels request, then log it
	case <-ctx.Done():
		log.Println("request canceled")
	}
}
