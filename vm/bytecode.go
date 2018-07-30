package vm

import "fmt"

// Opcode represents the directive of a bytecode instruction
type Opcode int

// Instructions maps Opcode ints to instruction strings
var Instructions = []string{
	"halt",
	"const",
	"store",
	"fetch",
	"gstore",
	"gfetch",
	"pop",
	"add",
	"sub",
	"mul",
	"div",
	"and",
	"or",
	"xor",
	"ifeq",
	"lt",
	"lte",
	"gt",
	"gte",
	"goto",
	"print", // temporary instruction
	"call",
	"ret",
}

// Instruction declarations
const (
	Halt Opcode = iota

	Const  // Pushes constant onto stack
	Store  // Stores top of stack into local variable
	Fetch  // Pushes local variable onto stack
	GStore // Pops item and creates as a global
	GFetch // Pushes items from globals onto stack
	Pop    // Pop Item from stack

	Add // Pushes sum of top two items on the stack
	Sub // Pushes substraction of top two items on the stack
	Mul // Pushes multiplication of top two items on the stack
	Div // Pushes division of top two items on the stack
	And // Bitwise AND
	Or  // Bitwise OR
	Xor // Bitwise XOR

	// Control Instructions
	IfEqual
	IfLessThan
	IfLessThanOrEqual
	IfGreaterThan
	IfGreaterThanOrEqual
	Goto

	Print // Temporary instruction for debugging purposes

	// Function Instructions
	Call // location, n args
	Return
)

// Instruction is an Opcode and optional Operand(s)
type Instruction struct {
	Code     Opcode
	Operands []Value
	Index    int
}

// NewInstruction returns reference to a Instruction
func NewInstruction(op Opcode, ops []Value) *Instruction {
	return &Instruction{
		Code:     op,
		Operands: ops,
		Index:    0,
	}
}

// Display returns a printed string of an instruction
func (i *Instruction) Display() string {
	line := Instructions[i.Code]
	if len(i.Operands) < 1 {
		line += "\t\t"
		return line
	} else if len(i.Operands) == 1 {
		line += "\t"
	}

	for _, op := range i.Operands {
		line += "\t"
		if op.Kind == StringValue {
			line += op.Content.(string)
		} else if op.Kind == NumberValue {
			line += fmt.Sprintf("%.2f", op.Content.(float64))
		}
	}

	return line
}

// NextOperand returns the next operand within the instruction
func (i *Instruction) NextOperand() Value {
	if i.Index < len(i.Operands)-1 {
		i.Index++
	} else {
		i.Index = 0
	}
	val := i.Operands[i.Index]
	return val
}
