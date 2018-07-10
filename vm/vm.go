package vm

type InterpretResult int

const (
	INTERPRET_OK InterpretResult = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

// Process represents the state of a single
// process executed on the virtual machine
type Process struct {
	chunk Chunk
}

func (p *Process) interpret(chunk Chunck) InterpretResult {
	p.chunk = chunk
	vm.ip
}
