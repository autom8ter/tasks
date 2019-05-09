package proxy

import (
	"github.com/autom8ter/tasks/pkg/functions"
	"github.com/gorilla/mux"
	"net/http"
)

type Proxy struct {
	router *mux.Router
}

func NewProxy(fns ...functions.RouterFunc) *Proxy {
	r := mux.NewRouter()
	for _, f := range fns {
		f(r)
	}
	return &Proxy{
		router: r,
	}
}

func (p *Proxy) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, p.router)
}
