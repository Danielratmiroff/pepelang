package main

import (
	"fmt"
	"os"
	"os/user"
	"pepe/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hola %s! Bienvenido a Pepelang!!\n",
		user.Username)
	fmt.Printf("Puedes escribir cualquier comando aca :)\n")

	repl.Start(os.Stdin, os.Stdout)
}
