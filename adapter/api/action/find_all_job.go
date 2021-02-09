package action

import (
	"net/http"

	"github.com/fabianoleittes/code-challenge-levee/adapter/api/logging"
	"github.com/fabianoleittes/code-challenge-levee/adapter/api/response"
	"github.com/fabianoleittes/code-challenge-levee/adapter/logger"
	"github.com/fabianoleittes/code-challenge-levee/usecase"
)

type FindAllJobAction struct {
	uc  usecase.FindAllJobUseCase
	log logger.Logger
}

func NewFindAllJobAction(uc usecase.FindAllJobUseCase, log logger.Logger) FindAllJobAction {
	return FindAllJobAction{
		uc:  uc,
		log: log,
	}
}

func (a FindAllJobAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "find_all_job"

	output, err := a.uc.Execute(r.Context())
	if err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning job list")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(a.log, logKey, http.StatusOK).Log("success when returning job list")

	response.NewSuccess(output, http.StatusOK).Send(w)
}
