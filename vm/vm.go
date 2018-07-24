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
	program []Instruction
	trace   bool
}

// NewRunner returns reference to an instance of a Runner
func NewRunner(program []Instruction, size int, main int, trace bool) *Runner {
	return &Runner{
		ip:      main,
		stack:   NewStack(size),
		globals: NewStack(512),
		program: program,
		trace:   trace,
	}
}

func (r Runner) nextOperand() Operand {
	r.ip++
	op := r.program[r.ip]
	return op.(Operand)
}

// Run will begin executing the program loaded into the Runner
func (r *Runner) Run() {
loop:
	for r.ip < len(r.program) {
		// Fetch
		instr := r.program[r.ip]

		// Decode & Execute
		switch instr {
		case Halt:
			break loop
		case Const:
			operand := r.nextOperand().Value
			r.stack.Push(operand)
		case Store:
			address := r.nextOperand()
			r.stack.Store(address)
		case Fetch:
			address := r.nextOperand()
			r.stack.Fetch(address)
		case GStore:
			address := r.nextOperand()
			r.globals.Store(address)
		case GFetch:
			address := r.nextOperand()
			r.globals.Fetch(address)
		case Pop:
			r.stack.Pop()
		case Add:
			a := r.stack.Pop().(float64)
			b := r.stack.Pop().(float64)
			r.stack.Push(a + b)
		case Sub:
			a := r.stack.Pop().(float64)
			b := r.stack.Pop().(float64)
			r.stack.Push(b - a)
		case Mul:
			a := r.stack.Pop().(float64)
			b := r.stack.Pop().(float64)
			r.stack.Push(a * b)
		case Div:
			a := r.stack.Pop().(float64)
			b := r.stack.Pop().(float64)
			r.stack.Push(b / a)
		case And:
			a := r.stack.Pop().(float64)
			b := r.stack.Pop().(float64)
			r.stack.Push(byte(b) & byte(a)) // TODO: Is this the right thing to #TODO?
		case Or:
			a := r.stack.Pop().(float64)
			b := r.stack.Pop().(float64)
			r.stack.Push(byte(b) | byte(a))
		case Xor:
			a := r.stack.Pop().(float64)
			b := r.stack.Pop().(float64)
			r.stack.Push(byte(b) ^ byte(a))
		case IfEqual:
			addr := r.nextOperand()
			a := r.stack.Pop().(float64)
			b := r.stack.Pop().(float64)
			if a == b {
				r.ip = int(addr.Value.(float64))
			}
		case IfLessThan:
			addr := r.nextOperand()
			a := r.stack.Pop().(float64)
			b := r.stack.Pop().(float64)
			if a < b {
				r.ip = int(addr.Value.(float64))
			}
		case IfLessThanOrEqual:
			addr := r.nextOperand()
			a := r.stack.Pop().(float64)
			b := r.stack.Pop().(float64)
			if a <= b {
				r.ip = int(addr.Value.(float64))
			}
		case IfGreaterThan:
			addr := r.nextOperand()
			a := r.stack.Pop().(float64)
			b := r.stack.Pop().(float64)
			if a > b {
				r.ip = int(addr.Value.(float64))
			}
		case IfGreaterThanOrEqual:
			addr := r.nextOperand()
			a := r.stack.Pop().(float64)
			b := r.stack.Pop().(float64)
			if a >= b {
				r.ip = int(addr.Value.(float64))
			}
		case Goto:
			addr := r.nextOperand()
			r.ip = int(addr.Value.(float64))
		case Call:
			// args are expected to be on the stack already
			addr := r.nextOperand()
			nargs := r.nextOperand()
			r.stack.Push(nargs.Value.(float64))
			r.stack.Push(float64(r.fp))
			r.stack.Push(float64(r.ip))
			r.fp = r.stack.pointer // fp points to the return address on the stack
			r.ip = int(addr.Value.(float64))
		case Return:
			retVal := r.stack.Pop()
			r.stack.pointer = r.fp
			r.ip = int(r.stack.Pop().(float64))
			r.fp = int(r.stack.Pop().(float64))
			nargs := int(r.stack.Pop().(float64))

			// Pop off all args
			for i := 0; i < nargs; i++ {
				r.stack.Pop()
			}

			// Leave result on stack
			r.stack.Push(retVal)
		}

		if r.trace {
			// Print current instruction
			out := instr.Display()
			if out != "" {
				fmt.Printf("%04d: ", r.ip)
				fmt.Print(instr.Display())
				fmt.Printf("\tstack=%v \t(%v)\n", r.stack.data, r.stack.Peek())
			}

			// Print stack frame
			// fmt.Printf("%v", r.stack)
		}

		// Repeat
		r.ip++
	}
}
