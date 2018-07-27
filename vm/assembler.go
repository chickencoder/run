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

var instructionOperand = map[string]int{
	"halt":   0,
	"const":  1,
	"store":  1,
	"fetch":  1,
	"gstore": 1,
	"gfetch": 1,
	"pop":    0,
	"add":    0,
	"sub":    0,
	"mul":    0,
	"div":    0,
	"and":    0,
	"or":     0,
	"xor":    0,
	"ifeq":   0,
	"lt":     0,
	"lte":    0,
	"gt":     0,
	"gte":    0,
	"goto":   1,
	"print":  0, // temporary instruction
	"call":   2,
	"ret":    0,
}

func indexOf(element string, elements []string) int {
	for index, elem := range elements {
		if element == elem {
			return index
		}
	}
	return -1
}

func isNumber(token string) bool {
	r, err := regexp.Compile(`(\d+(\.\d+)?)`)
	if err != nil {
		fmt.Println("Assembler: Internal Error")
		os.Exit(1)
	}
	return r.MatchString(token)
}

func isWord(char string) bool {
	r, err := regexp.Compile(`(\S+)`)
	if err != nil {
		fmt.Println("Assembler: Internal Error")
		os.Exit(1)
	}

	return r.MatchString(char)
}

func isQuote(char string) bool {
	return char == `"`
}

func isNotQuote(char string) bool {
	return char != `"`
}

func tokenize(source string) []string {
	var tokens []string
	current := 0

	// Remove commented lines
	lines := strings.Split(source, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			source = strings.Replace(source, line, "", -1)
		}
	}

	// Remove unecessary whitespace
	source = strings.Replace(source, "\n", " ", -1)
	source = strings.Replace(source, "\t", " ", -1)
	source += "\n"

	for current < len([]rune(source)) {
		char := string([]rune(source)[current])

		if char == " " {
			current++
			continue
		}

		if isQuote(char) {
			value := ""

			current++
			char = string([]rune(source)[current])

			for isNotQuote(char) {
				value += char
				current++
				char = string([]rune(source)[current])
			}

			current++
			value = `"` + value + `"`

			tokens = append(tokens, value)
			continue
		}

		if isWord(char) {
			value := ""

			for isWord(char) {
				value += char
				current++
				char = string([]rune(source)[current])
			}

			tokens = append(tokens, value)
			continue
		}
		break
	}

	return tokens
}

func parseOperand(token string) Value {
	// Is operand a string?
	if token == "nil" {
		return Nil
	} else if strings.HasPrefix(token, "\"") && strings.HasSuffix(token, "\"") {
		innerString := strings.TrimSuffix(strings.TrimPrefix(token, "\""), "\"")
		return Value{
			Kind:    StringValue,
			Content: innerString,
		}
	} else if isNumber(token) {
		val, err := strconv.ParseFloat(token, 64)
		if err != nil {
			fmt.Println("Assembler: Couldn't parse float")
			os.Exit(1)
		}
		return Value{
			Kind:    NumberValue,
			Content: val,
		}
	}
	fmt.Printf("Assembler: Could not parse operand: %s\n", token)
	os.Exit(1)

	return Nil // Will never return
}

// Assemble scans a source string into a slice of
// instructions that can be fed into a vm instance
func Assemble(source string) []*Instruction {
	var instructions []*Instruction
	var ip int
	var count int
	tokens := tokenize(source)

	for count < len(tokens) {
		// Only increment ip once we've successfully
		// consumed an instruction
		token := tokens[count]

		// Is token a label
		if strings.HasSuffix(token, ":") {
			label := strings.TrimSuffix(token, ":")
			address := strconv.Itoa(ip)

			// Replace all instances of label with literal address
			for i, item := range tokens {
				if item == label {
					tokens[i] = address
				}
			}
		}

		// Is token an opcode
		if indexOf(token, Instructions) != -1 {
			var operands []Value

			// Skip over operand
			// count++

			// Determine how many operands are required
			nops := instructionOperand[token]

			// Parse Operands
			for i := 0; i < nops; i++ {
				count++
				operands = append(operands, parseOperand(tokens[count]))
			}

			// Append to instructions
			instructions = append(instructions, NewInstruction(
				Opcode(indexOf(token, Instructions)), operands,
			))

			// Increment the ip
			ip++
		}

		count++
	}

	return instructions
}
