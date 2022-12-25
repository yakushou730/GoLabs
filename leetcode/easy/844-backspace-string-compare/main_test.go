package main

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				s: "ab#c",
				t: "ad#c",
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				s: "ab##",
				t: "c#d#",
			},
			want: true,
		},
		{
			name: "example 3",
			args: args{
				s: "a#c",
				t: "b",
			},
			want: false,
		},
		{
			name: "example 4",
			args: args{
				s: "xywrrmp",
				t: "xywrrmu#p",
			},
			want: true,
		},
		{
			name: "example 5",
			args: args{
				s: "y#fo##f",
				t: "y#f#o##f",
			},
			want: true,
		},
		{
			name: "example 6",
			args: args{
				s: "a##c",
				t: "#a#c",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
