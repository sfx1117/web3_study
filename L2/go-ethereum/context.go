package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")
	select {
	case <-time.After(time.Second * 10):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		serverError := http.StatusInternalServerError
		http.Error(w, err.Error(), serverError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

//server: hello handler started
//server: context canceled
//server: hello handler ended
//server: hello handler started
//server: hello handler ended
