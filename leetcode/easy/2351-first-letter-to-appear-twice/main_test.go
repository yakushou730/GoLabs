package main

import "testing"

func Test_repeatedCharacter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "example 1",
			args: args{s: "abccbaacz"},
			want: 'c',
		},
		{
			name: "example 2",
			args: args{s: "abcdd"},
			want: 'd',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedCharacter(tt.args.s); got != tt.want {
				t.Errorf("repeatedCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
