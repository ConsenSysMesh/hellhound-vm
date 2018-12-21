package dispatchers

import "github.com/ConsenSys/hellhound-vm/hh"

func Arithmetic() []hh.Operation {
	return []hh.Operation{
		hh.NewKiWave(hh.ADD, add),
		hh.NewKiWave(hh.SUB, sub),
		hh.NewKiWave(hh.MUL, mul),
		hh.NewKiWave(hh.DIV, div),
	}
}

func add(vm hh.VM, _ *hh.Ki) error {
	x, y := vm.Stack().Pop(), vm.Stack().Peek()
	y.Add(x, y)
	return nil
}

func sub(vm hh.VM, _ *hh.Ki) error {
	x, y := vm.Stack().Pop(), vm.Stack().Peek()
	y.Sub(x, y)
	return nil
}

func mul(vm hh.VM, _ *hh.Ki) error {
	x, y := vm.Stack().Pop(), vm.Stack().Peek()
	y.Mul(x, y)
	return nil
}

func div(vm hh.VM, _ *hh.Ki) error {
	x, y := vm.Stack().Pop(), vm.Stack().Peek()
	y.Div(x, y)
	return nil
}
