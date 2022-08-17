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
			name: "fialing",
			args: args{
				ctx:      context.TODO(),
				planFile: "# plan.header/0\n\nHi! This is my == plan file. 😀✍ == \n\nTwitter: @_aakarim\nGithub: @aakarim\n\n---\n\n# plan.project/plan \n\n\n- [x] add a daemon script\n- [ ] bug - if a file was reverted to a previous version (saved locally, reverted to previous version, saved), it will overwrite. There should be a way to detect this and action the revert, since that was the user's intention.\n- [x] add Saga magic comment/section \n- [ ] the onboarding experience should end with the user runnin the `plan fresh` command. `plan init` will add the header section with the default characters, but the user will need to type in `plan fresh` to get the section for the day.\n\n\n---\n\n# plan.day/2022-08-17\n\n- [ ] new plan should have some signal emojis so it's easy to scan and see which section you should jump to - also should be jumpable easily \n- [ ] set up plan as a doubly linked list\n\n# plan.day/2022-08-16\n\n- [ ] add some random tasks to fresh \n- [ ] implement `plan sync` by first downloading the latest version\n\t- [ ] sync conflicts should be resolved by parsing the file and going section by section, rather than diffing the whole file, because sections may have been added and the information may be confusing to resolve. The diffs _sbould_ be stored in the .plan file, though because it won't be so confusing that it overcomes the friction of having multiple files to represent conflicts.\n- [x] add components for email filtering\n- [x] email ted\n- [x] chase rebekka\n\n# plan.day/2022-08-15\n\n- [x] add a better header\n- [ ] install to $PATH\n- [x] fix $PATH on pland\n\n# plan.day/2022-08-14\n\n- [x] Read Part II of PLG Onboarding\n- [x] Study Org mode\n- [x] Plan `plan.fresh`\n- [x] write parser for plans\n- [x] write stringifier for plans\n\n",
			},
			want:    &PlanFile{},
			wantErr: false,
		},
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
