package action

import (
	"encoding/json"
	"net/http"

	"github.com/fabianoleittes/code-challenge-levee/adapter/api/logging"
	"github.com/fabianoleittes/code-challenge-levee/adapter/api/response"
	"github.com/fabianoleittes/code-challenge-levee/adapter/logger"
	"github.com/fabianoleittes/code-challenge-levee/adapter/validator"
	"github.com/fabianoleittes/code-challenge-levee/usecase"
	"github.com/fabianoleittes/code-challenge-levee/usecase/input"
)

type CreateJobAction struct {
	uc        usecase.CreateJob
	log       logger.Logger
	validator validator.Validator
}

func NewCreateJobAction(uc usecase.CreateJob, log logger.Logger, v validator.Validator) CreateJobAction {
	return CreateJobAction{
		uc:        uc,
		log:       log,
		validator: v,
	}
}

func (a CreateJobAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "create_job"

	var jobInput input.Job
	if err := json.NewDecoder(r.Body).Decode(&jobInput); err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("error when decoding json")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	if errs := a.validateInput(jobInput); len(errs) > 0 {
		logging.NewError(
			a.log,
			response.ErrInvalidInput,
			logKey,
			http.StatusBadRequest,
		).Log("invalid input")

		response.NewErrorMessage(errs, http.StatusBadRequest).Send(w)
		return
	}

	output, err := a.uc.Execute(r.Context(), jobInput)
	if err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when creating a new job")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(a.log, logKey, http.StatusCreated).Log("success creating job")

	response.NewSuccess(output, http.StatusCreated).Send(w)
}

func (a CreateJobAction) validateInput(input input.Job) []string {
	var msgs []string

	err := a.validator.Validate(input)
	if err != nil {
		for _, msg := range a.validator.Messages() {
			msgs = append(msgs, msg)
		}
	}

	return msgs
}
