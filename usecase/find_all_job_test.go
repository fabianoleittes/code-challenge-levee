package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/fabianoleittes/code-challenge-levee/domain"
)

type mockJobRepoFindAll struct {
	domain.JobRepository

	result []domain.Job
	err    error
}

func (m mockJobRepoFindAll) FindAll(_ context.Context) ([]domain.Job, error) {
	return m.result, m.err
}

type mockFindAllJobPresenter struct {
	result []FindAllJobOutput
}

func (m mockFindAllJobPresenter) Output(_ []domain.Job) []FindAllJobOutput {
	return m.result
}

func TestFindAllJobInteractor_Execute(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		repository    domain.JobRepository
		presenter     FindAllJobPresenter
		expected      []FindAllJobOutput
		expectedError interface{}
	}{
		{
			name: "Success when returning the Job list",
			repository: mockJobRepoFindAll{
				result: []domain.Job{
					domain.NewJob(
						"68de685e-3a37-431f-ba6b-dcd0076e5138",
						"1",
						"Software engineer",
						"draft",
						"2",
						time.Time{},
						time.Time{},
					),
					domain.NewJob(
						"02ae7bf9-09b3-429f-a8da-a03d80940c3b",
						"2",
						"JS developer",
						"draft",
						"3",
						time.Time{},
						time.Time{},
					),
				},
				err: nil,
			},
			presenter: mockFindAllJobPresenter{
				result: []FindAllJobOutput{
					{
						ID:         "68de685e-3a37-431f-ba6b-dcd0076e5138",
						PartnerID:  "1",
						Title:      "Software engineer",
						Status:     "draft",
						CategoryID: "2",
						ExpiresAt:  time.Time{}.String(),
						CreatedAt:  time.Time{}.String(),
					},
					{
						ID:         "02ae7bf9-09b3-429f-a8da-a03d80940c3b",
						PartnerID:  "2",
						Title:      "JS developer",
						Status:     "draft",
						CategoryID: "3",
						ExpiresAt:  time.Time{}.String(),
						CreatedAt:  time.Time{}.String(),
					},
				},
			},
			expected: []FindAllJobOutput{
				{
					ID:         "68de685e-3a37-431f-ba6b-dcd0076e5138",
					PartnerID:  "1",
					Title:      "Software engineer",
					Status:     "draft",
					CategoryID: "2",
					ExpiresAt:  time.Time{}.String(),
					CreatedAt:  time.Time{}.String(),
				},
				{
					ID:         "02ae7bf9-09b3-429f-a8da-a03d80940c3b",
					PartnerID:  "2",
					Title:      "JS developer",
					Status:     "draft",
					CategoryID: "3",
					ExpiresAt:  time.Time{}.String(),
					CreatedAt:  time.Time{}.String(),
				},
			},
		},
		{
			name: "Success when returning the empty Job list",
			repository: mockJobRepoFindAll{
				result: []domain.Job{},
				err:    nil,
			},
			presenter: mockFindAllJobPresenter{
				result: []FindAllJobOutput{},
			},
			expected: []FindAllJobOutput{},
		},
		{
			name: "Error when returning the list of Jobs",
			repository: mockJobRepoFindAll{
				result: []domain.Job{},
				err:    errors.New("error"),
			},
			presenter: mockFindAllJobPresenter{
				result: []FindAllJobOutput{},
			},
			expectedError: "error",
			expected:      []FindAllJobOutput{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var uc = NewFindAllJobInteractor(tt.repository, tt.presenter, time.Second)

			result, err := uc.Execute(context.Background())
			if (err != nil) && (err.Error() != tt.expectedError) {
				t.Errorf("[TestCase '%s'] Result: '%v' | ExpectedError: '%v'", tt.name, err, tt.expectedError)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.name, result, tt.expected)
			}
		})
	}
}
