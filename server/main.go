package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// empty struct designed for sending signal
var sig = make(chan struct{})

func main() {
	// block only if notify have been sent
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/notify", handleNotify)

	// it's simple to handle multi-thread task
	go func() {
		time.Sleep(10 * time.Second)
		// blank identifier, equivalent to /bin/null
		_, err := http.Get("http://127.0.0.1:8888/notify")
		if err != nil {
			log.Println(err.Error())
		}
	}()

	log.Fatal(http.ListenAndServe(":8888", nil))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	select { // not switch-case, select is designed for channels
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
