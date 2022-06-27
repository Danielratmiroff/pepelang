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
		fields := strings.Fields(line)
		if fields[0] == "r" {
			// todo:
			// pass a custom file name
			// check for correct file extension
			// reuse the lexer/parser code (below as well)
			fileBytes, err := ioutil.ReadFile("./hey.txt")
			if err != nil {
				log.Fatal("Error reading file:", err)
			}

			line := (string(fileBytes))

			l := lexer.New(line)
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

		} else {
			l := lexer.New(line)
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
}

// Prints error message to the console
func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
