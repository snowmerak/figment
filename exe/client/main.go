package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/lucas-clemente/quic-go/http3"
)

func main() {
	client := http.Client{
		Transport: &http3.RoundTripper{},
	}
	resp, err := client.Get("https://localhost:8080")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(io.ReadAll(resp.Body))
}
