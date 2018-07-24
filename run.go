package main

import (
	"github.com/chickencoder/run/vm"
)

func main() {
	program := []vm.Instruction{
		vm.Const, vm.NewOperand(10.0, "number"),
		vm.Const, vm.NewOperand(20.0, "number"),
		vm.Add,
		vm.Halt,
	}
	runner := vm.NewRunner(program, 5, 0, true)
	runner.Run()
}
