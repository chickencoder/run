package vm

// ErrorKind describes the nature of an error message
type ErrorKind int

// Errors contains string representations of ErrorKinds
// for simple formatting
var Errors = []string{
	"StackError",
	"ValueError",
	"CodeError",
}

const (
	StackError ErrorKind = iota
	ValueError
	CodeError
)
