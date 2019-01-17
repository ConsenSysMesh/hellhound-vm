package core

import (
	"errors"
	"github.com/ConsenSys/hellhound-vm/mocks"
	"github.com/golang/mock/gomock"
	"math/big"
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

func Test_tanden_Burn_Nominal_Case(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dispatcherMock := mocks.NewMockDispatcher(ctrl)
	dispatcherMock.
		EXPECT().
		Dispatch(gomock.Eq(hh.LOADKEY)).
		Return(func(_ hh.Tanden, _ hh.Ki) error {
			return nil
		}, nil)

	tanden := tanden{
		dispatcher: dispatcherMock,
	}
	ki := NewKi(
		hh.Kokyu{
			byte(hh.LOADKEY),
			byte(hh.STOP),
		},
	)
	err := tanden.Burn(ki)
	if err != nil {
		t.Error("Burn should not fail in normal conditions")
	}
}

func Test_tanden_Burn_Dispatch_Err(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dispatcherMock := mocks.NewMockDispatcher(ctrl)
	dispatcherMock.
		EXPECT().
		Dispatch(gomock.Eq(hh.LOADKEY)).
		Return(nil, errors.New("dispatch error"))

	tanden := tanden{
		dispatcher: dispatcherMock,
	}
	ki := NewKi(
		hh.Kokyu{
			byte(hh.LOADKEY),
		},
	)
	err := tanden.Burn(ki)
	if err == nil {
		t.Error("Burn should return error when dispatcher returns error")
	}
}

func Test_tanden_Burn_Wave_Err(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dispatcherMock := mocks.NewMockDispatcher(ctrl)
	dispatcherMock.
		EXPECT().
		Dispatch(gomock.Eq(hh.LOADKEY)).
		Return(func(_ hh.Tanden, _ hh.Ki) error {
			return errors.New("wave error")
		}, nil)

	tanden := tanden{
		dispatcher: dispatcherMock,
	}
	ki := NewKi(
		hh.Kokyu{
			byte(hh.LOADKEY),
		},
	)
	err := tanden.Burn(ki)
	if err == nil {
		t.Error("Burn should return error when wave execution returns error")
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
		{
			name:   "nominal case",
			fields: fields{stack: NewStack()},
			want:   NewStack(),
		},
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
		{
			name:   "nominal case",
			fields: fields{sp: 123},
			want:   123,
		},
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
	v := tanden{}.Version()
	if v != version {
		t.Errorf("invalid version : %s", v)
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
		{
			name:   "nominal case",
			fields: fields{keystore: NewKeystore()},
			want:   NewKeystore(),
		},
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
		{
			name:   "nominal case",
			fields: fields{registerSet: NewRegisterSet()},
			want:   NewRegisterSet(),
		},
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
		{
			name:   "nominal case",
			fields: fields{keystore: NewKeystore(), registerSet: NewRegisterSet()},
			want:   []hh.HankoInputProvider{NewKeystore(), NewRegisterSet()},
		},
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
		{
			name: "nominal case",
			fields: fields{
				keystore:   &keystore{
					keys: []*hh.Key{hh.NewKey(0x00,0x00, []byte{0x00})},
				},
				registerSet: &registerSet{
					values: [][]byte{{0x00}, {0x01}},
				},
				stack:       &stack{
					data: []*big.Int{big.NewInt(1)},
				},
			},
		},
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
