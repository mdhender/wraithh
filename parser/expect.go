// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package parser

import "fmt"

func (p *Parser) expectAssemble() *AssembleNode {
	o := &AssembleNode{line: p.line}
	p.skipsp()
	if id := p.acceptID(); id == nil {
		o.err = fmt.Errorf("expected id: got %q", p.first(p.expectLine().value, 12))
		return o
	} else {
		o.id = id.id
	}
	p.skipsp()
	if qty := p.acceptQuantity(); qty == nil {
		o.err = fmt.Errorf("expected quantity: got %q", p.first(p.expectLine().value, 12))
		return o
	} else {
		o.qty = qty.amount
	}
	p.skipsp()
	if o.what = p.acceptProduct(); o.what == nil {
		o.err = fmt.Errorf("expected product: got %q", p.first(p.expectLine().value, 12))
		return o
	}
	p.skipsp()
	if o.mine = p.acceptLocation(); o.mine == nil {
		o.factory = p.acceptProduct()
	}
	p.skipsp()
	if eol := p.acceptEOL(); eol == nil {
		o.err = fmt.Errorf("expected eol: got %q", p.first(p.expectLine().value, 12))
		return o
	}
	return o
}

func (p *Parser) expectBombard() *BombardNode {
	o := &BombardNode{line: p.line}
	p.skipsp()
	if id := p.acceptID(); id == nil {
		o.err = fmt.Errorf("expected id: got %q", p.first(p.expectLine().value, 12))
		return o
	} else {
		o.id = id.id
	}
	p.skipsp()
	if id := p.acceptID(); id == nil {
		o.err = fmt.Errorf("expected targetId: got %q", p.first(p.expectLine().value, 12))
		return o
	} else {
		o.targetId = id.id
	}
	p.skipsp()
	if pct := p.acceptPercentage(); pct == nil {
		o.err = fmt.Errorf("expected pctCommitted: got %q", p.first(p.expectLine().value, 12))
		return o
	} else {
		o.pctCommitted = pct.value
	}
	p.skipsp()
	if eol := p.acceptEOL(); eol == nil {
		o.err = fmt.Errorf("expected eol: got %q", p.first(p.expectLine().value, 12))
		return o
	}
	return o
}

func (p *Parser) expectInvade() *InvadeNode {
	o := &InvadeNode{line: p.line}
	p.skipsp()
	if id := p.acceptID(); id == nil {
		o.err = fmt.Errorf("expected id: got %q", p.first(p.expectLine().value, 12))
		return o
	} else {
		o.id = id.id
	}
	p.skipsp()
	if id := p.acceptID(); id == nil {
		o.err = fmt.Errorf("expected targetId: got %q", p.first(p.expectLine().value, 12))
		return o
	} else {
		o.targetId = id.id
	}
	p.skipsp()
	if pct := p.acceptPercentage(); pct == nil {
		o.err = fmt.Errorf("expected pctCommitted: got %q", p.first(p.expectLine().value, 12))
		return o
	} else {
		o.pctCommitted = pct.value
	}
	p.skipsp()
	if eol := p.acceptEOL(); eol == nil {
		o.err = fmt.Errorf("expected eol: got %q", p.first(p.expectLine().value, 12))
		return o
	}
	return o
}

func (p *Parser) expectLine() *LineNode {
	if p.iseof() {
		return &LineNode{line: p.line, value: "**eof**"}
	}
	n := &LineNode{line: p.line}
	var value []byte
	p.skipsp()
	for pc := byte(0); !p.iseof(); {
		ch := p.nextch()
		if ch == '\n' {
			break
		} else if p.isspace(ch) && p.isspace(pc) {
			// compress multiple spaces
		} else {
			value = append(value, ch)
		}
		pc = ch
	}
	n.value = string(value)
	return n
}

func (p *Parser) expectRaid() *RaidNode {
	o := &RaidNode{line: p.line}
	p.skipsp()
	if id := p.acceptID(); id == nil {
		o.err = fmt.Errorf("expected id: got %q", p.first(p.expectLine().value, 12))
		return o
	} else {
		o.id = id.id
	}
	p.skipsp()
	if id := p.acceptID(); id == nil {
		o.err = fmt.Errorf("expected targetId: got %q", p.first(p.expectLine().value, 12))
		return o
	} else {
		o.targetId = id.id
	}
	p.skipsp()
	if pct := p.acceptPercentage(); pct == nil {
		o.err = fmt.Errorf("expected pctCommitted: got %q", p.first(p.expectLine().value, 12))
		return o
	} else {
		o.pctCommitted = pct.value
	}
	p.skipsp()
	if o.product = p.acceptProduct(); o.product == nil {
		if o.profession = p.acceptProfession(); o.profession == nil {
			if o.resource = p.acceptResource(); o.resource == nil {
				o.err = fmt.Errorf("expected material: got %q", p.first(p.expectLine().value, 12))
				return o
			}
		}
	}
	if eol := p.acceptEOL(); eol == nil {
		o.err = fmt.Errorf("expected eol: got %q", p.first(p.expectLine().value, 12))
		return o
	}
	return o
}

func (p *Parser) expectSupport() *SupportNode {
	o := &SupportNode{line: p.line}
	p.skipsp()
	if id := p.acceptID(); id == nil {
		o.err = fmt.Errorf("expected id: got %q", p.first(p.expectLine().value, 12))
		return o
	} else {
		o.id = id.id
	}
	p.skipsp()
	if id := p.acceptID(); id == nil {
		o.err = fmt.Errorf("expected supportId: got %q", p.first(p.expectLine().value, 12))
		return o
	} else {
		o.supportId = id.id
	}
	p.skipsp()
	if id := p.acceptID(); id != nil {
		o.targetId = id.id
	}
	p.skipsp()
	if pct := p.acceptPercentage(); pct == nil {
		o.err = fmt.Errorf("expected pctCommitted: got %q", p.first(p.expectLine().value, 12))
		return o
	} else {
		o.pctCommitted = pct.value
	}
	p.skipsp()
	if eol := p.acceptEOL(); eol == nil {
		o.err = fmt.Errorf("expected eol: got %q", p.first(p.expectLine().value, 12))
		return o
	}
	return o
}
