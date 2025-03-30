package http

import "net/http"

type Router struct {
	*http.ServeMux
}

func NewRouter() *Router {
	return &Router{http.NewServeMux()}
}

func (r *Router) AddGroup(basePath string, handler http.Handler) {
	r.Handle(basePath+"/", http.StripPrefix(basePath, handler))
}