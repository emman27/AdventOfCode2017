package swarm

import (
	"reflect"
	"testing"
)

func Test_parseCommand(t *testing.T) {
	type args struct {
		s string
		i int
	}
	tests := []struct {
		name string
		args args
		want Particle
	}{
		{
			args: args{s: "p=<-3162,-1301,225>, v=<92,-52,59>, a=<18,19,-10>", i: 99},
			want: Particle{
				Position:     Geometry{-3162, -1301, 225},
				Velocity:     Geometry{92, -52, 59},
				Acceleration: Geometry{18, 19, -10},
				ID:           99,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseCommand(tt.args.s, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartA(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{filename: "./sample_input.txt"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartA(tt.args.filename); got != tt.want {
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
		{args: args{filename: "./sample_input_b.txt"}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartB(tt.args.filename); got != tt.want {
				t.Errorf("PartB() = %v, want %v", got, tt.want)
			}
		})
	}
}
