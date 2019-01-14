package core

import (
	"testing"

	"github.com/ConsenSys/hellhound-vm/hh"
)

func TestKiWaveNotFound_Error(t *testing.T) {
	type fields struct {
		OpCode hh.OpCode
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "nominal case",
			fields: fields{OpCode:hh.LOADKEY},
			want: "instruction not found for opcode : 30",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inf := KiWaveNotFound{
				OpCode: tt.fields.OpCode,
			}
			if got := inf.Error(); got != tt.want {
				t.Errorf("KiWaveNotFound.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
