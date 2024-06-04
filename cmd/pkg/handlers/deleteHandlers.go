package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func (h *Handler) SingleBillDELETE(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

  log.Info().Msg(id)

  err := h.bills.Delete(id)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	message := map[string]interface{}{
		"Confirmed": true,
		"Message":   fmt.Sprintf("Succesfully erased the bill with id %s", id),
	}

	render.JSON(w, r, message)
}
