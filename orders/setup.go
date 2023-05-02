// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

import (
	"fmt"
	"strconv"
	"strings"
)

func parseSetup(toks tokens) (*Setup, tokens) {
	o := Setup{line: toks[0].Line}

	t, rest := toks.next()
	if t.Kind != "text" {
		o.AddError("location: expected text, got %q", t.Text)
	} else {
		o.location = t.Text
		if t, rest = rest.next(); t.Kind != "text" {
			o.AddError("kind: expected 'ship' or 'colony', got %q", t.Text)
		} else {
			if strings.ToLower(t.Text) == "colony" {
				o.kind = "colony"
			} else if strings.ToLower(t.Text) == "ship" {
				o.kind = "ship"
			} else {
				o.AddError("kind: expected 'ship' or 'colony', got %q", t.Text)
			}

			if t, rest = rest.next(); t.Kind != "number" {
				o.errors = append(o.errors, fmt.Errorf("id: expected number, got %q", t.Text))
			} else {
				o.id = t.Number

				if t, rest = rest.next(); t.Kind != "text" {
					o.AddError("action: expected 'transfer', got %q", t.Text)
				} else if strings.ToLower(t.Text) != "transfer" {
					o.AddError("action: expected 'transfer', got %q", t.Text)
				} else {
					o.action = "transfer"

					for len(rest) != 0 && strings.ToLower(rest[0].Text) != "end" {
						t, rest = rest.next()
						if t.Text == "," || t.Kind == "eol" {
							continue
						}
						fields := strings.Fields(t.Text)
						qty, kind := "", ""
						switch len(fields) {
						case 0, 1:
							o.AddError("item: expected qty name (tech-level), got %q", t.Text)
							continue
						case 2:
							qty, kind = fields[0], fields[1]
						default:
							qty, kind = fields[0], strings.Join(fields[1:], " ")
						}
						var item transferItem
						var err error
						if item.qty, err = strconv.Atoi(qty); err != nil {
							o.AddError("item: expected qty name (tech-level), got %q", t.Text)
							continue
						}
						dtls := strings.Split(kind, "-")
						if len(dtls) == 1 {
							item.item = dtls[0]
						} else {
							item.item = strings.Join(dtls[0:len(dtls)-1], "-")
							if item.techLevel, err = strconv.Atoi(dtls[len(dtls)-1]); err != nil {
								o.AddError("item: expected qty name (tech-level), got %q", t.Text)
								continue
							}
						}
						o.items = append(o.items, item)
					}
				}
			}
		}
	}

	// setup should end with an end command
	if len(rest) == 0 {
		o.AddError("expected end: got eof")
	} else if strings.ToLower(rest[0].Text) != "end" {
		o.AddError("expected end: got %q", rest[0].Text)
	} else {
		rest = rest[1:]
	}

	// consume extra arguments, if any
	for foundEol := false; len(rest) != 0 && !foundEol; rest = rest[1:] {
		if foundEol = rest[0].Kind == "eol"; !foundEol {
			o.AddError("expected eol: got %q", rest[0].Text)
		}
	}

	return &o, rest
}

type Setup struct {
	line        int
	id          int    // id of unit establishing ship or colony
	action      string // must be 'transfer'
	location    string // location being set up
	kind        string
	targetId    int // id of unit being attacked
	pctCommited int
	items       []transferItem
	errors      []error
}

type transferItem struct {
	item      string // name
	qty       int
	techLevel int // optional tech level
}

func (o *Setup) Id() int { return o.id }
func (o *Setup) AddError(format string, args ...any) {
	o.errors = append(o.errors, fmt.Errorf(format, args...))
}
func (o *Setup) Errors() []error { return o.errors }
func (o *Setup) Execute() error  { panic("!") }
func (o *Setup) Line() int       { return o.line }
func (o *Setup) String() string {
	s := fmt.Sprintf("setup %d loc %q kind %q %q", o.id, o.location, o.kind, o.action)
	for _, item := range o.items {
		if item.techLevel == 0 {
			s += fmt.Sprintf(" (%q %d)", item.item, item.qty)
		} else {
			s += fmt.Sprintf(" (%q %d)", fmt.Sprintf("%s-%d", item.item, item.techLevel), item.qty)
		}
	}
	s += " end"
	return s
}
