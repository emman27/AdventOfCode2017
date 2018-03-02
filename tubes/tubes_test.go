package tubes

import (
	"reflect"
	"testing"
)

func TestPartA(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{filename: "sample_input.txt"}, want: "ABCDEF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartA(tt.args.filename); got != tt.want {
				t.Errorf("PartA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_packet_move(t *testing.T) {
	type fields struct {
		X         int
		Y         int
		Direction direction
		Trail     []byte
	}
	type args struct {
		d diagram
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   packet
	}{
		{
			fields: fields{X: 7, Y: 8, Direction: RIGHT, Trail: []byte{}},
			args:   args{d: diagram{Points: map[int]map[int]byte{7: map[int]byte{7: '-', 8: '+'}, 6: map[int]byte{8: '|'}}}},
			want:   packet{X: 6, Y: 8, Direction: UP, Trail: []byte{}, Moves: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := packet{
				X:         tt.fields.X,
				Y:         tt.fields.Y,
				Direction: tt.fields.Direction,
				Trail:     tt.fields.Trail,
			}
			if got := p.move(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("packet.move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartB(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{filename: "sample_input.txt"}, want: 38},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartB(tt.args.filename); got != tt.want {
				t.Errorf("PartB() = %v, want %v", got, tt.want)
			}
		})
	}
}
