package stream

import (
	"testing"
)

func Test_removeGarbage(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Basic", args: args{s: "<>"}, want: ""},
		{name: "Random Chars", args: args{s: "<Random characters>"}, want: ""},
		{name: "Ignore extras", args: args{s: "<<<<>"}, want: ""},
		{name: "Loads of trash", args: args{s: "<{o\"i!a,<{i<a>"}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeGarbage(tt.args.s); got != tt.want {
				t.Errorf("removeGarbage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_score(t *testing.T) {
	type args struct {
		s    string
		base int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Simple", args: args{s: "{}", base: 1}, want: 1},
		{name: "Nest", args: args{s: "{{{}}}", base: 1}, want: 6},
		{name: "Nest Double", args: args{s: "{{},{}}", base: 1}, want: 5},
		{name: "Nest Multi", args: args{s: "{{{},{},{{}}}}", base: 1}, want: 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := score(tt.args.s, tt.args.base); got != tt.want {
				t.Errorf("score() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartA(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Simple", args: args{s: "{}"}, want: 1},
		{name: "Nest", args: args{s: "{{{}}}"}, want: 6},
		{name: "Nest Double", args: args{s: "{{},{}}"}, want: 5},
		{name: "Nest Multi", args: args{s: "{{{},{},{{}}}}"}, want: 16},
		{name: "Garbage", args: args{s: "{<a>,<a>,<a>,<a>}"}, want: 1},
		{name: "Garbage more", args: args{s: "{{<ab>},{<ab>},{<ab>},{<ab>}}"}, want: 9},
		{name: "Garbage with ignores", args: args{s: "{{<!!>},{<!!>},{<!!>},{<!!>}}"}, want: 9},
		{name: "Everything is a problem", args: args{s: "{{<a!>},{<a!>},{<a!>},{<ab>}}"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartA(tt.args.s); got != tt.want {
				t.Errorf("PartA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartB(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{s: "<>"}, want: 0},
		{args: args{s: "<random characters>"}, want: 17},
		{args: args{s: "<{!>}>"}, want: 2},
		{args: args{s: "<<<<>"}, want: 3},
		{args: args{s: "!!"}, want: 0},
		{args: args{s: "!!!>"}, want: 0},
		{args: args{s: "<{o\"i!a,<{i<a>"}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartB(tt.args.s); got != tt.want {
				t.Errorf("PartB() = %v, want %v", got, tt.want)
			}
		})
	}
}
