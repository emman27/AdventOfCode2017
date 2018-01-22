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
