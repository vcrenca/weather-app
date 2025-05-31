package v1

import (
	"errors"
	"fmt"
	"github.com/go-chi/render"
	"net/http"
	httpports "weather-api/internal/ports/http"
)

func errHandler(w http.ResponseWriter, r *http.Request, err error) {
	var requiredParamErr *RequiredParamError
	if errors.As(err, &requiredParamErr) {
		badRequest(w, r, fmt.Sprintf("%s query parameter is required", requiredParamErr.ParamName))
		return
	}

	internalServerError(w, r)
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusInternalServerError)
	render.Respond(w, r, httpports.ErrorResponse{
		Title:    "Internal server error",
		Status:   500,
		Instance: r.URL.Path,
	})
}

func badRequest(w http.ResponseWriter, r *http.Request, detail string) {
	render.Status(r, http.StatusBadRequest)
	render.Respond(w, r, httpports.ErrorResponse{
		Title:    "Bad request",
		Status:   400,
		Detail:   detail,
		Instance: r.URL.Path,
	})
}
