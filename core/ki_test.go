package core

import (
	"reflect"
	"testing"

	"github.com/ConsenSys/hellhound-vm/hh"
)

func TestNewKi(t *testing.T) {
	type args struct {
		kokyu hh.Kokyu
	}
	tests := []struct {
		name string
		args args
		want hh.Ki
	}{
		{
			name: "nominal case",
			args: args{kokyu: []byte{0x00}},
			want: &ki{kokyu: []byte{0x00}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKi(tt.args.kokyu); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ki_Kokyu(t *testing.T) {
	type fields struct {
		kokyu hh.Kokyu
		pc    int
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name:   "nominal case",
			fields: fields{kokyu: []byte{0x00}},
			want:   []byte{0x00},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ki := ki{
				kokyu: tt.fields.kokyu,
				pc:    tt.fields.pc,
			}
			if got := ki.Kokyu(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ki.Kokyu() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ki_PC(t *testing.T) {
	type fields struct {
		kokyu hh.Kokyu
		pc    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "nominal case",
			fields: fields{pc: 99},
			want:   99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ki := ki{
				kokyu: tt.fields.kokyu,
				pc:    tt.fields.pc,
			}
			if got := ki.PC(); got != tt.want {
				t.Errorf("ki.PC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ki_Burnable(t *testing.T) {
	type fields struct {
		kokyu hh.Kokyu
		pc    int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "nominal case",
			fields: fields{
				kokyu: []byte{byte(hh.LOADKEY)},
				pc:    0,
			},
			want: true,
		},
		{
			name: "not burnable",
			fields: fields{
				kokyu: []byte{byte(hh.LOADKEY), byte(hh.STOP)},
				pc:    1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ki := ki{
				kokyu: tt.fields.kokyu,
				pc:    tt.fields.pc,
			}
			if got := ki.Burnable(); got != tt.want {
				t.Errorf("ki.Burnable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ki_Drained(t *testing.T) {
	type fields struct {
		kokyu hh.Kokyu
		pc    int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "nominal case",
			fields: fields{
				kokyu: []byte{byte(hh.LOADKEY), byte(hh.STOP)},
				pc:    0,
			},
			want: false,
		},
		{
			name: "drained",
			fields: fields{
				kokyu: []byte{byte(hh.LOADKEY)},
				pc:    2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ki := ki{
				kokyu: tt.fields.kokyu,
				pc:    tt.fields.pc,
			}
			if got := ki.Drained(); got != tt.want {
				t.Errorf("ki.Drained() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ki_GetAndMovePCForward(t *testing.T) {
	type fields struct {
		kokyu hh.Kokyu
		pc    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ki := &ki{
				kokyu: tt.fields.kokyu,
				pc:    tt.fields.pc,
			}
			if got := ki.GetAndMovePCForward(); got != tt.want {
				t.Errorf("ki.GetAndMovePCForward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ki_GetAndMovePCForwardN(t *testing.T) {
	type fields struct {
		kokyu hh.Kokyu
		pc    int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ki := &ki{
				kokyu: tt.fields.kokyu,
				pc:    tt.fields.pc,
			}
			if got := ki.GetAndMovePCForwardN(tt.args.n); got != tt.want {
				t.Errorf("ki.GetAndMovePCForwardN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ki_MovePCForward(t *testing.T) {
	type fields struct {
		kokyu hh.Kokyu
		pc    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ki := &ki{
				kokyu: tt.fields.kokyu,
				pc:    tt.fields.pc,
			}
			if got := ki.MovePCForward(); got != tt.want {
				t.Errorf("ki.MovePCForward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ki_MovePCForwardN(t *testing.T) {
	type fields struct {
		kokyu hh.Kokyu
		pc    int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ki := &ki{
				kokyu: tt.fields.kokyu,
				pc:    tt.fields.pc,
			}
			if got := ki.MovePCForwardN(tt.args.n); got != tt.want {
				t.Errorf("ki.MovePCForwardN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ki_ProvideHankoInput(t *testing.T) {
	type fields struct {
		kokyu hh.Kokyu
		pc    int
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ki := ki{
				kokyu: tt.fields.kokyu,
				pc:    tt.fields.pc,
			}
			if got := ki.ProvideHankoInput(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ki.ProvideHankoInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
