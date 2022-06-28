package repl

import (
	"bufio"
	"strings"
	"testing"
)

func TestReadExpectedInput(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"pp ./test_program.ppl", `let hey = fn(a) { return a + 3; } hey(1)`},
		{"5 + 5", "5 + 5"},
	}
	for _, tt := range tests {
		newRead := strings.NewReader(tt.input)
		scanner := bufio.NewScanner(newRead)

		scanned := scanner.Scan()
		if !scanned {
			t.Fatalf("Error initialising file scanner")
		}

		line := scanner.Text()
		fileLine := ReadExpectedInput(line)

		if fileLine != tt.expectedValue {
			t.Fatalf("fileLine values do not match, got=%s want=%s", fileLine, tt.expectedValue)
		}

	}
}
