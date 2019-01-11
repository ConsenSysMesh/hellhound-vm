package core

import (
	"reflect"
	"testing"

	"github.com/ConsenSys/hellhound-vm/hh"
)

func TestKeystoreCfg(t *testing.T) {
	type args struct {
		setters []hh.KeystoreOption
	}
	tests := []struct {
		name string
		args args
		want TandenConfigurer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeystoreCfg(tt.args.setters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeystoreCfg() = %v, want %v", got, tt.want)
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
		want hh.Keystore
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKeystore(tt.args.setters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKeystore() = %v, want %v", got, tt.want)
			}
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeystoreFromOptions(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeystoreFromOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_keystore_Store(t *testing.T) {
	type fields struct {
		keys []*hh.Key
	}
	type args struct {
		slot int
		key  *hh.Key
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
			ks := keystore{
				keys: tt.fields.keys,
			}
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
		want    *hh.Key
		wantErr bool
	}{
		// TODO: Add test cases.
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
			if !reflect.DeepEqual(got, tt.want) {
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
		// TODO: Add test cases.
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

func Test_keystore_validSlot(t *testing.T) {
	type fields struct {
		keys []*hh.Key
	}
	type args struct {
		slot int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ks := keystore{
				keys: tt.fields.keys,
			}
			if got := ks.validSlot(tt.args.slot); got != tt.want {
				t.Errorf("keystore.validSlot() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ks := keystore{
				keys: tt.fields.keys,
			}
			if got := ks.ProvideHankoInput(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keystore.ProvideHankoInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
