package main

import (
	"log"
	"net/http"
)

func main() {
	_, err := http.Get("http://127.0.0.1:8888/notify")
	if err != nil {
		log.Println(err.Error())
	}
}
