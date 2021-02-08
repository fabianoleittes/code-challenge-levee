package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fabianoleittes/code-challenge-levee/adapter/api/action"
	"github.com/fabianoleittes/code-challenge-levee/adapter/api/middleware"
	"github.com/fabianoleittes/code-challenge-levee/adapter/logger"
	"github.com/fabianoleittes/code-challenge-levee/adapter/presenter"
	"github.com/fabianoleittes/code-challenge-levee/adapter/repository"
	"github.com/fabianoleittes/code-challenge-levee/adapter/validator"
	"github.com/fabianoleittes/code-challenge-levee/usecase"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type gorillaMux struct {
	router     *mux.Router
	middleware *negroni.Negroni
	log        logger.Logger
	db         repository.SQL
	validator  validator.Validator
	port       Port
	ctxTimeout time.Duration
}

func newGorillaMux(
	log logger.Logger,
	db repository.SQL,
	validator validator.Validator,
	port Port,
	t time.Duration,
) *gorillaMux {
	return &gorillaMux{
		router:     mux.NewRouter(),
		middleware: negroni.New(),
		log:        log,
		db:         db,
		validator:  validator,
		port:       port,
		ctxTimeout: t,
	}
}

func (g gorillaMux) Listen() {
	g.setAppHandlers(g.router)
	g.middleware.UseHandler(g.router)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", g.port),
		Handler:      g.middleware,
	}

	g.log.WithFields(logger.Fields{"port": g.port}).Infof("Starting HTTP Server")
	if err := server.ListenAndServe(); err != nil {
		g.log.WithError(err).Fatalln("Error starting HTTP server")
	}
}

func (g gorillaMux) setAppHandlers(router *mux.Router) {
	api := router.PathPrefix("/v1").Subrouter()

	api.Handle("/jobs", g.buildCreateJobAction()).Methods(http.MethodPost)
	api.HandleFunc("/health", action.HealthCheck).Methods(http.MethodGet)
}

func (g gorillaMux) buildCreateJobAction() *negroni.Negroni {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		var (
			uc = usecase.NewCreateJobInteractor(
				repository.NewJobSQL(g.db),
				presenter.NewJobPresenter(),
				g.ctxTimeout,
			)
			act = action.NewCreateJobAction(uc, g.log, g.validator)
		)

		act.Execute(res, req)
	}

	return negroni.New(
		negroni.HandlerFunc(middleware.NewLogger(g.log).Execute),
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}
