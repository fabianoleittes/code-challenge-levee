package action

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/fabianoleittes/code-challenge-levee/infrastructure/log"
	"github.com/fabianoleittes/code-challenge-levee/usecase"
)

type mockFindAllJob struct {
	result []usecase.FindAllJobOutput
	err    error
}

func (m mockFindAllJob) Execute(_ context.Context) ([]usecase.FindAllJobOutput, error) {
	return m.result, m.err
}

func TestFindAllJobAction_Execute(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name               string
		ucMock             usecase.FindAllJobUseCase
		expectedBody       string
		expectedStatusCode int
	}{
		{
			name: "FindAllJobAction success one Job",
			ucMock: mockFindAllJob{
				result: []usecase.FindAllJobOutput{
					{

						ID:         "68de685e-3a37-431f-ba6b-dcd0076e5138",
						PartnerID:  "2",
						Title:      "Jr golang dev",
						Status:     "draft",
						CategoryID: "1",
						ExpiresAt:  time.Time{}.String(),
						CreatedAt:  time.Time{}.String(),
					},
				},
				err: nil,
			},
			expectedBody: `[{
				"id":"68de685e-3a37-431f-ba6b-dcd0076e5138",
				"partner_id":"2",
				"title":"Jr golang dev",
				"status":"draft",
				"category_id": "1",
				"expires_at":"0001-01-01 00:00:00 +0000 UTC"
				"created_at":"0001-01-01 00:00:00 +0000 UTC"
				}]`,
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "FindAllJobAction success empty",
			ucMock: mockFindAllJob{
				result: []usecase.FindAllJobOutput{},
				err:    nil,
			},
			expectedBody:       `[]`,
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "FindAllJobAction generic error",
			ucMock: mockFindAllJob{
				err: errors.New("error"),
			},
			expectedBody:       `{"errors":["error"]}`,
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/jobs", nil)

			var (
				w      = httptest.NewRecorder()
				action = NewFindAllJobAction(tt.ucMock, log.LoggerMock{})
			)

			action.Execute(w, req)

			if w.Code != tt.expectedStatusCode {
				t.Errorf(
					"[TestCase '%s'] The handler returned an unexpected HTTP status code: returned '% v' expected '% v'",
					tt.name,
					w.Code,
					tt.expectedStatusCode,
				)
			}

			var result = strings.TrimSpace(w.Body.String())
			if !strings.EqualFold(result, tt.expectedBody) {
				t.Errorf(
					"[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tt.name,
					result,
					tt.expectedBody,
				)
			}
		})
	}
}
