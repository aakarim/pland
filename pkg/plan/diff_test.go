package plan

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_prettyDiff(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				str1: "hi",
				str2: "hi",
			},
			want: "",
		},
		{
			args: args{
				str1: "hello\nthis\nis\nmy\nname",
				str2: "hello\nthis\nis\nmy\nface",
			},
			want: ">>>> Local copy\nhello\nthis\nis\nmy\nname\n<<<< Server copy\nhello\nthis\nis\nmy\nface",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prettyDiff(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("prettyDiff() = %v, want %v; diff = %v", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}
