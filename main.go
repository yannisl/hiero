// Playing with hieroglyphics
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

//
type Scanner struct {
	s bufio.Scanner
}

//
func NewScanner(r io.Reader) Scanner {
	sc := *bufio.NewScanner(r)
	split := bufio.ScanRunes
	sc.Split(split)
	return Scanner{s: sc}
}

//
func (s *Scanner) Scan() bool {
	return s.s.Scan()
}

//
func (s *Scanner) Text() string {
	return s.s.Text()
}

//
func (s *Scanner) ScanToDigit() string {

	return ""

}

// Symbol is a struct that represents a block of
// hieroglyppic symbols
type Symbol struct {
	code string // alphanumeric string
	pos  int    // 1 - 4 for stcking
}

//
type symbolStack struct {
	s []Symbol
}

// Que is a que structure that represents the stacking
// of the symbols.
//
// Example:
//   [[N16], [N21*Z1], []]
// The last in from the reader will be at the bottom region.
func (sc *Scanner) emit(tok string) {
	hasColon := strings.Contains(tok, ":")
	hasAsterisk := strings.Contains(tok, "*") 
	hasBoth := hasColon && hasAsterisk

	switch {
	case hasColon && !hasBoth:
		ss:= strings.Split(tok, ":")
		typeSet(ss)
	case hasBoth: 
		ss := strings.Split(tok, ":")
		fmt.Println(ss)
	case hasBoth:
	default:
		typeSet([]string{tok})
	}
}

type Metrics struct {
	width float64
	height float64
	depth float64
}

type Glyph struct {
	angle float64
	codepoint int
	Metrics Metrics
}

// At this point we need the metrics? 
func typeSet(ss []string) {
	hlist := ss
	if len(ss)==1 {
	   fmt.Printf("\\boxone{%s}\n", hlist)
	}
	if len(ss)==2 {
	   fmt.Printf("\\vboxtwo{%s}{%s}\n", ss[0], ss[1])
	}	
}


func Lex(input string) {
	var buf bytes.Buffer
	sc := NewScanner(strings.NewReader(input))
	for sc.Scan() {
		c := sc.Text()
		switch c {
		case "\n", "\r", " ", "\t":
		case "!":
			sc.emit(buf.String())
			buf.Reset()
			fmt.Println("")
		case "-":
			sc.emit(buf.String())
			buf.Reset()
		default:
			buf.WriteString(c)

		}
	}

	if err := sc.s.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}

}

func main() {
	// See http://www.catchpenny.org/codage/ for Manuel de Codage
	const input = `M23-X1:R4-X8-Q2:D4-W17-R14-G4-R8-O29:V30-U23-
	               N26-D58-O49:Z1-F13:N31-Z2-V30:N16:N21*Z1-D45:N25!
	               N26-D58-O49:Z1-F13:N31-Z2-V30:N16:N21*Z1-D45:N26!
	               N26-D58-O49:Z1-F13:N31-Z2-V30:N16:N21*Z1-D45:N27!`
	Lex(input)
}
