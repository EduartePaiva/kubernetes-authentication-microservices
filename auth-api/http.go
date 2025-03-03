package main

import "net/http"

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr}
}

func (h *httpServer) Run() error {
	router := http.NewServeMux()
	return http.ListenAndServe(h.addr, router)
}
