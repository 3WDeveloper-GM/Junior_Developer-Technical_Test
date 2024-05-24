package app

import (
	"errors"
	"fmt"
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

		value, err := r.Cookie("Bearer")
		if err != nil {
			fmt.Println(err)
		}

    if value != nil {
      fmt.Println(value.Value)
    }

		app.Logger.Info().Interface("request information", struct {
			Method string `json:"method"`
			Path   string `json:"path"`
			Auth   string `json:"authHeader"`
		}{
			Method: r.Method,
			Path:   r.URL.Path,
			Auth:   r.Header.Get("Authorization"),
		}).
			Msg(message)

		c := httptest.NewRecorder()
		next.ServeHTTP(c, r)

		for k, v := range c.Result().Header {
			w.Header()[k] = v
		}
		w.WriteHeader(c.Code)
		_, err = c.Body.WriteTo(w)
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
			r = app.Dependencies.Context.ContextSetUser(r, domain.AnonUser)
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

		if auth.ValidateTokenLength(v, currentToken); !v.Valid() {
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

		r = app.Dependencies.Context.ContextSetUser(r, user)

		next.ServeHTTP(w, r)
	})
}

func (app *Application) RequireUserAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := app.Dependencies.Context.ContextGetUser(r)

		if user.IsAnonymous() {
			app.Dependencies.Handlers.AuthenticationFailedResponse(w, r)
			return
		}

		if !user.Activated {
			app.Dependencies.Handlers.ActivationNeededResponse(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *Application) RequirePermissions(code string, next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		user := app.Dependencies.Context.ContextGetUser(r)

		permissions, err := app.Dependencies.Models.Permits.GetPermissionsFromUser(user.SysID)
		if err != nil {
			app.Dependencies.Handlers.InternalServerErrorResponse(w, r, err)
			return
		}

		if !permissions.Include(code) {
			app.Dependencies.Handlers.NotPermittedResponse(w, r)
			return
		}

		next.ServeHTTP(w, r)
	}

	return app.RequireUserAuth(fn)
}
