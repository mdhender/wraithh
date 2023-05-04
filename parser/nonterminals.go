// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package parser

func (p *Parser) acceptBlankLine() *BlankNode {
	saved := p.clone()
	if p.iseof() {
		return nil
	}
	p.skipsp()
	if p.iseof() {
		return &BlankNode{line: p.line}
	} else if p.nextch() == '\n' {
		return &BlankNode{line: p.line - 1}
	}
	p.restore(saved)
	return nil
}

func (p *Parser) acceptCommand() *CommandNode {
	saved := p.clone()
	var t []byte
	for p.isalpha(p.peekch()) {
		t = append(t, p.nextch())
	}
	if len(t) == 0 {
		p.restore(saved)
		return nil
	}
	if ch := p.peekch(); !(p.iseof() || ch == '\n' || p.isspace(ch)) {
		p.restore(saved)
		return nil
	}
	cmd := string(t)
	switch cmd {
	case "bombard":
		return &CommandNode{line: p.line, value: string(t)}
	}
	p.restore(saved)
	return nil
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
