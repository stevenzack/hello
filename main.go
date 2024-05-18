package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	e := http.ListenAndServe(":8080", nil)
	if e != nil {
		log.Panic(e)
		return
	}
}
