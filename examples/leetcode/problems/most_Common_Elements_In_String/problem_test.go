package most_Common_Elements_In_String

import (
	"reflect"
	"testing"
)

func Test_mostCommonElementsInAString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "happy path 1",
			args: args{str: "ioioioiopp"},
			want: map[string]int{
				"i": 4,
				"o": 4,
			},
		},
		{
			name: "happy path 1",
			args: args{str: "dddpppmmmccccccccgggggggg"},
			want: map[string]int{
				"c": 8,
				"g": 8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostCommonElementsInAString(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mostCommonElementsInAString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "happy path 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "happy path 2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			want: []int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate1(tt.args.nums, tt.args.k)
		})
	}
}
