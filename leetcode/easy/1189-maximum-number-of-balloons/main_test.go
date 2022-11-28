package main

import "testing"

func Test_maxNumberOfBalloons(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{text: "nlaebolko"},
			want: 1,
		},
		{
			name: "example 2",
			args: args{text: "loonbalxballpoon"},
			want: 2,
		},
		{
			name: "example 3",
			args: args{text: "leetcode"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumberOfBalloons(tt.args.text); got != tt.want {
				t.Errorf("maxNumberOfBalloons() = %v, want %v", got, tt.want)
			}
		})
	}
}
