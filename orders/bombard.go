// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

import "fmt"

func parseBombard(toks tokens) (*Bombard, tokens) {
	o := Bombard{line: toks[0].Line}

	t, rest := toks.next()
	if t.Kind != "number" {
		o.AddError("id: expected number, got %q", t.Text)
	} else {
		o.id = t.Number
		if t, rest = rest.next(); t.Kind != "number" {
			o.AddError("targetId: expected number, got %q", t.Text)
		} else {
			o.targetId = t.Number
			if t, rest = rest.next(); t.Kind != "percent" {
				o.AddError("pctCommitted: expected percentage, got %q", t.Text)
			} else {
				o.pctCommitted = t.Number
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

type Bombard struct {
	line         int
	id           int // id of unit being ordered
	targetId     int // id of unit being attacked
	pctCommitted int
	errors       []error
}

func (o *Bombard) Id() int { return o.id }
func (o *Bombard) AddError(format string, args ...any) {
	o.errors = append(o.errors, fmt.Errorf(format, args...))
}
func (o *Bombard) Errors() []error { return o.errors }
func (o *Bombard) Execute() error  { panic("!") }
func (o *Bombard) Line() int       { return o.line }
func (o *Bombard) String() string {
	return fmt.Sprintf("bombard %d %d %d%%", o.id, o.targetId, o.pctCommitted)
}
