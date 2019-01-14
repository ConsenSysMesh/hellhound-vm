package core

import (
	"reflect"
	"testing"

	"github.com/ConsenSys/hellhound-vm/hh"
)

func TestRegisterSetCfg(t *testing.T) {
	type args struct {
		setters []hh.RegisterSetOption
	}

	tandenMock := tanden{}

	tests := []struct {
		name                      string
		args                      args
		wantNil                   bool
		wantRegisterSetSlotNumber int
		tandenMock                tanden
	}{
		{
			name:                      "nominal case",
			args:                      args{setters: []hh.RegisterSetOption{hh.RegisterSetSlotNumber(8)}},
			wantNil:                   false,
			wantRegisterSetSlotNumber: 8,
			tandenMock:                tandenMock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RegisterSetCfg(tt.args.setters...)
			if got == nil && !tt.wantNil {
				t.Error("configurer should not be nil")
			}
			got(&tt.tandenMock)
			if len(tt.tandenMock.registerSet.Values()) != tt.wantRegisterSetSlotNumber {
				t.Error("invalid register set slot number")
			}
		})
	}
}

func TestNewRegisterSet(t *testing.T) {
	type args struct {
		setters []hh.RegisterSetOption
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nominal case",
			args: args{setters: []hh.RegisterSetOption{hh.RegisterSetSlotNumber(8)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewRegisterSet(tt.args.setters...)
		})
	}
}

func TestRegisterSetFromOptions(t *testing.T) {
	type args struct {
		opts hh.RegisterSetOptions
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{opts: defaultRegisterSetOptions},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterSetFromOptions(tt.args.opts)
		})
	}
}

func Test_registerSet_Store(t *testing.T) {
	type opts struct {
		setters []hh.RegisterSetOption
	}
	type args struct {
		slot  int
		entry []byte
	}
	tests := []struct {
		name    string
		opts    opts
		args    args
		wantErr bool
	}{
		{
			name: "nominal case",
			opts: opts{setters: []hh.RegisterSetOption{hh.RegisterSetSlotNumber(8)}},
			args: args{
				slot:  1,
				entry: []byte{0x00},
			},
		},
		{
			name: "invalid slot",
			opts: opts{setters: []hh.RegisterSetOption{hh.RegisterSetSlotNumber(8)}},
			args: args{
				slot:  -1,
				entry: []byte{0x00},
			},
			wantErr: true,
		},
		{
			name: "empty value",
			opts: opts{setters: []hh.RegisterSetOption{hh.RegisterSetSlotNumber(8)}},
			args: args{
				slot:  1,
				entry: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := NewRegisterSet(tt.opts.setters...)
			if err := rs.Store(tt.args.slot, tt.args.entry); (err != nil) != tt.wantErr {
				t.Errorf("registerSet.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_registerSet_Get(t *testing.T) {
	type fields struct {
		values [][]byte
	}
	type args struct {
		slot int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:   "nominal case",
			fields: fields{values: [][]byte{{0x00}}},
			args:   args{slot: 0,},
			want:   []byte{0x00},
		},
		{
			name:   "invalid slot",
			fields: fields{values: [][]byte{{0x00}}},
			args: args{slot: -1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := registerSet{values:tt.fields.values}
			got, err := rs.Get(tt.args.slot)
			if (err != nil) != tt.wantErr {
				t.Errorf("registerSet.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("registerSet.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_registerSet_ProvideHankoInput(t *testing.T) {
	type fields struct {
		values [][]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name:   "nominal case",
			fields: fields{values: [][]byte{{0x00}}},
			want:   []byte{0x00},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := registerSet{
				values: tt.fields.values,
			}
			if got := rs.ProvideHankoInput(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("registerSet.ProvideHankoInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
