package main

import "testing"

func Test_subarraySum(t *testing.T) {
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
				nums: []int{1, 1, 1},
				k:    2,
			},
			want: 2,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{1, 2, 3},
				k:    3,
			},
			want: 2,
		},
		{
			name: "example 3",
			args: args{
				nums: []int{1},
				k:    0,
			},
			want: 0,
		},
		{
			name: "example 4",
			args: args{
				nums: []int{-1, -1, 1},
				k:    0,
			},
			want: 1,
		},
		{
			name: "example 5",
			args: args{
				nums: []int{1, 2, 1, 2, 1},
				k:    3,
			},
			want: 4,
		},
		{
			name: "example 6",
			args: args{
				nums: []int{1, -1, 0},
				k:    0,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
