package main

import "testing"

func Test_waysToSplitArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{nums: []int{10, 4, -8, 7}},
			want: 2,
		},
		{
			name: "example 2",
			args: args{nums: []int{2, 3, 1, 0}},
			want: 2,
		},
		{
			name: "example 3",
			args: args{nums: []int{0, 0}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToSplitArray(tt.args.nums); got != tt.want {
				t.Errorf("waysToSplitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
