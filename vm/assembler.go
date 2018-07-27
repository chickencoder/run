package vm

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Each instruction is on a seperate lines
// lines that start with a '#' are commented and are ignored
// if the first word on a line is suffixed with a colon,
// then the label is replaced with the current ip and all instances
// of that label are replaced by the literal address

// Assemble scans a source string into a slice of
// instructions that can be fed into a vm instance
func Assemble(source string) []Instruction {
	var instructions []Instruction

	for count, line := range strings.Split(strings.TrimSuffix(source, "\n"), "\n") {
		words := strings.Split(line, " ")

		if strings.HasSuffix(strings.TrimSuffix(words[0], " "), ":") {
			label := strings.TrimSuffix(words[0], ":")
			address := strconv.Itoa(count)
			source = strings.Replace(source, words[0], "", -1)   // Remove all inplace labels
			source = strings.Replace(source, label, address, -1) // Replace labels with addresses
			words = words[1:]                                    // Remove label from words
		}

		var current *Instruction
		opcode := words[0]
		operands := words[1:]

		for i, op := range Instructions {
			if op == opcode {
				code := Opcode(i)

				if len(operands) > 0 {
					var ops []Value
					r, err := regexp.Compile(`(\d+(\.\d+)?)`)

					if err != nil {
						fmt.Printf("regexp did not compile")
						os.Exit(1)
					}

					for _, op := range operands {
						if strings.HasPrefix("\"", op) && strings.HasSuffix("\"", op) {
							val := strings.TrimPrefix(op, "\"")
							val = strings.TrimSuffix(val, "\"")

							ops = append(ops, Value{
								Kind:    StringValue,
								Content: val,
							})
						} else if r.MatchString(op) {
							num, _ := strconv.ParseFloat(op, 64)

							ops = append(ops, Value{
								Kind:    NumberValue,
								Content: num,
							})
						} else {
							fmt.Println("Assembling Error")
							os.Exit(1)
						}
					}

					current = NewInstruction(code, ops)
					break
				}
			}
		}

		instructions = append(instructions, *current)
	}

	fmt.Println(source)
	return instructions
}
