package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) genericErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message interface{}) {
	envelope := map[string]interface{}{
		"status": statusCode,
		"error":  message,
	}

	render.Status(r, statusCode)
	render.JSON(w, r, envelope)
}

func (h *Handler) InternalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.genericErrorResponse(w, r, http.StatusInternalServerError, err.Error())
}

func (h *Handler) NotFoundErrorResponse(w http.ResponseWriter, r *http.Request) {
	h.genericErrorResponse(w, r, http.StatusNotFound, "Not Found")
}

func (h *Handler) NotAllowedErrorResponse(w http.ResponseWriter, r *http.Request) {
	h.genericErrorResponse(w, r, http.StatusMethodNotAllowed, fmt.Sprintf("%s Method not allowed", r.Method))
}

func (h *Handler) ValidationErrorResponse(w http.ResponseWriter, r *http.Request, messages interface{}) {
	h.genericErrorResponse(w, r, http.StatusBadRequest, messages)
}

func (h *Handler) InvalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	h.genericErrorResponse(w, r, http.StatusUnauthorized, "invalid authorization/authentication credentials")
}
