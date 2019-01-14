package core

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ConsenSys/hellhound-vm/hh"
)

func Test_stack_Data(t *testing.T) {
	type fields struct {
		data []*big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   []*big.Int
	}{
		{
			name: "nominal case",
			fields: fields{
				data: []*big.Int{big.NewInt(0), big.NewInt(1),},
			},
			want: []*big.Int{big.NewInt(0), big.NewInt(1),},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := stack{
				data: tt.fields.data,
			}
			if got := stack.Data(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stack.Data() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_Push(t *testing.T) {
	type fields struct {
		data []*big.Int
	}
	type args struct {
		d *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*big.Int
	}{
		{
			name:   "nominal case",
			fields: fields{data: []*big.Int{big.NewInt(0), big.NewInt(1),}},
			args:   args{big.NewInt(2)},
			want:   []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &stack{
				data: tt.fields.data,
			}
			stack.Push(tt.args.d)
			if got := stack.Data(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stack.Push() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_stack_PushN(t *testing.T) {
	type fields struct {
		data []*big.Int
	}
	type args struct {
		ds []*big.Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*big.Int
	}{
		{
			name:   "nominal case",
			fields: fields{data: []*big.Int{big.NewInt(0), big.NewInt(1),}},
			args:   args{[]*big.Int{big.NewInt(2), big.NewInt(3)}},
			want:   []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2), big.NewInt(3)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &stack{
				data: tt.fields.data,
			}
			stack.PushN(tt.args.ds...)
			if got := stack.Data(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stack.PushN() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_stack_Pop(t *testing.T) {
	type fields struct {
		data []*big.Int
	}
	tests := []struct {
		name    string
		fields  fields
		wantRet *big.Int
	}{
		{
			name:    "nominal case",
			fields:  fields{data: []*big.Int{big.NewInt(1), big.NewInt(2)}},
			wantRet: big.NewInt(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &stack{
				data: tt.fields.data,
			}
			if gotRet := stack.Pop(); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("stack.Pop() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func Test_stack_Len(t *testing.T) {
	type fields struct {
		data []*big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "nominal case",
			fields: fields{data: []*big.Int{big.NewInt(1), big.NewInt(2)}},
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := stack{
				data: tt.fields.data,
			}
			if got := stack.Len(); got != tt.want {
				t.Errorf("stack.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_Peek(t *testing.T) {
	type fields struct {
		data []*big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   *big.Int
	}{
		{
			name:   "nominal case",
			fields: fields{data: []*big.Int{big.NewInt(1), big.NewInt(2)}},
			want:   big.NewInt(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := stack{
				data: tt.fields.data,
			}
			if got := stack.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stack.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_Get(t *testing.T) {
	type fields struct {
		data []*big.Int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *big.Int
	}{
		{
			name:   "nominal case",
			fields: fields{data: []*big.Int{big.NewInt(1), big.NewInt(2)}},
			args:   args{0},
			want:   big.NewInt(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &stack{
				data: tt.fields.data,
			}
			if got := stack.Get(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stack.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_Swap(t *testing.T) {
	type fields struct {
		data []*big.Int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*big.Int
	}{
		{
			name:   "nominal case",
			fields: fields{data: []*big.Int{big.NewInt(1), big.NewInt(2)}},
			args:   args{2},
			want:   []*big.Int{big.NewInt(2), big.NewInt(1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &stack{
				data: tt.fields.data,
			}
			stack.Swap(tt.args.n)
			if got := stack.Data(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stack.Swap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStackCfg(t *testing.T) {
	type args struct {
		setters []hh.StackOption
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nominal case",
			args: args{setters: []hh.StackOption{hh.StackSize(8)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StackCfg(tt.args.setters...);
			if got == nil {
				t.Errorf("StackCfg() should not be nil")
			}
			got(&tanden{})
		})
	}
}

func TestNewStack(t *testing.T) {
	type args struct {
		setters []hh.StackOption
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nominal case",
			args: args{setters: []hh.StackOption{hh.StackSize(8)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if nil == NewStack(tt.args.setters...) {
				t.Errorf("NewStack() should not be nil")
			}
		})
	}
}

func TestStackFromOptions(t *testing.T) {
	type args struct {
		opts hh.StackOptions
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nominal case",
			args: args{opts: defaultStackOptions},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if nil == StackFromOptions(tt.args.opts) {
				t.Errorf("StackFromOptions() should not be nil")
			}
		})
	}
}
