package main

import (
	"fmt"
	"net/http"

	"github.com/lucas-clemente/quic-go/http3"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello, world!"))
		w.WriteHeader(200)
	})
	fmt.Println("Listening on :8080")
	if err := http3.ListenAndServeQUIC("0.0.0.0:8080", "./localhost/cert.pem", "./localhost/key.pem", nil); err != nil {
		panic(err)
	}
}
