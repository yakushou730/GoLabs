package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				A: []int{3, 8, 9, 7, 6},
				K: 3,
			},
			want: []int{9, 7, 6, 3, 8},
		},
		{
			name: "example 2",
			args: args{
				A: []int{0, 0, 0},
				K: 1,
			},
			want: []int{0, 0, 0},
		},
		{
			name: "example 3",
			args: args{
				A: []int{1, 2, 3, 4},
				K: 4,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "example 4",
			args: args{
				A: []int{},
				K: 4,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
