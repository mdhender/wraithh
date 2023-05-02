// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

import "fmt"

func parseRaid(toks tokens) (*Raid, tokens) {
	o := Raid{line: toks[0].Line}

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
				o.pctCommited = t.Number
				if t, rest = rest.next(); t.Kind != "text" {
					o.AddError("material: expected text, got %q", t.Text)
				} else {
					o.material = t.Text
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

type Raid struct {
	line        int
	id          int // id of unit being ordered
	targetId    int // id of unit being attacked
	pctCommited int
	material    string // material to raid
	errors      []error
}

func (o *Raid) Id() int { return o.id }
func (o *Raid) AddError(format string, args ...any) {
	o.errors = append(o.errors, fmt.Errorf(format, args...))
}
func (o *Raid) Errors() []error { return o.errors }
func (o *Raid) Execute() error  { panic("!") }
func (o *Raid) Line() int       { return o.line }
func (o *Raid) String() string {
	return fmt.Sprintf("raid %d %d %d%%", o.id, o.targetId, o.pctCommited)
}
