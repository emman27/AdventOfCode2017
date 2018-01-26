package defrag

import (
	"reflect"
	"testing"
)

func Test_convertToBitmask(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{c: '0'}, want: "0000"},
		{args: args{c: '1'}, want: "0001"},
		{args: args{c: 'f'}, want: "1111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBitmask(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToBitmask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartA(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{input: "flqrgnkx"}, want: 8108},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartA(tt.args.input); got != tt.want {
				t.Errorf("PartA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartB(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{input: "flqrgnkx"}, want: 1242},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartB(tt.args.input); got != tt.want {
				t.Errorf("PartB() = %v, want %v", got, tt.want)
			}
		})
	}
}
