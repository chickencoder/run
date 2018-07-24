package vm

// Value represents an item on the stack
type Value interface{}

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
	return v
}

// Store takes an address, pops item off stack then stores
// it into local memory (heap growing up the stack)
func (s *Stack) Store(address Operand) {
	item := s.Pop()
	addr := address.Value.(int)
	s.data[len(s.data)-addr] = item
}

// Fetch looks up an address in memory then pushes
// the value at the address onto the stack
func (s *Stack) Fetch(address Operand) {
	addr := address.Value.(int)
	item := s.data[len(s.data)-addr]
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
