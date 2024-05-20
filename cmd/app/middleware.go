package app

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/3WDeveloper-GM/billings/cmd/pkg/auth"
	"github.com/3WDeveloper-GM/billings/cmd/pkg/domain"
	"github.com/3WDeveloper-GM/billings/cmd/pkg/handlers/validator"
	"github.com/3WDeveloper-GM/billings/cmd/pkg/models"
)

func (app *Application) VisitedRouteLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message := "got a request with the following:"

		app.Logger.Info().Interface("request information", struct {
			Method string `json:"method"`
			Path   string `json:"path"`
		}{
			Method: r.Method,
			Path:   r.URL.Path,
		}).
			Msg(message)

		c := httptest.NewRecorder()
		next.ServeHTTP(c, r)

		for k, v := range c.Result().Header {
			w.Header()[k] = v
		}
		w.WriteHeader(c.Code)
		_, err := c.Body.WriteTo(w)
		if err != nil {
			app.Dependencies.Handlers.InternalServerErrorResponse(w, r, err)
		}

		message = "sent the following response:"

		if res := c.Result().StatusCode; res <= 300 {
			app.Logger.Info().Interface("response information", struct {
				Status  string      `json:"status"`
				Headers interface{} `json:"headers"`
			}{
				Status:  c.Result().Status,
				Headers: c.Result().Header,
			}).
				Msg(message)
		} else {
			app.Logger.Error().Interface("response information", struct {
				StatusCode string      `json:"status"`
				Headers    interface{} `json:"headers"`
			}{
				StatusCode: c.Result().Status,
				Headers:    c.Result().Header,
			}).
				Msg(message)
		}
	})
}

func (app *Application) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			app.Dependencies.Context.ContextSetUser(r, domain.AnonUser)
			next.ServeHTTP(w, r)
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			app.Dependencies.Handlers.InvalidAuthenticationTokenResponse(w, r)
			return
		}

		currentToken := headerParts[1]

		v := validator.NewValidator()

		if domain.ValidatePasswordFromPlaintext(v, currentToken); !v.Valid() {
			app.Dependencies.Handlers.InvalidAuthenticationTokenResponse(w, r)
			return
		}

		user, err := app.Dependencies.Models.Users.GetForToken(auth.ScopeAuthentication, currentToken)
		if err != nil {
			switch {
			case errors.Is(err, models.ErrNoRows):
				app.Dependencies.Handlers.InvalidAuthenticationTokenResponse(w, r)
			default:
				app.Dependencies.Handlers.InternalServerErrorResponse(w, r, err)
			}
			return
		}

    r = app.Dependencies.Context.ContextSetUser(r,user)

    next.ServeHTTP(w,r)
	})
}
