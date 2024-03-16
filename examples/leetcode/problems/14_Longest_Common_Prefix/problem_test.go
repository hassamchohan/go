package problems

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy path 1",
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "happy path 2",
			args: args{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
		{
			name: "happy path 3",
			args: args{strs: []string{"hassam", "hassam123", "hassam456", "hassam789"}},
			want: "hassam",
		},
		{
			name: "happy path 4",
			args: args{strs: []string{}},
			want: "",
		},
		{
			name: "happy path 5",
			args: args{strs: []string{"dog"}},
			want: "dog",
		},
		{
			name: "happy path 6",
			args: args{strs: []string{"a"}},
			want: "a",
		},
		{
			name: "happy path 7",
			args: args{strs: []string{"cir", "car"}},
			want: "c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
