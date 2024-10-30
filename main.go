package main

import (
	"net/http"
)

var (
	port = "c1.play2go.cloud:22768"
)

func main() { //go run main.go
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ЦУУФУУ"))
	})

	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
