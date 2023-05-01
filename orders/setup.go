// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

import (
	"fmt"
	"strconv"
	"strings"
)

func parseSetup(toks tokens) (Order, tokens) {
	t, rest := toks.next()
	if t.Kind != "text" || strings.ToLower(toks[0].Text) != "setup" {
		return nil, toks
	}
	o := Setup{line: toks[0].Line}

	if t, rest = rest.next(); t.Kind != "text" {
		o.AddError(fmt.Errorf("location: expected text, got %q", t.Text))
	} else {
		o.location = t.Text
	}

	if t, rest = rest.next(); t.Kind != "text" {
		o.AddError(fmt.Errorf("kind: expected 'ship' or 'colony', got %q", t.Text))
	} else if strings.ToLower(t.Text) == "colony" {
		o.kind = "colony"
	} else if strings.ToLower(t.Text) == "ship" {
		o.kind = "ship"
	} else {
		o.AddError(fmt.Errorf("kind: expected 'ship' or 'colony', got %q", t.Text))
	}

	if t, rest = rest.next(); t.Kind != "number" {
		o.errors = append(o.errors, fmt.Errorf("id: expected number, got %q", t.Text))
	} else {
		o.id = t.Number
	}

	if t, rest = rest.next(); t.Kind != "text" {
		o.AddError(fmt.Errorf("action: expected 'transfer', got %q", t.Text))
	} else if strings.ToLower(t.Text) != "transfer" {
		o.action = "transfer"
	} else {
		o.AddError(fmt.Errorf("action: expected 'transfer', got %q", t.Text))
	}

	for id := 1; len(rest) != 0 && rest[0].Text != "\n"; id++ {
		t, rest = rest.next()
		if strings.ToLower(t.Text) == "end" {
			break
		}
		foo := strings.Fields(t.Text)
		if len(foo) != 2 {
			o.AddError(fmt.Errorf("item: expected qty name (tech-level), got %q", t.Text))
			continue
		}
		var item transferItem
		var err error
		if item.qty, err = strconv.Atoi(foo[0]); err != nil {
			o.AddError(fmt.Errorf("item: expected qty name (tech-level), got %q", t.Text))
			continue
		}
		if dtls := strings.Split(foo[1], "-"); len(dtls) == 1 {
			item.item = dtls[0]
		} else {
			item.item = strings.Join(dtls[0:len(dtls)-1], "-")
			if item.techLevel, err = strconv.Atoi(dtls[len(dtls)-1]); err != nil {
				o.AddError(fmt.Errorf("item: expected qty name (tech-level), got %q", t.Text))
				continue
			}
		}
	}
	if len(rest) == 0 || rest[0].Text != "\n" {
		o.AddError(fmt.Errorf("missing 'end'"))
	}

	// consume extra arguments, if any
	for len(rest) != 0 && rest[0].Kind != "eol" {
		o.errors = append(o.errors, fmt.Errorf("expected eol: got %q", rest[0].Text))
		rest = rest[1:]
	}

	return &o, nil
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

func (s *Setup) Id() int            { return s.id }
func (s *Setup) AddError(err error) { s.errors = append(s.errors, err) }
func (s *Setup) Errors() []error    { return s.errors }
func (s *Setup) Execute() error     { panic("!") }
func (s *Setup) Line() int          { return s.line }
func (s *Setup) String() string     { return fmt.Sprintf("setup ... end") }
