package main

import (
	"github.com/chickencoder/run/vm"
)

func main() {
	program := []vm.Instruction{
		vm.Add,                                 // sum: ADD
		vm.Return,                              // ret
		vm.Const, vm.NewOperand(10., "number"), // const 10
		vm.Const, vm.NewOperand(20., "number"), // const 20
		vm.Call, vm.NewOperand(0., "number"), vm.NewOperand(2., "number"), // call sum, 2
		vm.Halt,
	}
	runner := vm.NewRunner(program, 5, 2, true)
	runner.Run()
}
