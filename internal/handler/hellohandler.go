package handler

import (
	"fmt"
	"net/http"

	"github.com/irishpatrick/go-web-template/internal/util"
)

type HelloHandler interface {
	Prefix() string
	Router() *http.ServeMux
}

type HelloHandlerImpl struct {
	router *http.ServeMux
	prefix string
}

func (h HelloHandlerImpl) Prefix() string {
	return h.prefix
}

func (h HelloHandlerImpl) Router() *http.ServeMux {
	return h.router
}

func NewHelloHandler() HelloHandler {
	router := http.NewServeMux()
	h := &HelloHandlerImpl{
		prefix: "/hello/",
		router: router,
	}

	router.HandleFunc(util.Get(h, "/test"), h.sayHello)

	return h
}

func (h HelloHandlerImpl) sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, world")
}
