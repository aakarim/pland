package plan

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestParse(t *testing.T) {
	type args struct {
		ctx      context.Context
		planFile string
	}
	tests := []struct {
		name    string
		args    args
		want    *PlanFile
		wantErr bool
	}{
		{
			name: "empty file",
			args: args{
				ctx:      context.Background(),
				planFile: ``,
			},
			want:    &PlanFile{},
			wantErr: false,
		},
		{
			name: "with only days",
			args: args{
				ctx: context.Background(),
				planFile: `# plan.day/2022-08-14

- [x] Read Part II of PLG Onboarding
`,
			},
			want: &PlanFile{
				Days: []Day{
					{
						Contents: `- [x] Read Part II of PLG Onboarding`,
						Date:     time.Date(2022, 8, 14, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "with only days - no space",
			args: args{
				ctx: context.Background(),
				planFile: `>plan.day/2022-08-14

- [x] Read Part II of PLG Onboarding
`,
			},
			want: &PlanFile{
				Days: []Day{
					{
						Contents: `- [x] Read Part II of PLG Onboarding`,
						Date:     time.Date(2022, 8, 14, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: false,
		},

		{
			name: "with header token",
			args: args{
				ctx: context.Background(),
				planFile: `# plan.header

		Hi! This is my == plan file. 😀✍ ==

		---
		# plan.day/2022-08-14

		- [x] Read Part II of PLG Onboarding
		- [x] Study Org mode
		- [ ] write stringifier for plans
		`,
			},
			want: &PlanFile{
				Header: Header{
					token: `# plan.header`,
					Contents: `Hi! This is my == plan file. 😀✍ ==

		---`,
				},
				Days: []Day{
					{
						Contents: `- [x] Read Part II of PLG Onboarding
		- [x] Study Org mode
		- [ ] write stringifier for plans`,
						Date: time.Date(2022, 8, 14, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "forwards compatability",
			args: args{
				ctx: context.Background(),
				planFile: `# plan.header

		Hi! This is my == plan file. 😀✍ ==

		---
		# plan.project/lix

		- [ ] Read stuff

		---
		# plan.day/2022-08-14

		- [x] Read Part II of PLG Onboarding
		- [x] Study Org mode
		- [ ] write stringifier for plans`,
			},
			want: &PlanFile{
				Header: Header{
					token: `# plan.header`,
					Contents: `Hi! This is my == plan file. 😀✍ ==

		---`,
				},
				Days: []Day{
					{
						Contents: `- [x] Read Part II of PLG Onboarding
		- [x] Study Org mode
		- [ ] write stringifier for plans`,
						Date: time.Date(2022, 8, 14, 0, 0, 0, 0, time.UTC),
					},
				},
				ArbitrarySections: []ArbitrarySection{
					{
						Contents: `- [ ] Read stuff

		---`,
						token: "# plan.project/lix",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.ctx, tt.args.planFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Header.Contents, tt.want.Header.Contents) {
				t.Errorf("Header.Contents = %v, want %v", got.Header.Contents, tt.want.Header.Contents)
			}
			// Uncomment when we use the custom parser
			// if !reflect.DeepEqual(got.Header.token, tt.want.Header.token) {
			// 	t.Errorf("Header.token = %v, want %v", got.Header.token, tt.want.Header.token)
			// }
			// ASSUMPTION: that the slices are laid out in memory in the same way
			for i, d := range got.Days {
				if !reflect.DeepEqual(d.Contents, tt.want.Days[i].Contents) {
					dmp := diffmatchpatch.New()
					diffs := dmp.DiffMain(d.Contents, tt.want.Days[i].Contents, false)

					t.Errorf("Days = %v, want %v; diffs: %v", d.Contents, tt.want.Days[i].Contents, dmp.DiffPrettyText(diffs))
				}
				if !d.Date.Equal(tt.want.Days[i].Date) {
					t.Errorf("Days = %v, want %v", d.Date, tt.want.Days[i].Date)
				}
			}
			if !reflect.DeepEqual(got.ArbitrarySections, tt.want.ArbitrarySections) {
				t.Errorf("ArbitrarySections = %v, want %v", got.ArbitrarySections, tt.want.ArbitrarySections)
			}
		})
	}
}

func TestPlanFile_String(t *testing.T) {
	type fields struct {
		ParentVersion     int
		Header            Header
		ArbitrarySections []ArbitrarySection
		Days              []Day
		LastTouched       time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "",
			fields: fields{
				ParentVersion: 10,
				Header: Header{
					Contents: "Hi!",
				},
				ArbitrarySections: []ArbitrarySection{
					{
						Contents: `- [x] Done`,
						token:    "# plan.project/plan",
					},
				},
				Days: []Day{
					{
						Contents: `- [x] Done`,
						Date:     time.Date(2022, time.August, 17, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			want: "# plan.header/10\n\nHi!\n\n# plan.project/plan\n\n- [x] Done\n\n# plan.day/2022-08-17 🌱\n\n- [x] Done\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlanFile{
				ParentVersion:     tt.fields.ParentVersion,
				Header:            tt.fields.Header,
				ArbitrarySections: tt.fields.ArbitrarySections,
				Days:              tt.fields.Days,
				LastTouched:       tt.fields.LastTouched,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("PlanFile.String() = %v, want %v; diff = %v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}
