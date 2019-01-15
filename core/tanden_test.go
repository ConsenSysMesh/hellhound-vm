package core

import (
	"reflect"
	"testing"

	"github.com/ConsenSys/hellhound-vm/hh"
)

func TestNewTanden(t *testing.T) {
	type args struct {
		configurers []TandenConfigurer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "nominal case",
			args: args{configurers: []TandenConfigurer{
				KeystoreCfg(hh.KeystoreSlotNumber(8)),
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewTanden(tt.args.configurers...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTanden() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_tanden_Burn(t *testing.T) {
	type fields struct {
		keystore    hh.Keystore
		registerSet hh.RegisterSet
		dispatcher  hh.Dispatcher
		stack       hh.Stack
		sp          int
	}
	type args struct {
		ki hh.Ki
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tanden := &tanden{
				keystore:    tt.fields.keystore,
				registerSet: tt.fields.registerSet,
				dispatcher:  tt.fields.dispatcher,
				stack:       tt.fields.stack,
				sp:          tt.fields.sp,
			}
			if err := tanden.Burn(tt.args.ki); (err != nil) != tt.wantErr {
				t.Errorf("tanden.Burn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_tanden_Flows(t *testing.T) {
	type fields struct {
		keystore    hh.Keystore
		registerSet hh.RegisterSet
		dispatcher  hh.Dispatcher
		stack       hh.Stack
		sp          int
	}
	type args struct {
		ki hh.Ki
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    hh.KiWave
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tanden := &tanden{
				keystore:    tt.fields.keystore,
				registerSet: tt.fields.registerSet,
				dispatcher:  tt.fields.dispatcher,
				stack:       tt.fields.stack,
				sp:          tt.fields.sp,
			}
			got, err := tanden.Flows(tt.args.ki)
			if (err != nil) != tt.wantErr {
				t.Errorf("tanden.Flows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tanden.Flows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tanden_reset(t *testing.T) {
	type fields struct {
		keystore    hh.Keystore
		registerSet hh.RegisterSet
		dispatcher  hh.Dispatcher
		stack       hh.Stack
		sp          int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tanden := &tanden{
				keystore:    tt.fields.keystore,
				registerSet: tt.fields.registerSet,
				dispatcher:  tt.fields.dispatcher,
				stack:       tt.fields.stack,
				sp:          tt.fields.sp,
			}
			tanden.reset()
		})
	}
}

func Test_tanden_Stack(t *testing.T) {
	type fields struct {
		keystore    hh.Keystore
		registerSet hh.RegisterSet
		dispatcher  hh.Dispatcher
		stack       hh.Stack
		sp          int
	}
	tests := []struct {
		name   string
		fields fields
		want   hh.Stack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tanden := tanden{
				keystore:    tt.fields.keystore,
				registerSet: tt.fields.registerSet,
				dispatcher:  tt.fields.dispatcher,
				stack:       tt.fields.stack,
				sp:          tt.fields.sp,
			}
			if got := tanden.Stack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tanden.Stack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tanden_SP(t *testing.T) {
	type fields struct {
		keystore    hh.Keystore
		registerSet hh.RegisterSet
		dispatcher  hh.Dispatcher
		stack       hh.Stack
		sp          int
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
			tanden := tanden{
				keystore:    tt.fields.keystore,
				registerSet: tt.fields.registerSet,
				dispatcher:  tt.fields.dispatcher,
				stack:       tt.fields.stack,
				sp:          tt.fields.sp,
			}
			if got := tanden.SP(); got != tt.want {
				t.Errorf("tanden.SP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tanden_Version(t *testing.T) {
	type fields struct {
		keystore    hh.Keystore
		registerSet hh.RegisterSet
		dispatcher  hh.Dispatcher
		stack       hh.Stack
		sp          int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tanden := tanden{
				keystore:    tt.fields.keystore,
				registerSet: tt.fields.registerSet,
				dispatcher:  tt.fields.dispatcher,
				stack:       tt.fields.stack,
				sp:          tt.fields.sp,
			}
			if got := tanden.Version(); got != tt.want {
				t.Errorf("tanden.Version() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tanden_Keystore(t *testing.T) {
	type fields struct {
		keystore    hh.Keystore
		registerSet hh.RegisterSet
		dispatcher  hh.Dispatcher
		stack       hh.Stack
		sp          int
	}
	tests := []struct {
		name   string
		fields fields
		want   hh.Keystore
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tanden := tanden{
				keystore:    tt.fields.keystore,
				registerSet: tt.fields.registerSet,
				dispatcher:  tt.fields.dispatcher,
				stack:       tt.fields.stack,
				sp:          tt.fields.sp,
			}
			if got := tanden.Keystore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tanden.Keystore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tanden_RegisterSet(t *testing.T) {
	type fields struct {
		keystore    hh.Keystore
		registerSet hh.RegisterSet
		dispatcher  hh.Dispatcher
		stack       hh.Stack
		sp          int
	}
	tests := []struct {
		name   string
		fields fields
		want   hh.RegisterSet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tanden := tanden{
				keystore:    tt.fields.keystore,
				registerSet: tt.fields.registerSet,
				dispatcher:  tt.fields.dispatcher,
				stack:       tt.fields.stack,
				sp:          tt.fields.sp,
			}
			if got := tanden.RegisterSet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tanden.RegisterSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tanden_HankoInputProviders(t *testing.T) {
	type fields struct {
		keystore    hh.Keystore
		registerSet hh.RegisterSet
		dispatcher  hh.Dispatcher
		stack       hh.Stack
		sp          int
	}
	tests := []struct {
		name   string
		fields fields
		want   []hh.HankoInputProvider
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tanden := tanden{
				keystore:    tt.fields.keystore,
				registerSet: tt.fields.registerSet,
				dispatcher:  tt.fields.dispatcher,
				stack:       tt.fields.stack,
				sp:          tt.fields.sp,
			}
			if got := tanden.HankoInputProviders(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tanden.HankoInputProviders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tanden_Dump(t *testing.T) {
	type fields struct {
		keystore    hh.Keystore
		registerSet hh.RegisterSet
		dispatcher  hh.Dispatcher
		stack       hh.Stack
		sp          int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tanden := tanden{
				keystore:    tt.fields.keystore,
				registerSet: tt.fields.registerSet,
				dispatcher:  tt.fields.dispatcher,
				stack:       tt.fields.stack,
				sp:          tt.fields.sp,
			}
			tanden.Dump()
		})
	}
}
