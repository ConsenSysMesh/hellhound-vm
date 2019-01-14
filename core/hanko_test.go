package core

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/ConsenSys/hellhound-vm/hh"
)

func TestHankoSensei(t *testing.T) {
	HankoSensei()
}

func Test_sha3HankoSensei_Hanko(t *testing.T) {
	type args struct {
		input func() []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "nominal case",
			args: args{
				input: func() []byte {
					return []byte{0x00}
				},
			},
			want: fromHex("bc36789e7a1e281436464229828f817d6612f7b477d66591ff96a9e064bcc98a"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sha3HankoSensei{}
			if got := s.Hanko(hankoInputProviderFunc{tt.args.input}); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sha3HankoSensei.Hanko() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sha3HankoSensei_HankoN(t *testing.T) {
	type args struct {
		hankoInputProviders []hh.HankoInputProvider
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "nominal case",
			args:args{
				hankoInputProviders: []hh.HankoInputProvider{
					hankoInputProviderMock(func() []byte {
						return []byte{0x00}
					}),
					hankoInputProviderMock(func() []byte {
						return []byte{0x01}
					}),
				},
			},
			want: fromHex("49d03a195e239b52779866b33024210fc7dc66e9c2998975c0aa45c1702549d5"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sha3HankoSensei{}
			if got := s.HankoN(tt.args.hankoInputProviders...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sha3HankoSensei.HankoN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hanko(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hanko(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hanko() = %v, want %v", got, tt.want)
			}
		})
	}
}

type hankoInputProviderFunc struct {
	f func() []byte
}

func hankoInputProviderMock(f func() []byte) hh.HankoInputProvider{
	return hankoInputProviderFunc{f}
}

func (h hankoInputProviderFunc) ProvideHankoInput() []byte {
	return h.f()
}

func fromHex(in string) []byte{
	out, _ := hex.DecodeString(in)
	return out
}