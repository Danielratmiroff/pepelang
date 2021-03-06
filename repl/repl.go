package repl

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"pepe/evaluator"
	"pepe/lexer"
	"pepe/object"
	"pepe/parser"
	"strings"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	if len(os.Args) == 2 {
		args := os.Args[1]
		line := fmt.Sprintf("pp %v", args)
		input := ReadUserInput(line)
		startProgram(input, out, env)
	}

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		input := ReadUserInput(line)
		startProgram(input, out, env)
	}
}

func startProgram(input string, out io.Writer, env *object.Environment) {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParseErrors(out, p.Errors())
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}

// Read line or file
func ReadUserInput(line string) string {
	fields := strings.Fields(line)
	if fields[0] == "pp" {
		if len(fields) > 2 {
			log.Fatalf("Too many arguments, got=%d want=2", len(fields))
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
