package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	text := flag.String("text", "hello", "text to echo")
	addr := flag.String("addr", ":9100", "addr to listen")

	http.HandleFunc("/echo", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("%s", *text))
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
