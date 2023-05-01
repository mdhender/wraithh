// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

import "fmt"

func parseSupport(toks tokens) (Order, tokens) {
	t, rest := toks.next()
	if t.Kind != "number" {
		return nil, toks
	}
	sa := SupportAttack{line: t.Line, id: t.Number}

	if t, rest = rest.next(); t.Text != "support" {
		return nil, toks
	}

	if t, rest = rest.next(); t.Kind != "number" {
		sa.errors = append(sa.errors, fmt.Errorf("supportId: expected number, got %q", t.Text))
	} else {
		sa.supportId = t.Number
	}

	t, rest = rest.next()
	if t.Kind == "number" { // support attack
		sa.targetId = t.Number
		if t, rest = rest.next(); t.Kind != "percent" {
			sa.errors = append(sa.errors, fmt.Errorf("pctCommitted: expected percentage, got %q", t.Text))
		} else {
			sa.pctCommited = t.Number
		}

		// consume extra arguments, if any
		for len(rest) != 0 && rest[0].Kind != "eol" {
			sa.errors = append(sa.errors, fmt.Errorf("expected eol: got %q", rest[0].Text))
			rest = rest[1:]
		}

		return &sa, rest
	}

	// support defense
	sd := SupportDefend{line: sa.line, id: sa.id, errors: sa.errors}
	if t.Kind != "percent" {
		sd.errors = append(sa.errors, fmt.Errorf("pctCommitted: expected percentage, got %q", t.Text))
	} else {
		sd.pctCommited = t.Number
	}

	// consume extra arguments, if any
	for len(rest) != 0 && rest[0].Kind != "eol" {
		sd.errors = append(sd.errors, fmt.Errorf("expected eol: got %q", rest[0].Text))
		rest = rest[1:]
	}

	return &sd, rest
}

type SupportAttack struct {
	line        int
	id          int // id of unit being ordered
	supportId   int // id of unit being supported
	targetId    int // id of unit being attacked
	pctCommited int
	errors      []error
}

func (sa *SupportAttack) Id() int            { return sa.id }
func (sa *SupportAttack) AddError(err error) { sa.errors = append(sa.errors, err) }
func (sa *SupportAttack) Errors() []error    { return sa.errors }
func (sa *SupportAttack) Execute() error     { panic("!") }
func (sa *SupportAttack) Line() int          { return sa.line }
func (sa *SupportAttack) String() string {
	return fmt.Sprintf("support-attack %d %d %d %d%%", sa.id, sa.supportId, sa.targetId, sa.pctCommited)
}

type SupportDefend struct {
	line        int
	id          int // id of unit being ordered
	supportId   int // id of unit being supported
	pctCommited int
	errors      []error
}

func (sd *SupportDefend) Id() int            { return sd.id }
func (sd *SupportDefend) AddError(err error) { sd.errors = append(sd.errors, err) }
func (sd *SupportDefend) Errors() []error    { return sd.errors }
func (sd *SupportDefend) Execute() error     { panic("!") }
func (sd *SupportDefend) Line() int          { return sd.line }
func (sd *SupportDefend) String() string {
	return fmt.Sprintf("support-defend %d %d %d%%", sd.id, sd.supportId, sd.pctCommited)
}
