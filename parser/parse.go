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
