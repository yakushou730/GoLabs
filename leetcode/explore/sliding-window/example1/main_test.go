package main

import "testing"

func Test_findLength(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{3, 1, 2, 7, 4, 2, 1, 1, 5},
				k:    8,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLength(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
