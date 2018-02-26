package utils

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestQueue_Pop(t *testing.T) {
	type fields struct {
		Items []interface{}
		mux   sync.Mutex
	}
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      interface{}
		wantErr   bool
		wantAfter []interface{}
	}{
		{fields: fields{Items: []interface{}{}}, args: args{timeout: 1 * time.Second}, wantErr: true, wantAfter: []interface{}{}},
		{fields: fields{Items: []interface{}{1}}, args: args{timeout: 1 * time.Second}, want: 1, wantAfter: []interface{}{}},
		{fields: fields{Items: []interface{}{1, 2}}, args: args{timeout: 1 * time.Second}, want: 1, wantAfter: []interface{}{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				Items: tt.fields.Items,
				mux:   tt.fields.mux,
			}
			got, err := q.Pop(tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queue.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Pop() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(q.Items, tt.wantAfter) {
				t.Errorf("Queue.Pop() inconsistent end state = %v, want %v", got, tt.wantAfter)
			}
		})
	}
}
