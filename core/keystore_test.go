package core

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"

	"github.com/ConsenSys/hellhound-vm/hh"
)

func TestKeystoreCfg(t *testing.T) {
	type args struct {
		setters []hh.KeystoreOption
	}

	tandenMock := tanden{}

	tests := []struct {
		name string
		args args
		wantNil bool
		wantKeystoreSlotNumber int
		tandenMock tanden
	}{
		{
			name: "nominal case",
			args: args{setters: []hh.KeystoreOption{hh.KeystoreSlotNumber(8)}},
			wantNil:false,
			wantKeystoreSlotNumber: 8,
			tandenMock: tandenMock,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KeystoreCfg(tt.args.setters...)
			if got == nil && !tt.wantNil{
				t.Error("configurer should not be nil")
			}
			got(&tt.tandenMock)
			if len(tt.tandenMock.keystore.Keys()) != tt.wantKeystoreSlotNumber{
				t.Error("invalid keystore slot number")
			}
		})
	}
}

func TestNewKeystore(t *testing.T) {
	type args struct {
		setters []hh.KeystoreOption
	}
	tests := []struct {
		name string
		args args
		wantKeystoreSlotNumber int
	}{
		{
			name: "nominal case",
			args: args{setters: []hh.KeystoreOption{hh.KeystoreSlotNumber(8)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewKeystore(tt.args.setters...)
		})
	}
}

func TestKeystoreFromOptions(t *testing.T) {
	type args struct {
		opts hh.KeystoreOptions
	}
	tests := []struct {
		name string
		args args
		want hh.Keystore
	}{
		{
			name: "nominal case",
			args: args{defaultKeystoreOptions},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			KeystoreFromOptions(tt.args.opts)
		})
	}
}

func Test_keystore_Store(t *testing.T) {
	type opts struct {
		setters []hh.KeystoreOption
	}
	type args struct {
		slot int
		key  *hh.Key
	}
	keyMock := hh.NewKey(0x00, 0x00, bytes.Repeat([]byte{0x00}, 16))

	tests := []struct {
		name    string
		opts  opts
		args    args
		wantErr bool
	}{
		{
			name: "nominal case",
			opts: opts{setters: []hh.KeystoreOption{hh.KeystoreSlotNumber(8)}},
			args:args{
				slot:0,
				key: keyMock,
			},
			wantErr: false,
		},
		{
			name: "nil key",
			opts: opts{setters: []hh.KeystoreOption{hh.KeystoreSlotNumber(8)}},
			args:args{
				slot:0,
				key: nil,
			},
			wantErr: true,
		},
		{
			name: "invalid slot",
			opts: opts{setters: []hh.KeystoreOption{hh.KeystoreSlotNumber(8)}},
			args:args{
				slot:-1,
				key: keyMock,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ks := NewKeystore(tt.opts.setters...)
			if err := ks.Store(tt.args.slot, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("keystore.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_keystore_Get(t *testing.T) {
	type fields struct {
		keys []*hh.Key
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
		ignoreReturn bool
	}{
		{
			name: "nominal case",
			fields: fields{
				[]*hh.Key{
					hh.NewKey(0x00, 0x00, bytes.Repeat([]byte{0x00}, 16)),
				},
			},
			args:args{slot:0},
			want: bytes.Repeat([]byte{0x00}, 16),
			wantErr: false,
		},
		{
			name: "invalid slot",
			fields: fields{
				[]*hh.Key{
					hh.NewKey(0x00, 0x00, bytes.Repeat([]byte{0x00}, 16)),
				},
			},
			args:args{slot:-1},
			ignoreReturn: true,
			wantErr: true,
		},
		{
			name: "empty slot",
			fields: fields{
				[]*hh.Key{
					hh.NewKey(0x00, 0x00, bytes.Repeat([]byte{0x00}, 16)),
					nil,
				},
			},
			args:args{slot:1},
			wantErr: true,
			ignoreReturn:true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ks := keystore{
				keys: tt.fields.keys,
			}
			got, err := ks.Get(tt.args.slot)
			if (err != nil) != tt.wantErr {
				t.Errorf("keystore.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.ignoreReturn && !reflect.DeepEqual(got.Value, tt.want) {
				t.Errorf("keystore.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_keystore_Keys(t *testing.T) {
	type fields struct {
		keys []*hh.Key
	}
	tests := []struct {
		name   string
		fields fields
		want   []*hh.Key
	}{
		{
			name: "nominal case",
			fields: fields{
				[]*hh.Key{
					hh.NewKey(0x00, 0x00, bytes.Repeat([]byte{0x00}, 16)),
				},
			},
			want: []*hh.Key{
				hh.NewKey(0x00, 0x00, bytes.Repeat([]byte{0x00}, 16)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ks := keystore{
				keys: tt.fields.keys,
			}
			if got := ks.Keys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keystore.Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_keystore_ProvideHankoInput(t *testing.T) {
	type fields struct {
		keys []*hh.Key
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "nominal case",
			fields: fields{
				[]*hh.Key{
					hh.NewKey(0x00, 0x00, bytes.Repeat([]byte{0x00}, 16)),
				},
			},
			want: bytes.Repeat([]byte{0x00}, 16),

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ks := keystore{
				keys: tt.fields.keys,
			}
			if got := ks.ProvideHankoInput(); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(hex.EncodeToString(got))
				t.Errorf("keystore.ProvideHankoInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
