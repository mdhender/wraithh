// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

import "fmt"

func parseRaid(toks tokens) (Order, tokens) {
	t, rest := toks.next()
	if t.Kind != "number" {
		return nil, toks
	}
	o := Raid{line: t.Line, id: t.Number}

	if t, rest = rest.next(); t.Text != "raid" {
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

	if t, rest = rest.next(); t.Kind != "text" {
		o.errors = append(o.errors, fmt.Errorf("material: expected text, got %q", t.Text))
	} else {
		o.material = t.Text
	}

	// consume extra arguments, if any
	for len(rest) != 0 && rest[0].Kind != "eol" {
		o.errors = append(o.errors, fmt.Errorf("expected eol: got %q", rest[0].Text))
		rest = rest[1:]
	}

	return &o, rest
}

type Raid struct {
	line        int
	id          int // id of unit being ordered
	targetId    int // id of unit being attacked
	pctCommited int
	material    string
	errors      []error
}

func (r *Raid) Id() int            { return r.id }
func (r *Raid) AddError(err error) { r.errors = append(r.errors, err) }
func (r *Raid) Errors() []error    { return r.errors }
func (r *Raid) Execute() error     { panic("!") }
func (r *Raid) Line() int          { return r.line }
func (r *Raid) String() string {
	return fmt.Sprintf("raid %d %d %d%% %q", r.id, r.targetId, r.pctCommited, r.material)
}
