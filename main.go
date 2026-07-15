package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type Instruction struct {
	Type      rune
	Code      float64
	DebugLine string
}

type Program struct {
	Instructions []Instruction
}

func ParseLine(line string) (isNull bool, i Instruction) {
	if len(line) == 0 {
		isNull = true
		return
	}
	var first = rune(line[0])
	switch first {
	case 'G':
		return false, ParseGCode(line)
	case 'M':
		// TODO: implement M
		return true, Instruction{}
	case ';':
		return true, Instruction{}
	default:
		panic(fmt.Sprintf("Don't know how to parse the line: \"%s\"", line))
	}
}

func (p *Program) Parse(lines []string) {
	for _, line := range lines {
		isNull, i := ParseLine(line)
		if !isNull {
			p.Instructions = append(p.Instructions, i)
		}
	}
}

func main() {
	var args = os.Args
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage\n")
		os.Exit(1)
	}

	osReader, err := os.Open(args[1])
	if err != nil {
		panic(err)
	}
	defer osReader.Close()

	var reader = bufio.NewReader(osReader)
	var wasPrefix = false
	var lastLine []byte
	var lines []string

	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}

		if wasPrefix {
			line = append(lastLine, line...)
		}
		if isPrefix {
			lastLine = line
			// TODO: track line number?
			fmt.Fprintf(os.Stderr, "DEBUG: line exceeded buffer size\n")
		} else {
			lines = append(lines, string(line))
		}
		wasPrefix = isPrefix
	}

	var p Program
	p.Parse(lines)
}
