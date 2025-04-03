package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func CreateServer(port string, createHandler func(router chi.Router) http.Handler) http.Server {
	apiRouter := chi.NewRouter()
	apiRouter.Use(middleware.RequestID)
	apiRouter.Use(LogrusMiddleware())
	apiRouter.Use(middleware.Recoverer)

	rootRouter := chi.NewRouter()
	rootRouter.Mount("/api", createHandler(apiRouter))

	return http.Server{
		Addr:    ":" + port,
		Handler: rootRouter,
	}
}
