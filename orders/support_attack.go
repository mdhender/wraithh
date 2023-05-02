// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

import "fmt"

func parseSupportAttack(toks tokens) (*SupportAttack, tokens) {
	o := SupportAttack{line: toks[0].Line}

	t, rest := toks.next()
	if t.Kind != "number" {
		o.AddError("id: expected number, got %q", t.Text)
	} else {
		o.id = t.Number
		if t, rest = rest.next(); t.Kind != "number" {
			o.AddError("supportId: expected number, got %q", t.Text)
		} else {
			o.supportId = t.Number
			if t, rest = rest.next(); t.Kind != "number" {
				o.AddError("targetId: expected number, got %q", t.Text)
			} else {
				o.targetId = t.Number
				if t, rest = rest.next(); t.Kind != "percent" {
					o.AddError("pctCommitted: expected percentage, got %q", t.Text)
				} else {
					o.pctCommited = t.Number
				}
			}
		}
	}

	// consume extra arguments, if any
	for foundEol := false; len(rest) != 0 && !foundEol; rest = rest[1:] {
		if foundEol = rest[0].Kind == "eol"; !foundEol {
			o.AddError("expected eol: got %q", rest[0].Text)
		}
	}

	return &o, rest
}

type SupportAttack struct {
	line        int
	id          int // id of unit being ordered
	supportId   int // id of unit being supported
	targetId    int // id of unit being attacked
	pctCommited int
	errors      []error
}

func (o *SupportAttack) Id() int { return o.id }
func (o *SupportAttack) AddError(format string, args ...any) {
	o.errors = append(o.errors, fmt.Errorf(format, args...))
}
func (o *SupportAttack) Errors() []error { return o.errors }
func (o *SupportAttack) Execute() error  { panic("!") }
func (o *SupportAttack) Line() int       { return o.line }
func (o *SupportAttack) String() string {
	return fmt.Sprintf("support-attack %d %d %d %d%%", o.id, o.supportId, o.targetId, o.pctCommited)
}
