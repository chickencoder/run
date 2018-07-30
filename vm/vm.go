package vm

import (
	"fmt"
	"os"
)

// Runner represents an instance of the Run Virtual Machine
type Runner struct {
	ip      int
	fp      int
	stack   *Stack
	globals *Stack
	program []*Instruction
	trace   bool
	panic   bool
}

// NewRunner returns reference to an instance of a Runner
func NewRunner(program []*Instruction, size int, main int, trace bool) *Runner {
	return &Runner{
		ip:      main,
		stack:   NewStack(size),
		globals: NewStack(512),
		program: program,
		trace:   trace,
		panic:   false,
	}
}

// Throw will display a runtime error message
func Throw(kind ErrorKind, message string) {
	fmt.Printf("%s: %s\n", Errors[kind], message)
	os.Exit(1)
}

// Run will begin executing the program loaded into the Runner
func (r *Runner) Run() {
loop:
	for !r.panic && r.ip < len(r.program) {
		instr := r.program[r.ip]

		// Decode & Execute
		switch instr.Code {
		case Halt:
			break loop
		case Const:
			operand := instr.NextOperand()
			if operand == Nil {
				Throw(CodeError, fmt.Sprintf("expected operand from %s", instr.Display()))
			}
			item := r.stack.Push(operand)
			if item == Nil {
				Throw(StackError, "cannot add because stack is full")
			}
			r.ip++

		case Store:
			address := instr.NextOperand()
			if address == Nil {
				Throw(CodeError, fmt.Sprintf("expected operand from %s", instr.Display()))
			}
			// Check for nil value
			r.stack.Store(address)
			r.ip++

		case Fetch:
			address := instr.NextOperand()
			if address == Nil {
				Throw(CodeError, fmt.Sprintf("expected operand from %s", instr.Display()))
			}
			// Check for nil value
			r.stack.Fetch(address)
			r.ip++

		case GStore:
			address := instr.NextOperand()
			if address == Nil {
				Throw(CodeError, fmt.Sprintf("expected operand from %s", instr.Display()))
			}
			// Check for nil value
			r.globals.Store(address)
			r.ip++

		case GFetch:
			address := instr.NextOperand()
			if address == Nil {
				Throw(CodeError, fmt.Sprintf("expected operand from %s", instr.Display()))
			}
			// Check for nil value
			r.globals.Fetch(address)
			r.ip++

		case Pop:
			item := r.stack.Pop()
			if item == Nil {
				Throw(StackError, "cannot pop because stack is empty")
			}
			r.ip++

		case Add:
			a := r.stack.Pop()
			b := r.stack.Pop()

			if a == Nil || b == Nil {
				Throw(StackError, "cannot add because stack is empty")
			}

			if a.Kind == NumberValue && b.Kind == NumberValue {
				result := Value{
					Kind:    NumberValue,
					Content: a.Content.(float64) + b.Content.(float64),
				}
				item := r.stack.Push(result)
				if item == Nil {
					Throw(StackError, "cannot add because stack is full")
				}
			} else {
				Throw(ValueError, fmt.Sprintf("cannot add %s value to %s value", ValueKinds[a.Kind], ValueKinds[b.Kind]))
			}
			r.ip++

		case Sub:
			a := r.stack.Pop()
			b := r.stack.Pop()

			if a == Nil || b == Nil {
				Throw(StackError, "cannot sub because stack is empty")
			}

			if a.Kind == NumberValue && b.Kind == NumberValue {
				result := Value{
					Kind:    NumberValue,
					Content: b.Content.(float64) - a.Content.(float64),
				}
				item := r.stack.Push(result)
				if item == Nil {
					Throw(StackError, "cannot add because stack is full")
				}
				r.ip++
			} else {
				Throw(ValueError, fmt.Sprintf("cannot sub %s value from %s value", ValueKinds[b.Kind], ValueKinds[a.Kind]))
			}

		case Mul:
			a := r.stack.Pop()
			b := r.stack.Pop()

			if a == Nil || b == Nil {
				Throw(StackError, "cannot mul because stack is empty")
			}

			if a.Kind == NumberValue && b.Kind == NumberValue {
				result := Value{
					Kind:    NumberValue,
					Content: a.Content.(float64) * b.Content.(float64),
				}
				item := r.stack.Push(result)
				if item == Nil {
					Throw(StackError, "cannot add because stack is full")
				}
				r.ip++
			} else {
				Throw(ValueError, fmt.Sprintf("cannot mul %s value with %s value", ValueKinds[a.Kind], ValueKinds[b.Kind]))
			}

		case Div:
			a := r.stack.Pop()
			b := r.stack.Pop()

			if a == Nil || b == Nil {
				Throw(StackError, "cannot div because stack is empty")
			}

			if a.Kind == NumberValue && b.Kind == NumberValue {
				result := Value{
					Kind:    NumberValue,
					Content: b.Content.(float64) / a.Content.(float64),
				}
				item := r.stack.Push(result)
				if item == Nil {
					Throw(StackError, "cannot add because stack is full")
				}
				r.ip++
			} else {
				Throw(ValueError, fmt.Sprintf("cannot div %s value by %s value", ValueKinds[b.Kind], ValueKinds[a.Kind]))
			}

		case And:
			a := r.stack.Pop()
			b := r.stack.Pop()

			if a == Nil || b == Nil {
				Throw(StackError, "cannot and because stack is empty")
			}

			if a.Kind == NumberValue && b.Kind == NumberValue {
				result := Value{
					Kind:    NumberValue,
					Content: int(a.Content.(float64)) & int(b.Content.(float64)),
				}
				item := r.stack.Push(result)
				if item == Nil {
					Throw(StackError, "cannot add because stack is full")
				}
				r.ip++
			} else {
				Throw(ValueError, fmt.Sprintf("cannot and %s value with %s value", ValueKinds[a.Kind], ValueKinds[b.Kind]))
			}

		case Or:
			a := r.stack.Pop()
			b := r.stack.Pop()

			if a == Nil || b == Nil {
				Throw(StackError, "cannot or because stack is empty")
			}

			if a.Kind == NumberValue && b.Kind == NumberValue {
				result := Value{
					Kind:    NumberValue,
					Content: int(a.Content.(float64)) | int(b.Content.(float64)),
				}
				item := r.stack.Push(result)
				if item == Nil {
					Throw(StackError, "cannot add because stack is full")
				}
				r.ip++
			} else {
				Throw(ValueError, fmt.Sprintf("cannot or %s value with %s value", ValueKinds[a.Kind], ValueKinds[b.Kind]))
			}

		case Xor:
			a := r.stack.Pop()
			b := r.stack.Pop()

			if a == Nil || b == Nil {
				Throw(StackError, "cannot xor because stack is empty")
			}

			if a.Kind == NumberValue && b.Kind == NumberValue {
				result := Value{
					Kind:    NumberValue,
					Content: int(a.Content.(float64)) ^ int(b.Content.(float64)),
				}
				item := r.stack.Push(result)
				if item == Nil {
					Throw(StackError, "cannot add because stack is full")
				}
				r.ip++
			} else {
				Throw(ValueError, fmt.Sprintf("cannot xor %s value with %s value", ValueKinds[a.Kind], ValueKinds[b.Kind]))
			}

		case IfEqual:
			addr := instr.NextOperand()
			a := r.stack.Pop()
			b := r.stack.Pop()

			if a == Nil || b == Nil {
				Throw(StackError, "cannot make comparison because stack is empty")
			}

			if a.Kind == NumberValue && b.Kind == NumberValue {
				if a == b {
					r.ip = int(addr.Content.(float64))
				}
			} else if a.Kind == StringValue && b.Kind == StringValue {
				if a == b {
					r.ip = int(addr.Content.(float64))
				}
			} else {
				Throw(ValueError, fmt.Sprintf("cannot make comparison between %s value and %s value", ValueKinds[a.Kind], ValueKinds[b.Kind]))
			}

		case IfLessThan:
			addr := instr.NextOperand()
			a := r.stack.Pop()
			b := r.stack.Pop()

			if a == Nil || b == Nil {
				Throw(StackError, "cannot make comparison because stack is empty")
			}

			if a.Kind == NumberValue && b.Kind == NumberValue {
				if a.Content.(float64) < b.Content.(float64) {
					r.ip = int(addr.Content.(float64))
				}
			} else {
				Throw(ValueError, fmt.Sprintf("cannot make comparison between %s value and %s value", ValueKinds[a.Kind], ValueKinds[b.Kind]))
			}

		case IfLessThanOrEqual:
			addr := instr.NextOperand()
			a := r.stack.Pop()
			b := r.stack.Pop()

			if a == Nil || b == Nil {
				Throw(StackError, "cannot make comparison because stack is empty")
			}

			if a.Kind == NumberValue && b.Kind == NumberValue {
				if a.Content.(float64) <= b.Content.(float64) {
					r.ip = int(addr.Content.(float64))
				}
			} else {
				Throw(ValueError, fmt.Sprintf("cannot make comparison between %s value and %s value", ValueKinds[a.Kind], ValueKinds[b.Kind]))
			}

		case IfGreaterThan:
			addr := instr.NextOperand()
			a := r.stack.Pop()
			b := r.stack.Pop()

			if a == Nil || b == Nil {
				Throw(StackError, "cannot make comparison because stack is empty")
			}

			if a.Kind == NumberValue && b.Kind == NumberValue {
				if a.Content.(float64) > b.Content.(float64) {
					r.ip = int(addr.Content.(float64))
				}
			} else {
				Throw(ValueError, fmt.Sprintf("cannot make comparison between %s value and %s value", ValueKinds[a.Kind], ValueKinds[b.Kind]))
			}

		case IfGreaterThanOrEqual:
			addr := instr.NextOperand()
			a := r.stack.Pop()
			b := r.stack.Pop()

			if a == Nil || b == Nil {
				Throw(StackError, "cannot make comparison because stack is empty")
			}

			if a.Kind == NumberValue && b.Kind == NumberValue {
				if a.Content.(float64) >= b.Content.(float64) {
					r.ip = int(addr.Content.(float64))
				}
			} else {
				Throw(ValueError, fmt.Sprintf("cannot make comparison between %s value and %s value", ValueKinds[a.Kind], ValueKinds[b.Kind]))
			}

		case Goto:
			addr := instr.NextOperand()
			if addr == Nil {
				Throw(CodeError, fmt.Sprintf("expected operand from %s", instr.Display()))
			}
			r.ip = int(addr.Content.(float64))

		case Call:
			// TODO: Add error checking when pushing/popping
			// args are expected to be on the stack already
			addr := instr.NextOperand()
			nargs := instr.NextOperand()

			if addr == Nil || nargs == Nil {
				Throw(CodeError, fmt.Sprintf("expected address and nargs operands from %s", instr.Display()))
			}

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
			if retVal == Nil {
				Throw(CodeError, "no value returned from function")
			}

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

		case Print:
			fmt.Println(r.stack.Peek())
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
				fmt.Printf("\tstack %v \t(%v)", r.stack.data, r.stack.Peek())
				fmt.Printf("\t*%d\n", r.stack.pointer)
			}
		}
	}
}
