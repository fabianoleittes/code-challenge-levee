package domain

import (
	"reflect"
	"testing"
	"time"
)

func TestNewJob(t *testing.T) {

	t.Parallel()

	type args struct {
		ID         JobID
		PartnerID  string
		Title      string
		Status     string
		CategoryID string
		ExpiresAt  time.Time
		CreatedAt  time.Time
	}

	tests := []struct {
		name     string
		args     args
		expected Job
	}{
		{
			name: "Create Job instance",
			args: args{
				ID:         "",
				PartnerID:  "",
				Title:      "",
				Status:     "",
				CategoryID: "",
				ExpiresAt:  time.Time{},
				CreatedAt:  time.Time{},
			},
			expected: Job{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewJob(tt.args.ID, tt.args.PartnerID, tt.args.Title,
				tt.args.Status, tt.args.CategoryID, tt.args.ExpiresAt,
				tt.args.CreatedAt)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.name, result, tt.expected)
			}
		})
	}
}
