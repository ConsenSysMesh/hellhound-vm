package core

import (
	"fmt"
	"testing"
)

func TestHhvm_New(t *testing.T) {
	hhvm, err := NewTanden()
	if err != nil{
		t.Errorf("HHVM creation should not fail : %s", err.Error())
	}

	fmt.Println("HHVM version : ", hhvm.Version())
}
