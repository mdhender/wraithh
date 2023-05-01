// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

import "fmt"

func parseInvade(toks tokens) (Order, tokens) {
	t, rest := toks.next()
	if t.Kind != "number" {
		return nil, toks
	}
	o := Invade{line: t.Line, id: t.Number}

	if t, rest = rest.next(); t.Text != "invade" {
		return nil, toks
	}

	if t, rest = rest.next(); t.Kind != "number" {
		o.errors = append(o.errors, fmt.Errorf("targetId: expected number, got %q", t.Text))
	} else {
		o.targetId = t.Number
	}

	if t, rest = rest.next(); t.Kind != "percent" {
		o.errors = append(o.errors, fmt.Errorf("pctCommitted: expected percentage, got %q", t.Text))
	} else {
		o.pctCommited = t.Number
	}

	// consume extra arguments, if any
	for len(rest) != 0 && rest[0].Kind != "eol" {
		o.errors = append(o.errors, fmt.Errorf("expected eol: got %q", rest[0].Text))
		rest = rest[1:]
	}

	return &o, rest
}

type Invade struct {
	line        int
	id          int // id of unit being ordered
	targetId    int // id of unit being attacked
	pctCommited int
	errors      []error
}

func (i *Invade) Id() int            { return i.id }
func (i *Invade) AddError(err error) { i.errors = append(i.errors, err) }
func (i *Invade) Errors() []error    { return i.errors }
func (i *Invade) Execute() error     { panic("!") }
func (i *Invade) Line() int          { return i.line }
func (i *Invade) String() string {
	return fmt.Sprintf("invade %d %d %d%%", i.id, i.targetId, i.pctCommited)
}
