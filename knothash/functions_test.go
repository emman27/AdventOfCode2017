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
