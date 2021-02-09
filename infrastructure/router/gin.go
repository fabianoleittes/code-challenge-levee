package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fabianoleittes/code-challenge-levee/adapter/api/action"
	"github.com/fabianoleittes/code-challenge-levee/adapter/logger"
	"github.com/fabianoleittes/code-challenge-levee/adapter/presenter"
	"github.com/fabianoleittes/code-challenge-levee/adapter/repository"
	"github.com/fabianoleittes/code-challenge-levee/adapter/validator"
	"github.com/fabianoleittes/code-challenge-levee/usecase"
	"github.com/gin-gonic/gin"
)

type ginEngine struct {
	router     *gin.Engine
	log        logger.Logger
	db         repository.NoSQL
	validator  validator.Validator
	port       Port
	ctxTimeout time.Duration
}

func newGinServer(
	log logger.Logger,
	db repository.NoSQL,
	validator validator.Validator,
	port Port,
	t time.Duration,
) *ginEngine {
	return &ginEngine{
		router:     gin.New(),
		log:        log,
		db:         db,
		validator:  validator,
		port:       port,
		ctxTimeout: t,
	}
}

func (g ginEngine) Listen() {
	gin.SetMode(gin.ReleaseMode)
	gin.Recovery()

	g.setAppHandlers(g.router)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", g.port),
		Handler:      g.router,
	}

	g.log.WithFields(logger.Fields{"port": g.port}).Infof("Starting HTTP Server")
	if err := server.ListenAndServe(); err != nil {
		g.log.WithError(err).Fatalln("Error starting HTTP server")
	}
}

func (g ginEngine) setAppHandlers(router *gin.Engine) {
	router.POST("/v1/jobs", g.buildCreateJobAction())
	router.GET("/v1/jobs", g.buildFindAllJobAction())
	router.GET("/v1/health", g.healthcheck())
}

func (g ginEngine) buildCreateJobAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = usecase.NewCreateJobInteractor(
				repository.NewJobNoSQL(g.db),
				presenter.NewJobPresenter(),
				g.ctxTimeout,
			)
			act = action.NewCreateJobAction(uc, g.log, g.validator)
		)

		act.Execute(c.Writer, c.Request)
	}
}

func (g ginEngine) buildFindAllJobAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = usecase.NewFindAllJobInteractor(
				repository.NewJobNoSQL(g.db),
				presenter.NewFindAllJobPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindAllJobAction(uc, g.log)
		)

		act.Execute(c.Writer, c.Request)
	}
}

func (g ginEngine) healthcheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}
