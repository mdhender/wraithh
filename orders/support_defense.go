// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

import "fmt"

func parseSupportDefend(toks tokens) (*SupportDefend, tokens) {
	o := SupportDefend{line: toks[0].Line}

	t, rest := toks.next()
	if t.Kind != "number" {
		o.AddError("id: expected number, got %q", t.Text)
	} else {
		o.id = t.Number
		if t, rest = rest.next(); t.Kind != "number" {
			o.AddError("supportId: expected number, got %q", t.Text)
		} else {
			o.supportId = t.Number
			if t, rest = rest.next(); t.Kind != "percent" {
				o.AddError("pctCommitted: expected percentage, got %q", t.Text)
			} else {
				o.pctCommited = t.Number
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

type SupportDefend struct {
	line        int
	id          int // id of unit being ordered
	supportId   int // id of unit being supported
	pctCommited int
	errors      []error
}

func (o *SupportDefend) Id() int { return o.id }
func (o *SupportDefend) AddError(format string, args ...any) {
	o.errors = append(o.errors, fmt.Errorf(format, args...))
}
func (o *SupportDefend) Errors() []error { return o.errors }
func (o *SupportDefend) Execute() error  { panic("!") }
func (o *SupportDefend) Line() int       { return o.line }
func (o *SupportDefend) String() string {
	return fmt.Sprintf("support-defend %d %d %d%%", o.id, o.supportId, o.pctCommited)
}
