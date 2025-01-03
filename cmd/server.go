package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/irishpatrick/go-web-template/internal/handler"
	"go.uber.org/fx"
)

type Server = *http.Server
type Router = *http.ServeMux

func NewRouter(helloHandler handler.HelloHandler) Router {
	router := http.NewServeMux()

	router.Handle(helloHandler.Prefix(), helloHandler.Router())

	return router
}

func NewServer(lc fx.Lifecycle, router Router) Server {
	host := os.Getenv("HOST")
	if len(host) == 0 {
		host = "localhost"
	}
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

func RunServer() {
	fx.New(
		fx.Provide(
			handler.NewHelloHandler,
			NewRouter,
			NewServer,
		),
		fx.Invoke(func(Server) {}),
	).Run()
}
