package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func CreateServer(port string, createHandler func(router chi.Router) http.Handler) http.Server {
	rootRouter := chi.NewRouter()
	rootRouter.Mount("/api", createHandler(createApiRouter()))

	return http.Server{
		Addr:    ":" + port,
		Handler: rootRouter,
	}
}

func createApiRouter() *chi.Mux {
	apiRouter := chi.NewRouter()
	apiRouter.Use(middleware.RequestID)
	apiRouter.Use(logMiddleware)
	apiRouter.Use(middleware.Recoverer)
	return apiRouter
}
