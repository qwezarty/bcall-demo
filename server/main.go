package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var sig = make(chan struct{})

func main() {
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/notify", handleNotify)

	log.Println("server started, listen at 8888")
	http.ListenAndServe(":8888", nil)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	select {
	case <-sig:
		fmt.Fprintln(w, "hello, world!")
	case <-time.After(15 * time.Second):
		fmt.Fprintln(w, "oops, timing out!")
	}
}

func handleNotify(w http.ResponseWriter, r *http.Request) {
	select {
	case <-sig:
		log.Println("signal already sent")
	default:
		close(sig)
		log.Println("sending signal")
	}
}
