package vm

import (
	"fmt"
)

// Runner represents an instance of the Run Virtual Machine
type Runner struct {
	ip      int
	fp      int
	stack   *Stack
	globals *Stack
	program []*Instruction
	trace   bool
}

// NewRunner returns reference to an instance of a Runner
func NewRunner(program []*Instruction, size int, main int, trace bool) *Runner {
	return &Runner{
		ip:      main,
		stack:   NewStack(size),
		globals: NewStack(512),
		program: program,
		trace:   trace,
	}
}

// Run will begin executing the program loaded into the Runner
func (r *Runner) Run() {
loop:
	for r.ip < len(r.program) {
		instr := r.program[r.ip]

		// Decode & Execute
		switch instr.Code {
		case Halt:
			break loop
		case Const:
			operand := instr.NextOperand()
			r.stack.Push(operand)
			r.ip++

		case Store:
			address := instr.NextOperand()
			r.stack.Store(address)
			r.ip++

		case Fetch:
			address := instr.NextOperand()
			r.stack.Fetch(address)
			r.ip++

		case GStore:
			address := instr.NextOperand()
			r.globals.Store(address)
			r.ip++

		case GFetch:
			address := instr.NextOperand()
			r.globals.Fetch(address)
			r.ip++

		case Pop:
			r.stack.Pop()
			r.ip++

		case Add:
			a := r.stack.Pop()
			b := r.stack.Pop()
			if a.Kind == NumberValue && b.Kind == NumberValue {
				result := Value{
					Kind:    NumberValue,
					Content: a.Content.(float64) + b.Content.(float64),
				}
				r.stack.Push(result)
			}
			// Throw error
			r.ip++

		case Sub:
			a := r.stack.Pop()
			b := r.stack.Pop()
			if a.Kind == NumberValue && b.Kind == NumberValue {
				result := Value{
					Kind:    NumberValue,
					Content: b.Content.(float64) - a.Content.(float64),
				}
				r.stack.Push(result)
			}
			// Throw error
			r.ip++

		case Mul:
			a := r.stack.Pop()
			b := r.stack.Pop()
			if a.Kind == NumberValue && b.Kind == NumberValue {
				result := Value{
					Kind:    NumberValue,
					Content: a.Content.(float64) * b.Content.(float64),
				}
				r.stack.Push(result)
			}
			// Throw error
			r.ip++

		case Div:
			a := r.stack.Pop()
			b := r.stack.Pop()
			if a.Kind == NumberValue && b.Kind == NumberValue {
				result := Value{
					Kind:    NumberValue,
					Content: b.Content.(float64) / a.Content.(float64),
				}
				r.stack.Push(result)
			}
			// Throw error
			r.ip++

		case And:
			a := r.stack.Pop()
			b := r.stack.Pop()
			if a.Kind == NumberValue && b.Kind == NumberValue {
				result := Value{
					Kind:    NumberValue,
					Content: int(a.Content.(float64)) & int(b.Content.(float64)),
				}
				r.stack.Push(result)
			}
			// Throw error
			r.ip++

		case Or:
			a := r.stack.Pop()
			b := r.stack.Pop()
			if a.Kind == NumberValue && b.Kind == NumberValue {
				result := Value{
					Kind:    NumberValue,
					Content: int(a.Content.(float64)) | int(b.Content.(float64)),
				}
				r.stack.Push(result)
			}
			r.ip++

		case Xor:
			a := r.stack.Pop()
			b := r.stack.Pop()
			if a.Kind == NumberValue && b.Kind == NumberValue {
				result := Value{
					Kind:    NumberValue,
					Content: int(a.Content.(float64)) ^ int(b.Content.(float64)),
				}
				r.stack.Push(result)
			}
			r.ip++

		case IfEqual:
			addr := instr.NextOperand()
			a := r.stack.Pop()
			b := r.stack.Pop()
			if a.Kind == NumberValue && b.Kind == NumberValue {
				if a == b {
					r.ip = int(addr.Content.(float64))
				}
			}
			// Else string type
			if a == b {
				r.ip = int(addr.Content.(float64))
			}

		case IfLessThan:
			addr := instr.NextOperand()
			a := r.stack.Pop()
			b := r.stack.Pop()
			if a.Kind == NumberValue && b.Kind == NumberValue {
				if a.Content.(float64) < b.Content.(float64) {
					r.ip = int(addr.Content.(float64))
				}
			}
			// Throw Error if not number

		case IfLessThanOrEqual:
			addr := instr.NextOperand()
			a := r.stack.Pop()
			b := r.stack.Pop()
			if a.Kind == NumberValue && b.Kind == NumberValue {
				if a.Content.(float64) <= b.Content.(float64) {
					r.ip = int(addr.Content.(float64))
				}
			}
			// Throw Error if not number

		case IfGreaterThan:
			addr := instr.NextOperand()
			a := r.stack.Pop()
			b := r.stack.Pop()
			if a.Kind == NumberValue && b.Kind == NumberValue {
				if a.Content.(float64) > b.Content.(float64) {
					r.ip = int(addr.Content.(float64))
				}
			}
			// Throw Error if not number

		case IfGreaterThanOrEqual:
			addr := instr.NextOperand()
			a := r.stack.Pop()
			b := r.stack.Pop()
			if a.Kind == NumberValue && b.Kind == NumberValue {
				if a.Content.(float64) >= b.Content.(float64) {
					r.ip = int(addr.Content.(float64))
				}
			}
			// Throw Error if not number

		case Goto:
			addr := instr.NextOperand()
			r.ip = int(addr.Content.(float64))

		case Call:
			// TODO: Add error checking when pushing/popping
			// args are expected to be on the stack already
			addr := instr.NextOperand()
			nargs := instr.NextOperand()
			fpVal := Value{
				Kind:    NumberValue,
				Content: float64(r.fp),
			}
			ipVal := Value{
				Kind:    NumberValue,
				Content: float64(r.ip),
			}
			r.stack.Push(nargs)
			r.stack.Push(fpVal)
			r.stack.Push(ipVal)

			r.fp = r.stack.pointer // fp points to the return address on the stack
			r.ip = int(addr.Content.(float64))

		case Return:
			// TODO: add error checking
			retVal := r.stack.Pop()
			r.stack.pointer = r.fp
			r.ip = int(r.stack.Pop().Content.(float64))
			r.fp = int(r.stack.Pop().Content.(float64))
			nargs := int(r.stack.Pop().Content.(float64))

			// Pop off all args
			for i := 0; i < nargs; i++ {
				r.stack.Pop()
			}

			// Leave result on stack
			r.stack.Push(retVal)
			r.ip++

		default:
			fmt.Printf("Unrecognised opcode %d\n", instr.Code)
			break loop
		}

		if r.trace {
			out := instr.Display()
			if out != "" {
				fmt.Printf("%04d ", r.ip)
				fmt.Print(instr.Display())
				fmt.Printf("\tstack %v \t(%v)\n", r.stack.data, r.stack.Peek())
			}
		}
	}
}
