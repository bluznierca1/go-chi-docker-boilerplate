package handler

import "net/http"

type HelloHandler struct{}

func (h *HelloHandler) Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}
