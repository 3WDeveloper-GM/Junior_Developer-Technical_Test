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
	h.genericErrorResponse(w, r, http.StatusMethodNotAllowed, fmt.Sprintf("%s method not allowed", r.Method))
}

func (h *Handler) ValidationErrorResponse(w http.ResponseWriter, r *http.Request, messages interface{}) {
	h.genericErrorResponse(w, r, http.StatusBadRequest, messages)
}

func (h *Handler) InvalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	h.genericErrorResponse(w, r, http.StatusUnauthorized, "invalid authorization/authentication credentials")
}

func (h *Handler) InvalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")
	message := "invalid or missing authentication token"
	h.genericErrorResponse(w, r, http.StatusUnauthorized, message)
}

func (h *Handler) RowsNotFoundResponse(w http.ResponseWriter, r *http.Request) {
  h.genericErrorResponse(w,r,http.StatusNotFound,"Resource not found.")
}

func (h *Handler) ActivationNeededResponse(w http.ResponseWriter, r *http.Request) {
  h.genericErrorResponse(w,r,http.StatusUnauthorized, "user account must be activated in order to access this resource")
}

func (h *Handler) AuthenticationFailedResponse(w http.ResponseWriter, r *http.Request) {
  h.genericErrorResponse(w,r,http.StatusUnauthorized, "you must be authenticated to access this resource")
}

func (h *Handler) NotPermittedResponse(w http.ResponseWriter, r *http.Request) {
  h.genericErrorResponse(w,r,http.StatusUnauthorized,"your user account doesn't have the permissions to access this resource")
}

func (h *Handler) NotAuthenticatedResponse(w http.ResponseWriter, r *http.Request) {
  h.genericErrorResponse(w,r,http.StatusUnauthorized,"User account not authenticated")
}
