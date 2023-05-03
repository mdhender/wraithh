// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package parser

import "os"

func ParseFile(name string) (Nodes, error) {
	input, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return Parse(input), nil
}

func Parse(input []byte) Nodes {
	var nodes Nodes

	p := &Parser{line: 1, buffer: input}
	for !p.iseof() {
		if p.acceptBlankLine() != nil {
			continue
		}
		var n Node
		n = p.acceptOrder()
		if n == nil {
			n = p.expectLine()
		}
		nodes = append(nodes, n)
	}

	if p.acceptEOF() == nil {
		panic("assert(eof)")
	}

	return nodes
}

func (p *Parser) acceptOrder() Node {
	saved := p.clone()

	// command must be first word on the line
	var t []byte
	for p.isalpha(p.peekch()) {
		t = append(t, p.nextch())
	}
	if len(t) == 0 {
		p.restore(saved)
		return nil
	}
	// command must be terminated with eof, eol, or a space
	if !(p.iseof() || p.peekch() == '\n' || p.isspace(p.peekch())) {
		p.restore(saved)
		return nil
	}

	switch string(t) {
	case "assemble":
		return p.expectAssemble()
	case "bombard":
		return p.expectBombard()
	case "invade":
		return p.expectInvade()
	case "raid":
		return p.expectRaid()
	case "setup":
	case "support":
		return p.expectSupport()
	}
	p.restore(saved)
	return nil
}
