// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package parser

type Parser struct {
	line   int
	buffer []byte
}

type Token struct {
	Line  int
	Kind  string
	Value string
}

var (
	whtspace = []byte{' ', '\t', '\f', '\r', '\v'}
)

func (p Parser) Accept(kind string) (Parser, Token, bool) {
	// skip leading spaces
	_, p.buffer = runof(p.buffer, whtspace)

	t := Token{Line: p.line}

	if len(p.buffer) == 0 {
		t.Kind = "eof"
		return p, t, kind == "eof"
	} else if p.buffer[0] == '\n' {
		t.Kind, p.line = "eol", p.line+1
		return p, t, kind == "eof"
	}

	return p, Token{Line: p.line}, false
}

// getch returns the next character
func (p *Parser) getch() byte {
	if len(p.buffer) == 0 {
		return 0
	}
	ch := p.buffer[0]
	p.buffer = p.buffer[1:]
	return ch
}

func (p *Parser) next() Token {
	// skip leading spaces
	_, p.buffer = runof(p.buffer, whtspace)

	t := Token{Line: p.line}

	if len(p.buffer) == 0 {
		t.Kind = "eof"
		return t
	} else if p.buffer[0] == '\n' {
		t.Kind, p.line = "eol", p.line+1
		return t
	}

	t.Kind = "unknown"
	return t
}

// peekch returns the next character
func (p *Parser) peekch() byte {
	if len(p.buffer) == 0 {
		return 0
	}
	return p.buffer[0]
}
