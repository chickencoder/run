package vm

import "fmt"

// ValueKind represents either a String or Number
// typed value on the stack
type ValueKind int

const (
	NumberValue ValueKind = iota
	StringValue
)

// Value represents an item on the stack
type Value struct {
	Kind    ValueKind
	Content interface{}
}

func (v Value) String() string {
	if v.Kind == StringValue {
		return fmt.Sprint(v.Content)
	} else if v.Content != nil {
		return fmt.Sprintf("%.2f", v.Content.(float64))
	} else {
		return "nil"
	}
}

// Stack data structure for storing Value items
type Stack struct {
	pointer int
	size    int
	data    []Value
}

// NewStack returns stack item
func NewStack(size int) *Stack {
	return &Stack{
		pointer: -1,
		data:    make([]Value, size),
		size:    size,
	}
}

// Push puts an item on the top of the stack
// Stack grows downwards
func (s *Stack) Push(item Value) {
	s.pointer++
	s.data[s.pointer] = item
}

// Pop removes an item from the top of the stack
// and then returns a Value type
func (s *Stack) Pop() Value {
	v := s.data[s.pointer]
	s.pointer--

	// TODO: Throw underflow error
	// Don't allow stack underflow
	if s.pointer < 0 {
		s.pointer = 0
	}
	return v
}

// Peek returns the item on the top of the stack
// without popping it off
func (s *Stack) Peek() Value {
	// Always return valid
	return s.data[s.pointer]
}

// Store takes an address, pops item off stack then stores
// it into local memory (heap growing up the stack)
func (s *Stack) Store(address Value) {
	// TODO: Throw errors
	item := s.Pop()
	addr := int(address.Content.(float64))
	s.data[len(s.data)-addr-1] = item
}

// Fetch looks up an address in memory then pushes
// the value at the address onto the stack
func (s *Stack) Fetch(address Value) {
	// TODO: throw errors
	addr := int(address.Content.(float64))
	item := s.data[len(s.data)-addr-1]
	s.Push(item)
}

// func (s *Stack) String() string {
// 	var str string
// 	for i, item := range s.data {
// 		if i == s.pointer {
// 			str += fmt.Sprintf("_%v_", item)
// 		} else {
// 			str += fmt.Sprintf("%v", item)
// 		}
// 	}
// 	return str
// }
