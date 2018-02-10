package permutation

import (
	"testing"
)

func Test_spin(t *testing.T) {
	type args struct {
		s string
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "abcde", i: 1}, want: "eabcd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spin(tt.args.s, tt.args.i); got != tt.want {
				t.Errorf("spin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_exchange(t *testing.T) {
	type args struct {
		s string
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "eabcd", a: 3, b: 4}, want: "eabdc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exchange(tt.args.s, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("exchange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partner(t *testing.T) {
	type args struct {
		s string
		a rune
		b rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "eabdc", a: 'e', b: 'b'}, want: "baedc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partner(tt.args.s, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("partner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readFile(tt.args.filename); got != tt.want {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_effectivePartA(t *testing.T) {
	got := effectivePartA("abcdefghijklmnop")
	want := PartA("./input.txt", "abcdefghijklmnop")
	if want != got {
		t.Errorf("effectivePartA() = %v, want %v", got, want)
	}

}
