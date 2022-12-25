package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{N: 9},
			want: 2,
		},
		{
			name: "example 2",
			args: args{N: 1041},
			want: 5,
		},
		{
			name: "example 3",
			args: args{N: 32},
			want: 0,
		},
		{
			name: "example 4",
			args: args{N: 1024},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.N); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
