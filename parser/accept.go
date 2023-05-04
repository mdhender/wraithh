// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package parser

import (
	"fmt"
)

func (p *Parser) Accept(k Kind) (node Node) {
	saved := p.clone()
	p.skipsp()
	switch k {
	case CARGO:
		node = p.acceptProduct()
		if node == nil {
			p.restore(saved)
			node = p.acceptProfession()
		}
	case COMMA:
		node = p.acceptComma()
	case COMMAND:
		node = p.acceptCommand()
	case EOF:
		node = p.acceptEOF()
	case EOL:
		node = p.acceptEOL()
	case ID:
		node = p.acceptID()
	case INTEGER:
		node = p.acceptInteger()
	case LOCATION:
		node = p.acceptLocation()
	case NAME:
		node = p.acceptName()
	case PERCENTAGE:
		node = p.acceptPercentage()
	case PRODUCT:
		node = p.acceptProduct()
	case PROFESSION:
		node = p.acceptProfession()
	case QUANTITY:
		node = p.acceptQuantity()
	case RESOURCE:
		node = p.acceptResource()
	default:
		node = nil
	}
	if node == nil {
		p.restore(saved)
	}
	return node
}

func (p *Parser) Expect(k Kind) (Node, error) {
	n := p.Accept(k)
	if n == nil {
		return nil, fmt.Errorf("%d: unexpected input", p.line)
	}
	return n, nil
}
