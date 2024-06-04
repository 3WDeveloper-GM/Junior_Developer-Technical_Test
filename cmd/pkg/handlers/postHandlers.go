package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/3WDeveloper-GM/billings/cmd/pkg/auth"
	"github.com/3WDeveloper-GM/billings/cmd/pkg/domain"
	jsonbobjects "github.com/3WDeveloper-GM/billings/cmd/pkg/domain/jsonbObjects"
	"github.com/3WDeveloper-GM/billings/cmd/pkg/handlers/validator"
	"github.com/3WDeveloper-GM/billings/cmd/pkg/models"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func (h *Handler) SendJson(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Message string `json:"message"`
		Integer int    `json:"integer"`
	}

	err := render.DecodeJSON(r.Body, &input)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	var output struct {
		OriginalMessage struct {
			Message string `json:"message"`
			Integer int    `json:"integer"`
		} `json:"original"`
		Confirmation string `json:"confirmation"`
	}

	output.OriginalMessage = input
	output.Confirmation = "got a message"

	render.JSON(w, r, output)
}

func (h *Handler) BillingPOST(w http.ResponseWriter, r *http.Request) {
	var data struct {
		BillID       string                   `json:"idFactura"`
		Date         string                   `json:"fechaEmision"`
		TotalAmmount int                      `json:"montoTotal"`
		Details      jsonbobjects.JsonObjects `json:"detalles,omitempty"`
		Misc         jsonbobjects.JsonObjects `json:"miscelaneo,omitempty"`
	}

	err := render.DecodeJSON(r.Body, &data)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	user := h.context.ContextGetUser(r)

	input := &domain.Bill{
		BillID:       data.BillID,
		Date:         data.Date,
		TotalAmmount: &data.TotalAmmount,
		Details:      data.Details,
		Misc:         data.Misc,
	}

	input.Provider.ProviderID = user.ProviderID
	input.Provider.Name = user.Name

	// fmt.Println(*input)

	valid := validator.NewValidator()

	if !input.ValidateBill(valid) {
		h.ValidationErrorResponse(w, r, valid.Errors)
		return
	}

	err = h.bills.Create(input)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	response := map[string]interface{}{
		"Confirmed": true,
		"payload":   input,
	}

	render.JSON(w, r, response)
}

func (h *Handler) UserCreatePOST(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Name     string `json:"nombreUsuario"`
		Email    string `json:"correoUsuario"`
		PassWord string `json:"contraUsuario"`
	}

	err := render.DecodeJSON(r.Body, &data)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	input := &domain.Users{
		ProviderID: uuid.NewString(),
		Name:       data.Name,
		Email:      data.Email,
		Activated:  true,
	}

	err = input.Password.Set(data.PassWord)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	valid := validator.NewValidator()

	if !input.ValidateUser(valid) {
		h.ValidationErrorResponse(w, r, valid.Errors)
		return
	}

	err = h.users.Create(input)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrDuplicateMail):
			valid.AddErrorKey("emails", models.ErrDuplicateMail.Error())
			h.ValidationErrorResponse(w, r, valid.Errors)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	err = h.permissions.GrantPermissionToUser(
		input.SysID,
		h.permissions.GenerateUserPermissions()...)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	response := map[string]interface{}{
		"message": struct {
			Confirmed bool   `json:"confirmed"`
			Info      string `json:"info"`
		}{
			Confirmed: true,
			Info:      "User Created!",
		},
		"payload": input,
	}

	render.JSON(w, r, response)
}

func (h *Handler) CreateAuthTokenPOST(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"correoUsuario"`
		Password string `json:"contraUsuario"`
	}

	err := render.DecodeJSON(r.Body, &input)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	valid := validator.NewValidator()

	domain.ValidateEmail(valid, input.Email)
	domain.ValidatePasswordFromPlaintext(valid, input.Password)

	if !valid.Valid() {
		h.ValidationErrorResponse(w, r, valid.Errors)
		return
	}

	user, err := h.users.Fetch(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRows):
			h.InvalidCredentialsResponse(w, r)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	match, err := user.Password.Matches(input.Password)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	if !match {
		h.InvalidCredentialsResponse(w, r)
		return
	}

	token, err := h.tokens.New(user.SysID, 24*time.Hour, auth.ScopeAuthentication)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	userInfo := map[string]interface{}{
		"usuario": user.Name,
		"id":      user.ProviderID,
	}

	message := map[string]interface{}{
		"created":   true,
		"message":   "login completed successfully",
		"resultado": userInfo,
		"token":     token.Plaintext,
	}

	cookie := http.Cookie{
		Name:     "Bearer",
		Value:    token.Plaintext,
		Expires:  token.Expiry,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
		Path:     "/",
		MaxAge:   24 * 60,
	}

	http.SetCookie(w, &cookie)

	render.JSON(w, r, message)
}

func (h *Handler) CreateAutheticationTokenJWTPOST(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserMail string `json:"correoUsuario"`
		Password string `json:"contraUsuario"`
	}

	err := render.DecodeJSON(r.Body, &input)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	valid := validator.NewValidator()

	domain.ValidateEmail(valid, input.UserMail)
	domain.ValidatePasswordFromPlaintext(valid, input.Password)

	if !valid.Valid() {
		h.ValidationErrorResponse(w, r, valid.Errors)
		return
	}

	user, err := h.users.Fetch(input.UserMail)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRows):
			h.InvalidCredentialsResponse(w, r)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	match, err := user.Password.Matches(input.Password)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	if !match {
		h.InvalidCredentialsResponse(w, r)
		return
	}
}

func (h *Handler) LoginJWTtokenPOST(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := render.DecodeJSON(r.Body, &input)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	valid := validator.NewValidator()

	domain.ValidateEmail(valid, input.Email)
	domain.ValidatePasswordFromPlaintext(valid, input.Password)

	if !valid.Valid() {
		h.ValidationErrorResponse(w, r, valid.Errors)
		return
	}

	user, err := h.users.Fetch(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRows):
			h.InvalidCredentialsResponse(w, r)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	match, err := user.Password.Matches(input.Password)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
	}

	if !match {
		h.InvalidCredentialsResponse(w, r)
		return
	}

	jwtToken, err := h.jwtToken.CreateToken(user.Name, user.Email)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	message := map[string]interface{}{
		"info":      "token created",
		"token":     jwtToken,
		"userName":  user.Name,
		"userEmail": user.Email,
	}

	render.JSON(w, r, message)
}
