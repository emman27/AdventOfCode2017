package spinlock

import (
	"reflect"
	"testing"
)

func TestPartA(t *testing.T) {
	type args struct {
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{steps: 3}, want: 638},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartA(tt.args.steps); got != tt.want {
				t.Errorf("PartA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_performSpin(t *testing.T) {
	type args struct {
		buffer  []int
		currPos int
		steps   int
		val     int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 int
	}{
		{args: args{buffer: []int{0}, currPos: 0, steps: 3, val: 1}, want: []int{0, 1}, want1: 1},
		{args: args{buffer: []int{0, 1}, currPos: 1, steps: 3, val: 2}, want: []int{0, 2, 1}, want1: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := performSpin(tt.args.buffer, tt.args.currPos, tt.args.steps, tt.args.val)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("performSpin() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("performSpin() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
