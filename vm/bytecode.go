package vm

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

	// Function Instructions
	Call // location, n args
	Return
)

// Opcode represents the directive of a bytecode instruction
type Opcode int

// Operand represents a stack value
type Operand struct {
	Value interface{}
	Type  string
}

// NewOperand returns Operand instance
func NewOperand(val interface{}, typ string) Operand {
	return Operand{
		Value: val,
		Type:  typ,
	}
}

// Instruction represents a printable Opcode or Operand
type Instruction interface {
	Display() string
}

// Display method implements Instruction.Print for Opcodes
func (o Opcode) Display() string {
	return Instructions[o]
}

// Display method implements Instruction.Print for Operators
func (o Operand) Display() string {
	// if o.Type == "string" {
	// 	return o.Value.(string)
	// } else if o.Type == "number" {
	// 	num := o.Value.(float64)
	// 	return strconv.FormatFloat(num, 'f', 6, 64)
	// }
	// return "undefined"
	return ""
}
