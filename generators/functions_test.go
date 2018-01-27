package generators

import "testing"

func Test_next(t *testing.T) {
	type args struct {
		g generator
		i int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{args: args{g: factorA, i: 1092455}, want: 1181022009},
		{args: args{g: factorA, i: 1181022009}, want: 245556042},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := next(tt.args.g, tt.args.i); got != tt.want {
				t.Errorf("next() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		{args: args{a: 65, b: 8921}, want: 588},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartA(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("PartA() = %v, want %v", got, tt.want)
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
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toBinary(tt.args.i); got != tt.want {
				t.Errorf("toBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
