package problems

import "testing"

func Test_winnerOfGame(t *testing.T) {
	type args struct {
		colors string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				colors: "AAABABB",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				colors: "AA",
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				colors: "ABBBBBBBAAA",
			},
			want: false,
		},
		{
			name: "Example 4",
			args: args{
				colors: "ABBBAAABBBAAAB",
			},
			want: false,
		},
		{
			name: "Example 5",
			args: args{
				colors: "A",
			},
			want: false,
		},
		{
			name: "Example 6",
			args: args{
				colors: "AABB",
			},
			want: false,
		},
		{
			name: "Example 7",
			args: args{
				colors: "AAAABBBB",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winnerOfGame(tt.args.colors); got != tt.want {
				t.Errorf("winnerOfGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
