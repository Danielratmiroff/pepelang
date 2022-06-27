package repl

import (
	"bufio"
	"io/ioutil"
	"ratmy/evaluator"
	"ratmy/lexer"
	"ratmy/object"
	"ratmy/parser"
	"strings"
	"testing"
)

func TestPPLFiles(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"pp test_program.ppl", 4},
	}

	for _, tt := range tests {
		newRead := strings.NewReader(tt.input)
		scanner := bufio.NewScanner(newRead)
		env := object.NewEnvironment()

		scanned := scanner.Scan()
		if !scanned {
			t.Fatalf("Error initialising file scanner")
		}

		line := scanner.Text()
		fields := strings.Fields(line)

		if fields[0] != "pp" {
			t.Fatalf("Missing read file command as input. got=%s, expected='pp'", fields[0])
		}

		if len(fields) > 2 {
			t.Fatalf("Too many input arguments got=%d want=2", len(fields))
		}

		f := fields[1]
		fileBytes, err := ioutil.ReadFile(f)
		if err != nil {
			t.Fatalf("Error reading file 'test_program.ppl' got=%s, Error: %s", tt.input, err)
		}

		fileLine := string(fileBytes)

		l := lexer.New(fileLine)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			t.Fatalf("Found parse errors: %s", p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated == nil {
			t.Fatalf("Empty code evaluation, want=%s", evaluated.Inspect())
		}

		if evaluated.Inspect() != tt.expectedValue {
			t.Fatalf("Wrong evaluated value, got=%s, want=%s", evaluated.Inspect(), tt.expectedValue)
		}

	}
}
