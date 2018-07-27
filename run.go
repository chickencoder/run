package main

import (
	"fmt"

	"github.com/chickencoder/run/vm"
)

func main() {
	// trace := flag.Bool("trace", false, "Trace the program execution")
	// flag.Parse()

	// program := []vm.Instruction{
	// 	vm.Add,                                  // sum: ADD
	// 	vm.Return,                               // ret
	// 	vm.Const, vm.NewOperand(10.2, "number"), // const 10.2
	// 	vm.Const, vm.NewOperand(20.89, "number"), // const 20.89
	// 	vm.Call, vm.NewOperand(0., "number"), vm.NewOperand(2.0, "number"), // call sum, 2
	// 	vm.Pop,
	// 	vm.Halt,
	// }

	// tenValue := vm.Value{
	// 	Kind:    vm.NumberValue,
	// 	Content: 10.0,
	// }

	// twoValue := vm.Value{
	// 	Kind:    vm.NumberValue,
	// 	Content: 2.0,
	// }

	// zeroValue := vm.Value{
	// 	Kind:    vm.NumberValue,
	// 	Content: 0.0,
	// }

	// oneValue := vm.Value{
	// 	Kind:    vm.NumberValue,
	// 	Content: 1.0,
	// }

	// functionLocation := vm.Value{
	// 	Kind:    vm.NumberValue,
	// 	Content: 0.0,
	// }

	// stringValue := vm.Value{
	// 	Kind:    vm.StringValue,
	// 	Content: "Hello World",
	// }

	// sum:
	//		load x
	//		load y
	//		add
	//		ret
	//
	// 		const 10
	//		const 2
	//		store x
	//		store y
	//		call sum, 2
	//		halt
	// x is at address 0x00
	// y is at address 0x01

	// program := []*vm.Instruction{
	// 	vm.NewInstruction(vm.Const, []vm.Value{twoValue}),
	// 	vm.NewInstruction(vm.Store, []vm.Value{zeroValue}),
	// 	vm.NewInstruction(vm.Fetch, []vm.Value{zeroValue}),
	// 	vm.NewInstruction(vm.Halt, nil),
	// }

	// program := []*vm.Instruction{
	// 	vm.NewInstruction(vm.Fetch, []vm.Value{oneValue}),
	// 	vm.NewInstruction(vm.Fetch, []vm.Value{twoValue}),
	// 	vm.NewInstruction(vm.Add, nil),
	// 	vm.NewInstruction(vm.Return, nil),
	// 	vm.NewInstruction(vm.Const, []vm.Value{tenValue}),
	// 	vm.NewInstruction(vm.Const, []vm.Value{twoValue}),
	// 	vm.NewInstruction(vm.Store, []vm.Value{oneValue}),
	// 	vm.NewInstruction(vm.Store, []vm.Value{twoValue}),
	// 	vm.NewInstruction(vm.Call, []vm.Value{functionLocation, twoValue}),
	// 	vm.NewInstruction(vm.Halt, nil),
	// }

	// Testing adding two incompatible values
	// program := []*vm.Instruction{
	// 	vm.NewInstruction(vm.Const, []vm.Value{stringValue}),
	// 	vm.NewInstruction(vm.Const, []vm.Value{twoValue}),
	// 	vm.NewInstruction(vm.Add, nil),
	// 	vm.NewInstruction(vm.Halt, nil),
	// }

	// Testing const with wrong number of operands
	// program := []*vm.Instruction{
	// 	vm.NewInstruction(vm.Pop, []vm.Value{}),
	// 	vm.NewInstruction(vm.Halt, nil),
	// }

	// const 10
	// out
	// ret
	// call 0, 0
	// halt

	// program := []*vm.Instruction{
	// 	vm.NewInstruction(vm.Const, []vm.Value{tenValue}),
	// 	vm.NewInstruction(vm.Return, nil),
	// 	vm.NewInstruction(vm.Call, []vm.Value{functionLocation, zeroValue}),
	// 	vm.NewInstruction(vm.Halt, nil),
	// }
	// runner := vm.NewRunner(program, 10, 0, *trace)
	// runner.Run()

	fmt.Println(vm.Assemble("main: const 10\nconst main"))
}
