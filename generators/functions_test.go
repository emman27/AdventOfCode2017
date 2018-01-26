package generators

import (
	"testing"
)

func TestPartA(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{a: int64(65), b: int64(8921)}, want: 588},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartA(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("PartA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genNext(t *testing.T) {
	type args struct {
		factor int64
		prev   int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{args: args{factor: int64(16807), prev: int64(65)}, want: int64(1092455)},
		{args: args{factor: int64(16807), prev: int64(1181022009)}, want: int64(245556042)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genNext(tt.args.factor, tt.args.prev); got != tt.want {
				t.Errorf("genNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toBinary(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{i: int64(1092455)}, want: "00000000000100001010101101100111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toBinary(tt.args.i); got != tt.want {
				t.Errorf("toBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
