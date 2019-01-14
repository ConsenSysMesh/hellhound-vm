package core

import (
	"github.com/ConsenSys/hellhound-vm/dispatchers"
	"github.com/ConsenSys/hellhound-vm/hh"
	"reflect"
	"testing"
)

func Test_dispatcher_Dispatch(t *testing.T) {
	d := NewDispatcher(
		dispatchers.Keystore(),
	)
	var tests = []struct {
		name    string
		opcode  hh.OpCode
		wantErr bool
	}{
		{"nominal case", hh.OpCode(hh.LOADKEY), false},
		{"0xFF unknown opcode", hh.OpCode(0xFF), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := d.Dispatch(tt.opcode)
			if !tt.wantErr {

			}
			if tt.wantErr && err == nil {
				t.Errorf("Dispatch(%X) should fail", tt.opcode)
			}
		})
	}

}

func TestDispatcher(t *testing.T) {
	tests := []struct {
		name string
		want TandenConfigurer
	}{
		{
			name: "nominal case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Dispatcher()
			if d == nil{
				t.Error("dispatcher should not be nil")
			}
		})
	}
}

func TestDispatcherCfg(t *testing.T) {
	type args struct {
		subDispatchers [][]hh.Operation
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nominal case",
			args: args{subDispatchers: [][]hh.Operation{dispatchers.Arithmetic()}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configurer := DispatcherCfg(tt.args.subDispatchers...)
			if configurer == nil {
				t.Error("configurer should not be nil")
			}
			configurer(&tanden{})
		})
	}
}

func TestNewDispatcher(t *testing.T) {
	type args struct {
		subDispatchers [][]hh.Operation
	}
	tests := []struct {
		name string
		args args
		want hh.Dispatcher
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDispatcher(tt.args.subDispatchers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDispatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
