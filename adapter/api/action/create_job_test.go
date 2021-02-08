package action

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/fabianoleittes/code-challenge-levee/infrastructure/log"
	"github.com/fabianoleittes/code-challenge-levee/infrastructure/validation"
	"github.com/fabianoleittes/code-challenge-levee/usecase"
	"github.com/fabianoleittes/code-challenge-levee/usecase/input"
	"github.com/fabianoleittes/code-challenge-levee/usecase/output"
)

type mockJobCreateJob struct {
	result output.Job
	err    error
}

func (m mockJobCreateJob) Execute(_ context.Context, _ input.Job) (output.Job, error) {
	return m.result, m.err
}

func TestCreateJobAction_Execute(t *testing.T) {
	t.Parallel()

	// expiresAt, _ := time.Parse(time.RFC3339, "2021-02-20T15:04:05Z")
	validator, _ := validation.NewValidatorFactory(validation.InstanceGoPlayground)

	type args struct {
		rawPayload []byte
	}

	tests := []struct {
		name               string
		args               args
		ucMock             usecase.CreateJob
		expectedBody       string
		expectedStatusCode int
	}{
		{
			name: "CreateJobAction success",
			args: args{
				rawPayload: []byte(
					`{
						"partner_id": "2",
						"title": "Jr Ruby Dev",
						"category_id": "1",
						"expires_at": "0001-01-01 00:00:00"
					}`,
				),
			},
			ucMock: mockJobCreateJob{
				result: output.Job{
					ID:         "68de685e-3a37-431f-ba6b-dcd0076e5138",
					PartnerID:  "2",
					Title:      "Jr Ruby Dev",
					Status:     "draft",
					CategoryID: "1",
					ExpiresAt:  time.Time{}.String(),
					CreatedAt:  time.Time{}.String(),
				},
				err: nil,
			},
			expectedBody: `{
				"id":"68de685e-3a37-431f-ba6b-dcd0076e5138",
				"partner_id":"2",
				"title":"Jr Ruby Dev",
				"status":"draft",
				"category_id": "1",
				"expires_at":"0001-01-01 00:00:00 +0000 UTC"
				"created_at":"0001-01-01 00:00:00 +0000 UTC"
				}`,
			expectedStatusCode: http.StatusCreated,
		},
		{
			name: "CreateJobAction success",
			args: args{
				rawPayload: []byte(
					`{
						"partner_id": "1",
						"title": "Sr Ruby Dev",
						"category_id": "1",
						"expires_at": "0001-01-01 00:00:00"
					}`,
				),
			},
			ucMock: mockJobCreateJob{
				result: output.Job{
					ID:         "68de685e-3a37-431f-ba6b-dcd0076e5138",
					PartnerID:  "1",
					Title:      "Sr Ruby Dev",
					Status:     "draft",
					CategoryID: "2",
					ExpiresAt:  time.Time{}.String(),
					CreatedAt:  time.Time{}.String(),
				},
				err: nil,
			},
			expectedBody: `{
				"id":"68de685e-3a37-431f-ba6b-dcd0076e5138",
				"partner_id":"1",
				"title":"Sr Ruby Dev",
				"status":"draft",
				"category_id": "2",
				"expires_at":"0001-01-01 00:00:00 +0000 UTC"
				"created_at":"0001-01-01 00:00:00 +0000 UTC"
				}`,
			expectedStatusCode: http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(
				http.MethodPost,
				"/jobs",
				bytes.NewReader(tt.args.rawPayload),
			)

			var (
				w      = httptest.NewRecorder()
				action = NewCreateJobAction(tt.ucMock, log.LoggerMock{}, validator)
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
