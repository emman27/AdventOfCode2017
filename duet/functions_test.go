package duet

import (
	"testing"
)

func TestPartA(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{filename: "./sample_input.txt", want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartA(tt.filename); got != tt.want {
				t.Errorf("PartA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartB(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{filename: "sample_input_b.txt"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartB(tt.args.filename); got != tt.want {
				t.Errorf("PartB() = %v, want %v", got, tt.want)
			}
		})
	}
}
