package main

import (
	"os"

	"millanuka.com/lisp-interpreter/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
