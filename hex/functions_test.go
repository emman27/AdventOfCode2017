package hex

import (
	"testing"
)

func TestPartA(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{dir: "ne,ne,ne"}, want: 3},
		{args: args{dir: "ne,ne,sw,sw"}, want: 0},
		{args: args{dir: "ne,ne,s,s"}, want: 2},
		{args: args{dir: "se,sw,se,sw,sw"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartA(tt.args.dir); got != tt.want {
				t.Errorf("PartA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartB(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{dir: "ne,ne,ne"}, want: 3},
		{args: args{dir: "ne,ne,sw,sw"}, want: 2},
		{args: args{dir: "ne,ne,s,s"}, want: 2},
		{args: args{dir: "se,sw"}, want: 1},
		{args: args{dir: "se,sw,se"}, want: 2},
		{args: args{dir: "se,sw,se,sw"}, want: 2},
		{args: args{dir: "se,sw,se,sw,sw"}, want: 3},
		{args: args{dir: ""}, want: 0},
		{args: args{dir: "s,n"}, want: 1},
		{args: args{dir: "ne,s,sw,nw,n,ne,se"}, want: 1},
		{args: args{dir: "ne,ne,s,s,sw,sw,nw,nw,n,n,ne,ne,se,se"}, want: 2},
		{args: args{dir: "ne,ne,s,s,s"}, want: 3},
		{args: args{dir: "sw"}, want: 1},
		{args: args{dir: "nw"}, want: 1},
		{args: args{dir: "nw,se"}, want: 1},
		{args: args{dir: "se,nw"}, want: 1},
		{args: args{dir: "sw,ne"}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartB(tt.args.dir); got != tt.want {
				t.Errorf("PartB() = %v, want %v", got, tt.want)
			}
		})
	}
}
