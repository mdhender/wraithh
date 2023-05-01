// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

import "fmt"

func parseBombard(toks tokens) (Order, tokens) {
	t, rest := toks.next()
	if t.Kind != "number" {
		return nil, toks
	}
	o := Bombard{line: t.Line, id: t.Number}

	if t, rest = rest.next(); t.Text != "bombard" {
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

type Bombard struct {
	line        int
	id          int // id of unit being ordered
	targetId    int // id of unit being attacked
	pctCommited int
	errors      []error
}

func (b *Bombard) Id() int            { return b.id }
func (b *Bombard) AddError(err error) { b.errors = append(b.errors, err) }
func (b *Bombard) Errors() []error    { return b.errors }
func (b *Bombard) Execute() error     { panic("!") }
func (b *Bombard) Line() int          { return b.line }
func (b *Bombard) String() string {
	return fmt.Sprintf("bombard %d %d %d%%", b.id, b.targetId, b.pctCommited)
}
