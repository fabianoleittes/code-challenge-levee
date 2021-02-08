package presenter

import (
	"reflect"
	"testing"
	"time"

	"github.com/fabianoleittes/code-challenge-levee/domain"
	"github.com/fabianoleittes/code-challenge-levee/usecase/output"
)

func Test_createJobPresenter_Output(t *testing.T) {
	type args struct {
		job domain.Job
	}
	tests := []struct {
		name string
		args args
		want output.Job
	}{
		{
			name: "Create job output",
			args: args{
				job: domain.NewJob(
					"68de685e-3a37-431f-ba6b-dcd0076e5138",
					"2",
					"Jr Ruby Dev",
					"draft",
					"1",
					time.Time{},
					time.Time{},
				),
			},
			want: output.Job{
				ID:         "68de685e-3a37-431f-ba6b-dcd0076e5138",
				PartnerID:  "2",
				Title:      "Jr Ruby Dev",
				Status:     "draft",
				CategoryID: "1",
				ExpiresAt:  "0001-01-01T00:00:00Z",
				CreatedAt:  "0001-01-01T00:00:00Z",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pre := NewJobPresenter()
			if got := pre.Output(tt.args.job); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%+v' | Want: '%+v'", tt.name, got, tt.want)
			}
		})
	}
}
