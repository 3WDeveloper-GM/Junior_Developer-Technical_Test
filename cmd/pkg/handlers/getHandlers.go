package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/3WDeveloper-GM/billings/cmd/pkg/domain"
	"github.com/3WDeveloper-GM/billings/cmd/pkg/handlers/validator"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h *Handler) WhoAmIGET(w http.ResponseWriter, r *http.Request) {
  user := h.context.ContextGetUser(r)

	message := map[string]interface{}{
		"message": "you are an authenticated user",
    "user":user,
	}

	render.JSON(w, r, message)
}

func (h *Handler) HealthCheckGET(w http.ResponseWriter, r *http.Request) {
	output := struct {
		Message    string `json:"message"`
		PortNumber int    `json:"port_number"`
	}{
		Message:    "Server up",
		PortNumber: h.portNumber,
	}

	render.JSON(w, r, output)
}

func (h *Handler) FetchBillGET(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	user := h.context.ContextGetUser(r)
	bill := &domain.Bill{}

	bill.Provider.ProviderID = user.ProviderID
	bill.Provider.Name = user.Name

	err = h.bills.Fetch(bill, n)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			h.RowsNotFoundResponse(w, r)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	render.JSON(w, r, bill)
}

func (h *Handler) BillsFetchByDateGET(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Query        string `json:"peticion"`
		StartingDate string `json:"fechaInicio"`
		EndingDate   string `json:"fechaFinal"`
		Results      int    `json:"numResultados"`
	}

	input.Query = "Facturas"

	qs := r.URL.Query()

	defaultEndingDate := time.Now().
		Add(24 * time.Hour).Format(time.DateOnly)
	defaultStartDate := time.Now().
		Add(-24 * time.Hour).Format(time.DateOnly)

	input.StartingDate = h.help.ReadString(qs, "startDate", defaultStartDate)
	input.EndingDate = h.help.ReadString(qs, "endDate", defaultEndingDate)

	start, err := h.help.ParseDate(input.StartingDate)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	end, err := h.help.ParseDate(input.EndingDate)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	input.StartingDate, input.EndingDate = start, end
	user := h.context.ContextGetUser(r)

	bills, err := h.bills.DateFetch(input.StartingDate, input.EndingDate, user)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	input.Results = len(bills)

	response := map[string]interface{}{
		"peticion":  input,
		"resultado": bills,
	}

	render.JSON(w, r, response)
}

func (h *Handler) FetchUserByMailGET(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"correo"`
	}

	err := render.DecodeJSON(r.Body, &input)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	valid := validator.NewValidator()
	domain.ValidateEmail(valid, input.Email)

	if !valid.Valid() {
		h.ValidationErrorResponse(w, r, valid.Errors)
		return
	}

	user, err := h.users.Fetch(input.Email)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
	}

	render.JSON(w, r, *user)
}
