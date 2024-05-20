package app

import (
	"net/http"
	"net/http/httptest"
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
      app.Dependencies.Handlers.InternalServerErrorResponse(w,r,err)
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
