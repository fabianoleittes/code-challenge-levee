package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/fabianoleittes/code-challenge-levee/domain"
	"github.com/fabianoleittes/code-challenge-levee/usecase/input"
	"github.com/fabianoleittes/code-challenge-levee/usecase/output"
)

type mockJobRepoStore struct {
	domain.JobRepository
	result domain.Job
	err    error
}

func (m mockJobRepoStore) Create(_ context.Context, _ domain.Job) (domain.Job, error) {
	return m.result, m.err
}

type mockJobPresenterStore struct {
	output.JobPresenter
	result output.Job
}

func (m mockJobPresenterStore) Output(_ domain.Job) output.Job {
	return m.result
}

func TestCreateJobInteractor_Execute(t *testing.T) {
	t.Parallel()

	expiresAt, _ := time.Parse(time.RFC3339, "2021-02-20T15:04:05Z")

	type args struct {
		input input.Job
	}

	tests := []struct {
		name          string
		args          args
		repository    domain.JobRepository
		presenter     output.JobPresenter
		expected      output.Job
		expectedError interface{}
	}{
		{
			name: "Create Job successful",
			args: args{
				input: input.Job{
					PartnerID:  "1",
					Title:      "Software engineer",
					Status:     "draft",
					CategoryID: "2",
					ExpiresAt:  expiresAt,
				},
			},
			repository: mockJobRepoStore{
				result: domain.NewJob(
					"68de685e-3a37-431f-ba6b-dcd0076e5138",
					"1",
					"Software engineer",
					"draft",
					"2",
					expiresAt,
					time.Now(),
				),
				err: nil,
			},
			presenter: mockJobPresenterStore{
				result: output.Job{
					ID:         "68de685e-3a37-431f-ba6b-dcd0076e5138",
					PartnerID:  "1",
					Title:      "Software engineer",
					Status:     "draft",
					CategoryID: "2",
					ExpiresAt:  expiresAt.Local().String(),
					CreatedAt:  time.Time{}.String(),
				},
			},
			expected: output.Job{
				ID:         "68de685e-3a37-431f-ba6b-dcd0076e5138",
				PartnerID:  "1",
				Title:      "Software engineer",
				Status:     "draft",
				CategoryID: "2",
				ExpiresAt:  expiresAt.Local().String(),
				CreatedAt:  time.Time{}.String(),
			},
		},
		{
			name: "Create Job successful",
			args: args{
				input: input.Job{
					PartnerID:  "2",
					Title:      "JS developer",
					Status:     "draft",
					CategoryID: "3",
					ExpiresAt:  time.Now(),
				},
			},
			repository: mockJobRepoStore{
				result: domain.NewJob(
					"02ae7bf9-09b3-429f-a8da-a03d80940c3b",
					"2",
					"JS developer",
					"draft",
					"3",
					expiresAt,
					time.Now(),
				),
				err: nil,
			},
			presenter: mockJobPresenterStore{
				result: output.Job{
					ID:         "02ae7bf9-09b3-429f-a8da-a03d80940c3b",
					PartnerID:  "2",
					Title:      "JS developer",
					Status:     "draft",
					CategoryID: "3",
					ExpiresAt:  expiresAt.Local().String(),
					CreatedAt:  time.Time{}.String(),
				},
			},
			expected: output.Job{
				ID:         "02ae7bf9-09b3-429f-a8da-a03d80940c3b",
				PartnerID:  "2",
				Title:      "JS developer",
				Status:     "draft",
				CategoryID: "3",
				ExpiresAt:  expiresAt.Local().String(),
				CreatedAt:  time.Time{}.String(),
			},
		},
		{
			name: "Create Job generic error",
			args: args{
				input: input.Job{
					PartnerID:  "",
					Title:      "",
					Status:     "draft",
					CategoryID: "",
					ExpiresAt:  time.Now(),
				},
			},
			repository: mockJobRepoStore{
				result: domain.Job{},
				err:    errors.New("error"),
			},
			presenter: mockJobPresenterStore{
				result: output.Job{},
			},
			expectedError: "error",
			expected:      output.Job{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var uc = NewCreateJobInteractor(tt.repository, tt.presenter, time.Second)

			result, err := uc.Execute(context.TODO(), tt.args.input)
			if (err != nil) && (err.Error() != tt.expectedError) {
				t.Errorf("[TestCase '%s'] Result: '%v' | ExpectedError: '%v'", tt.name, err, tt.expectedError)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.name, result, tt.expected)
			}
		})
	}
}
