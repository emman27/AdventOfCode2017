package knothash

import (
	"reflect"
	"testing"
)

func Test_stepForward(t *testing.T) {
	type args struct {
		s     []int
		steps int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{s: []int{1, 2, 3}, steps: 1}, want: []int{2, 3, 1}},
		{args: args{s: []int{1, 2, 3}, steps: 3}, want: []int{1, 2, 3}},
		{args: args{s: []int{1, 2, 3}, steps: 7}, want: []int{2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stepForward(tt.args.s, tt.args.steps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stepForward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		s []int
		l int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{s: []int{1, 2, 3}, l: 1}, want: []int{1, 2, 3}},
		{args: args{s: []int{1, 2, 3}, l: 2}, want: []int{2, 1, 3}},
		{args: args{s: []int{1, 2, 3}, l: 3}, want: []int{3, 2, 1}},
		{args: args{s: []int{1, 2, 3}, l: 0}, want: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.s, tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartA(t *testing.T) {
	type args struct {
		cmds []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartA(tt.args.cmds); got != tt.want {
				t.Errorf("PartA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_move(t *testing.T) {
	type args struct {
		cmds []int
		s    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{cmds: []int{3}, s: []int{0, 1, 2, 3, 4}}, want: 2},
		{args: args{cmds: []int{3, 4}, s: []int{0, 1, 2, 3, 4}}, want: 12},
		{args: args{cmds: []int{3, 4, 1}, s: []int{0, 1, 2, 3, 4}}, want: 12},
		{args: args{cmds: []int{3, 4, 1, 5}, s: []int{0, 1, 2, 3, 4}}, want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := move(tt.args.cmds, tt.args.s); got != tt.want {
				t.Errorf("move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_determineSequence(t *testing.T) {
	type args struct {
		cmds string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{cmds: "1,2,3"}, want: []int{49, 44, 50, 44, 51, 17, 31, 73, 47, 23}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determineSequence(tt.args.cmds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("determineSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xor(t *testing.T) {
	type args struct {
		orig []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{orig: []int{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22}}, want: 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xor(tt.args.orig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combineXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToHex(t *testing.T) {
	type args struct {
		res []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{res: []int{64, 7, 255}}, want: "4007ff"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToHex(tt.args.res); got != tt.want {
				t.Errorf("convertToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartB(t *testing.T) {
	type args struct {
		cmds string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{cmds: ""}, want: "a2582a3a0e66e6e86e3812dcb672a272"},
		{args: args{cmds: "AoC 2017"}, want: "33efeb34ea91902bb2f59c9920caa6cd"},
		{args: args{cmds: "1,2,3"}, want: "3efbe78a8d82f29979031a4aa0b16a9d"},
		{args: args{cmds: "1,2,4"}, want: "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartB(tt.args.cmds); got != tt.want {
				t.Errorf("PartB() = %v, want %v", got, tt.want)
			}
		})
	}
}
