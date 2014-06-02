package main

/*
#include <stdio.h>
*/
import "C"
import (
	"fmt"
	"github.com/k0kubun/go-termios"
)

func main() {
	var saveTerm, tempTerm termios.Termios

	// Get current terminal parameters
	if err := saveTerm.GetAttr(termios.Stdin); err != nil {
		panic(err)
	}
	tempTerm = saveTerm

	// Change terminal parameters
	tempTerm.IFlag &= termios.IGNCR  // ignore received CR
	tempTerm.LFlag ^= termios.ICANON // disable canonical mode
	tempTerm.LFlag ^= termios.ECHO   // disable echo of input
	tempTerm.LFlag ^= termios.ISIG   // disable signal
	tempTerm.CC[termios.VMIN] = 1    // number of bytes to read()
	tempTerm.CC[termios.VTIME] = 0   // timeout of read()
	if err := tempTerm.SetAttr(termios.Stdin, termios.TCSANOW); err != nil {
		panic(err)
	}

	// Read each key input
	readKey()

	// Set original terminal parameters
	if err := saveTerm.SetAttr(termios.Stdin, termios.TCSANOW); err != nil {
		panic(err)
	}
}

func readKey() {
	println("Input some keys (hit 'q' to finish):")

	var ch byte
	for ch != 'q' {
		ch = getc()
		if ch != '\r' && ch != '\n' {
			fmt.Printf("input = %c\n", ch)
		}
	}
}

func getc() byte {
	return byte(C.fgetc(C.stdin))
}
