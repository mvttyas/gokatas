// Package count counts lines. It was created with a test-first approach (TDD).
// It provides a default input, which is STDIN, with the possibility to change
// it. Source: https://bitfieldconsulting.com/books/tools
//
// If you want to climb a mountain, begin at the top. -- Zen saying
//
// Level: intermediate
// Topics: default options, scripting
package count

import (
	"bufio"
	"io"
	"os"
)

type Counter struct {
	Input io.Reader
}

func NewCounter() Counter {
	return Counter{Input: os.Stdin}
}

func (c Counter) Lines() int {
	var lines int
	scanner := bufio.NewScanner(c.Input)
	for scanner.Scan() {
		lines++
	}
	return lines
}