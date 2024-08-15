package handler

import handleriface "myapp/internal/api/handler/handler_iface"

type ApiHandlers struct {
	HelloHandler handleriface.IHelloHandler
}

func InitHandlers() *ApiHandlers {
	helloHandler := NewHelloHandler()

	return &ApiHandlers{
		HelloHandler: helloHandler,
	}
}
