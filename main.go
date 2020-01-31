package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	var c string
	flag.StringVar(&c, "code", "", "brainfuck program to execute")

	flag.Parse()

	// _, tokens := lex(c)
	// i := NewInterpreter(tokens, newMem(), os.Stdout, os.Stdin)
	m := newMem()
	i := New(c, os.Stdout, os.Stdin)
	i.Run(m)
	log.Printf("%#v", m)
}
