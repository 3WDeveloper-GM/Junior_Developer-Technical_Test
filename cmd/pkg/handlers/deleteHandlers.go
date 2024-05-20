package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h *Handler) SingleBillDELETE(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	err = h.bills.Delete(int(n))
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	message := map[string]interface{}{
		"Confirmed": true,
		"Message":   fmt.Sprintf("Succesfully erased the bill with id number %d", n),
	}

	render.JSON(w, r, message)
}
