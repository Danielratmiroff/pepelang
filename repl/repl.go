package repl

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"ratmy/evaluator"
	"ratmy/lexer"
	"ratmy/object"
	"ratmy/parser"
	"strings"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		input := readExpectedInput(line)

		l := lexer.New(input)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

// Read line or file
func readExpectedInput(line string) string {
	fields := strings.Fields(line)
	if fields[0] == "pp" {
		if len(fields) > 2 {
			log.Fatalf("Too many arguments, expected 2, got:%d", len(fields))
		}

		f := fields[1]
		fileBytes, err := ioutil.ReadFile(f)
		if err != nil {
			log.Fatal("Error reading file:", err)
		}

		return string(fileBytes)

	} else {
		return line
	}
}

// Prints error message to the console
func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
