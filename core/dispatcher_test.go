package core

import (
	"github.com/ConsenSys/hellhound-vm/dispatchers"
	"github.com/ConsenSys/hellhound-vm/hh"
	"testing"
)

var (
	d = NewDispatcher(
		dispatchers.Keystore(),
		dispatchers.RegisterSet(),
	)
)

func TestDispatcher_Dispatch(t *testing.T) {

	var tests = []struct {
		name    string
		opcode  hh.Opcode
		wantErr bool
	}{
		{"0xFF unknown opcode", hh.Opcode(0xFF), true},
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
