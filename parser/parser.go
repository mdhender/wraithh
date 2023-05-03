// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package parser implements a recursive-descent parser for orders.
package parser

import "bytes"

type Parser struct {
	buffer []byte
	line   int
}

func (p *Parser) clone() *Parser {
	return &Parser{buffer: p.buffer, line: p.line}
}

func (p *Parser) first(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n]
}

func (p *Parser) isalpha(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}

func (p *Parser) isalnum(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ('0' <= ch && ch <= '9')
}

func (p *Parser) isdigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (p *Parser) iseof() bool {
	return len(p.buffer) == 0
}

func (p *Parser) islower(ch byte) bool {
	return 'a' <= ch && ch <= 'z'
}

func (p *Parser) isspace(ch byte) bool {
	return bytes.IndexByte([]byte{' ', '\f', '\r', '\t', '\v'}, ch) != -1
}

func (p *Parser) isupper(ch byte) bool {
	return 'A' <= ch && ch <= 'Z'
}

func (p *Parser) nextch() (ch byte) {
	if p.iseof() {
		return 0
	}
	ch, p.buffer = p.buffer[0], p.buffer[1:]
	if ch == '\n' {
		p.line++
	}
	return ch
}

func (p *Parser) peekch() byte {
	if p.iseof() {
		return 0
	}
	return p.buffer[0]
}

func (p *Parser) restore(r *Parser) {
	if r != nil {
		p.buffer = r.buffer
		p.line = r.line
	}
}

func (p *Parser) skipsp() {
	for bytes.IndexByte([]byte{' ', '\f', '\r', '\t', '\v'}, p.peekch()) != -1 {
		_ = p.nextch()
	}
}
