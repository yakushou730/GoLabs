package main

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "example 1",
			args: args{s: []byte("hello")},
			want: []byte("olleh"),
		},
		{
			name: "example 2",
			args: args{s: []byte("Hannah")},
			want: []byte("hannaH"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			if string(tt.args.s) != string(tt.want) {
				t.Errorf("after reverseString(), s = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}
