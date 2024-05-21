package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/3WDeveloper-GM/billings/cmd/pkg/context"
	"github.com/3WDeveloper-GM/billings/cmd/pkg/handlers"
	"github.com/3WDeveloper-GM/billings/cmd/pkg/handlers/validator"
	"github.com/3WDeveloper-GM/billings/cmd/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const portnumber = 4040

type Application struct {
	Server       *chi.Mux
	Dependencies *dependency
	Config       *Config
	Logger       *zerolog.Logger
}

type dependency struct {
	Handlers handlers.Handler
	Models   models.AppModels
	Context  context.ContextKey
	Valid    validator.Validator
}

func (a *Application) setDependencies() {

  newCtx := context.NewContext()

	mods := models.InitializeAppModels(a.Config.DB)
	handler := handlers.NewHandlerInstance(
    portnumber, mods.Bills, 
    mods.Users, mods.Tokens,
    mods.Permits, newCtx,
    )


	depends := &dependency{
		Handlers: *handler,
		Models:   *mods,
		Context:  *newCtx,
	}

	a.Dependencies = depends
}

func (a *Application) setServer() {
	a.Server = chi.NewRouter()
	a.Server.Use(a.VisitedRouteLogger)
  a.Server.Use(a.Authenticate)

	a.Server.NotFound(a.Dependencies.Handlers.NotFoundErrorResponse)
	a.Server.MethodNotAllowed(a.Dependencies.Handlers.NotAllowedErrorResponse)
}

func NewApplication() *Application {
	a := &Application{}
	a.setLogger()

	err := a.setConfiguration()
	if err != nil {
		a.Logger.Panic().Msg(err.Error())
	}
	err = a.setDBPool()
	if err != nil {
		a.Logger.Panic().Msg(err.Error())
	}

	a.setDependencies()
	a.setServer()

	a.setRoutes()
	return a
}

func (a *Application) setLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	a.Logger = &log.Logger

	a.Logger.Info().Msg("Logger configured")
}

func (a *Application) StartApp() {
	serverUrl := "Starting server in http://localhost:%d"
	a.Logger.Info().Msg(fmt.Sprintf(serverUrl, portnumber))

	server := http.Server{
		Addr:        fmt.Sprintf(":%d", portnumber),
		ReadTimeout: 3 * time.Second,
		IdleTimeout: 3 * time.Second,
		Handler:     a.Server,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
