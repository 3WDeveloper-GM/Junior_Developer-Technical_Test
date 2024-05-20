package handlers

import (
	"net/http"

	"github.com/3WDeveloper-GM/billings/cmd/pkg/domain"
	"github.com/go-chi/render"
)

func (h *Handler) UpdateBillPUT(w http.ResponseWriter, r *http.Request) {
	var input domain.Bill

	err := render.DecodeJSON(r.Body, &input)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	err = h.bills.Update(&input)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	message := map[string]interface{}{
		"message": "Successfully updated!",
		"payload": input,
	}

	render.JSON(w, r, message)
}
