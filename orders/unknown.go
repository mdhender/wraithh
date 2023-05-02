// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

import "fmt"

func parseUnknownOrder(tok token, toks tokens) (*Unknown, tokens) {
	o := &Unknown{
		line:    tok.Line,
		command: tok.Text,
		errors:  []error{fmt.Errorf("unknown order %q", tok.Text)},
	}
	// consume extra arguments, if any
	for foundEol := false; len(toks) != 0 && !foundEol; toks = toks[1:] {
		foundEol = toks[0].Kind == "eol"
	}
	return o, toks
}

type Unknown struct {
	line    int
	command string
	errors  []error
}

func (u *Unknown) Id() int { return 0 }
func (u *Unknown) AddError(format string, args ...any) {
	u.errors = append(u.errors, fmt.Errorf(format, args...))
}
func (u *Unknown) Errors() []error { return u.errors }
func (u *Unknown) Execute() error  { panic("!") }
func (u *Unknown) Line() int       { return u.line }
func (u *Unknown) String() string  { return fmt.Sprintf("unknown command %q", u.command) }
