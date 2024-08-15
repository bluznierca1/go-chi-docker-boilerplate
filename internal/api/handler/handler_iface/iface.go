package handleriface

import "net/http"

type IHelloHandler interface {
	Hello(w http.ResponseWriter, r *http.Request)
}
