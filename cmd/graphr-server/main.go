package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := NewGraphrServer()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello/", server.helloHandler)

	host := "127.0.0.1"
	port := "8000"

	addr := fmt.Sprintf("%s:%s", host, port)

	fmt.Println("Listening on", addr)
	http.ListenAndServe(addr, mux)
}

type GraphrServer struct {
}

func NewGraphrServer() *GraphrServer {
	return &GraphrServer{}
}

func (*GraphrServer) helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "got request %v", r.URL)
}
