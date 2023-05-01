// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

import "fmt"

func parseUnknownOrder(toks tokens) (Order, tokens) {
	o := &Unknown{
		line:   toks[0].Line,
		errors: []error{fmt.Errorf("%d: unknown order", toks[0].Line)},
	}
	for len(toks) != 0 && toks[0].Kind != "eol" {
		toks = toks[1:]
	}
	return o, toks
}

type Unknown struct {
	line   int
	errors []error
}

func (u *Unknown) Id() int            { return 0 }
func (u *Unknown) AddError(err error) { u.errors = append(u.errors, err) }
func (u *Unknown) Errors() []error    { return u.errors }
func (u *Unknown) Execute() error     { panic("!") }
func (u *Unknown) Line() int          { return u.line }
func (u *Unknown) String() string     { return fmt.Sprintf("unknown line %d", u.line) }
