package main

import (
	"fmt"
	"os"

	"github.com/rodrigo-dc/wavegen"
)

func main() {
	wave, err := wavegen.Sine(440, 44100)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print tab-separated values, ready to be used by GNU Plot
	for i, s := range wave {
		fmt.Println(i, "\t", s)
	}
}
