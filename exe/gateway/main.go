package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":11211",
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
