package problems

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy path 1",
			args: args{s: "()"},
			want: true,
		},
		{
			name: "happy path 2",
			args: args{s: "{[()]}"},
			want: true,
		},
		{
			name: "happy path 3",
			args: args{s: "()[]{}"},
			want: true,
		},
		{
			name: "happy path 4",
			args: args{s: "(]"},
			want: false,
		},
		{
			name: "happy path 5",
			args: args{s: "{}{}{}{}{}{}{}{}{}{}[][][][][][][][][]()()()()()()()()(){}{}{}{}{}{}{}()()()()()()()[][][]()[]"},
			want: true,
		},
		{
			name: "happy path 6",
			args: args{s: "{}{})))"},
			want: false,
		},
		{
			name: "happy path 7",
			args: args{s: "(((())))(((())))(("},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid1(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}