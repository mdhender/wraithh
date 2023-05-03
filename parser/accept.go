// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package parser

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
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

func (p *Parser) acceptComma() *CommaNode {
	if p.peekch() != ',' {
		return nil
	}
	_ = p.nextch()
	return &CommaNode{line: p.line}
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

func (p *Parser) acceptEOF() *EOFNode {
	if !p.iseof() {
		return nil
	}
	return &EOFNode{line: p.line}
}

func (p *Parser) acceptEOL() *EOLNode {
	if p.peekch() != '\n' {
		return nil
	}
	_ = p.nextch()
	return &EOLNode{line: p.line - 1}
}

func (p *Parser) acceptID() *IDNode {
	if n := p.acceptInteger(); n != nil {
		return &IDNode{line: p.line, id: n.value}
	}
	return nil
}

func (p *Parser) acceptInteger() *IntegerNode {
	saved := p.clone()
	var t []byte
	for p.isdigit(p.peekch()) {
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
	n, err := strconv.Atoi(string(t))
	if err != nil { // this can't happen
		panic("assert(err != nil)")
	}
	return &IntegerNode{line: p.line, value: n}
}

func (p *Parser) acceptLocation() *LocationNode {
	panic("!")
}

func (p *Parser) acceptName() *NameNode {
	saved := p.clone()
	if !p.isalpha(p.peekch()) {
		p.restore(saved)
		return nil
	}
	t := []byte{p.nextch()}
	for pc := byte(0); p.isalnum(p.peekch()) || p.isspace(p.peekch()); {
		ch := p.nextch()
		// collapse multiple spaces in a name to just a single space
		runOfSpaces := p.isspace(ch) && p.isspace(pc)
		if !runOfSpaces {
			t = append(t, p.nextch())
		}
		pc = ch
	}
	if ch := p.peekch(); !(p.iseof() || ch == '\n' || p.isspace(ch)) {
		p.restore(saved)
		return nil
	}
	t = bytes.TrimSpace(t)
	return &NameNode{line: p.line, value: string(t)}
}

func (p *Parser) acceptPercentage() *PercentageNode {
	saved := p.clone()
	var t []byte
	for p.isdigit(p.peekch()) {
		t = append(t, p.nextch())
	}
	if len(t) == 0 {
		p.restore(saved)
		return nil
	} else if p.nextch() != '%' {
		p.restore(saved)
		return nil
	}
	if ch := p.peekch(); !(p.iseof() || ch == '\n' || p.isspace(ch)) {
		p.restore(saved)
		return nil
	}
	n, err := strconv.Atoi(string(t))
	if err != nil { // this can't happen
		panic("assert(err != nil)")
	}
	return &PercentageNode{line: p.line, value: n}
}

func (p *Parser) acceptProduct() *ProductNode {
	saved := p.clone()
	if !p.isalpha(p.peekch()) {
		p.restore(saved)
		return nil
	}
	var t []byte
	for p.isalnum(p.peekch()) || p.peekch() == '-' {
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
	// product will be xxx, xxx-yyy, or xxx-yyy-tl
	// if product includes tl, we must extract it.
	product := strings.TrimSpace(strings.ToLower(string(t)))
	var techLevel int
	if fields := strings.Split(product, "-"); len(fields) == 1 {
		// product is xxx
		techLevel = 0
	} else {
		// product is xxx-yyy-tl or xxx-yyy
		firstFields, lastField := fields[:len(fields)-1], fields[len(fields)-1]
		if tl, err := strconv.Atoi(lastField); err != nil {
			techLevel = 0
		} else {
			product, techLevel = strings.Join(firstFields, "-"), tl
		}
	}
	switch product {
	case "anti-missile":
		return &ProductNode{line: p.line, _type: ANTIMISSILE, techLevel: techLevel}
	case "assault-craft":
		return &ProductNode{line: p.line, _type: ASSAULTCRAFT, techLevel: techLevel}
	case "assault-weapons":
		return &ProductNode{line: p.line, _type: ASSAULTWEAPONS, techLevel: techLevel}
	case "automation":
		return &ProductNode{line: p.line, _type: AUTOMATION, techLevel: techLevel}
	case "consumer-goods":
		return &ProductNode{line: p.line, _type: CONSUMERGOODS, techLevel: techLevel}
	case "energy-shield":
		return &ProductNode{line: p.line, _type: ENERGYSHIELD, techLevel: techLevel}
	case "energy-weapon":
		return &ProductNode{line: p.line, _type: ENERGYWEAPON, techLevel: techLevel}
	case "factory":
		return &ProductNode{line: p.line, _type: FACTORY, techLevel: techLevel}
	case "farm":
		return &ProductNode{line: p.line, _type: FARM, techLevel: techLevel}
	case "food":
		return &ProductNode{line: p.line, _type: FOOD, techLevel: techLevel}
	case "hyper-engine":
		return &ProductNode{line: p.line, _type: HYPERENGINE, techLevel: techLevel}
	case "life-support":
		return &ProductNode{line: p.line, _type: LIFESUPPORT, techLevel: techLevel}
	case "light-structural-unit":
		return &ProductNode{line: p.line, _type: LIGHTSTRUCTURAL, techLevel: techLevel}
	case "military-robot":
		return &ProductNode{line: p.line, _type: MILITARYROBOT, techLevel: techLevel}
	case "military-supplies":
		return &ProductNode{line: p.line, _type: MILITARYSUPPLIES, techLevel: techLevel}
	case "mine":
		return &ProductNode{line: p.line, _type: MINE, techLevel: techLevel}
	case "missile":
		return &ProductNode{line: p.line, _type: MISSILE, techLevel: techLevel}
	case "sensor":
		return &ProductNode{line: p.line, _type: SENSOR, techLevel: techLevel}
	case "space-drive":
		return &ProductNode{line: p.line, _type: SPACEDRIVE, techLevel: techLevel}
	case "structural-unit":
		return &ProductNode{line: p.line, _type: STRUCTURALUNIT, techLevel: techLevel}
	case "super-light-structural-unit":
		return &ProductNode{line: p.line, _type: SUPERLIGHTSTRUCTURAL, techLevel: techLevel}
	case "transport":
		return &ProductNode{line: p.line, _type: TRANSPORT, techLevel: techLevel}
	}
	p.restore(saved)
	return nil
}

func (p *Parser) acceptProfession() *ProfessionNode {
	saved := p.clone()
	if !p.isalpha(p.peekch()) {
		p.restore(saved)
		return nil
	}
	var t []byte
	for p.isalpha(p.peekch()) || p.peekch() == '-' {
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
	switch strings.ToLower(string(t)) {
	case "civilian":
		return &ProfessionNode{line: p.line, _type: CIVILIAN}
	case "construction-crew":
		return &ProfessionNode{line: p.line, _type: CONSTRUCTIONCREW}
	case "professional":
		return &ProfessionNode{line: p.line, _type: PROFESSIONAL}
	case "soldier":
		return &ProfessionNode{line: p.line, _type: SOLDIER}
	case "spy":
		return &ProfessionNode{line: p.line, _type: SPY}
	case "unskilled-worker":
		return &ProfessionNode{line: p.line, _type: UNSKILLEDWORKER}
	}
	p.restore(saved)
	return nil
}

func (p *Parser) acceptQuantity() *QuantityNode {
	if qty := p.acceptInteger(); qty != nil {
		return &QuantityNode{line: p.line, amount: qty.value}
	}
	return nil
}

func (p *Parser) acceptResource() *ResourceNode {
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
	switch strings.ToLower(string(t)) {
	case "fuel":
		return &ResourceNode{line: p.line, _type: FUEL}
	case "gold":
		return &ResourceNode{line: p.line, _type: GOLD}
	case "metallics":
		return &ResourceNode{line: p.line, _type: METALLICS}
	case "non-metallics":
		return &ResourceNode{line: p.line, _type: NONMETALLICS}
	}
	p.restore(saved)
	return nil
}
