// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package tokenizer

import "fmt"

// Token is a token from the input.
type Token struct {
	Line    int
	Kind    Kind
	Float   float64
	Integer int
	Text    string
}

// String implements the Stringer interface.
func (t Token) String() string {
	switch t.Kind {
	case COMMA:
		return fmt.Sprintf("{%d ','}", t.Line)
	case COMMENT:
		return fmt.Sprintf("{%d ;...", t.Line)
	case DEPOSITID, FACTGRP, MINEGRP, POPULATION, PRODUCT, RESEARCH, RESOURCE:
		return fmt.Sprintf("{%d %s}", t.Line, t.Text)
	case EOF:
		return fmt.Sprintf("{%d $$}", t.Line)
	case EOL:
		return fmt.Sprintf("{%d '\\n'}", t.Line)
	case FLOAT:
		return fmt.Sprintf("{%d %g}", t.Line, t.Float)
	case INTEGER:
		return fmt.Sprintf("{%d %d}", t.Line, t.Integer)
	case PARENCL:
		return fmt.Sprintf("{%d ')'}", t.Line)
	case PARENOP:
		return fmt.Sprintf("{%d '('}", t.Line)
	case PERCENTAGE:
		return fmt.Sprintf("{%d %d%%}", t.Line, t.Integer)
	case QTEXT:
		return fmt.Sprintf("{%d `%s`}", t.Line, t.Text)
	case SPACES:
		return fmt.Sprintf("{%d ...", t.Line)
	case TEXT:
		return fmt.Sprintf("{%d %q}", t.Line, t.Text)
	}
	return fmt.Sprintf("{%d %s %q}", t.Line, t.Kind, t.Text)
}

// Kind is the type of token.
type Kind int

// enums for Kind
const (
	EOF Kind = iota
	EOL
	COMMA
	COMMENT
	DEPOSITID
	FACTGRP
	FLOAT
	INTEGER
	MINEGRP
	PARENCL
	PARENOP
	PERCENTAGE
	POPULATION
	PRODUCT
	QTEXT
	RESEARCH
	RESOURCE
	SPACES
	TEXT
)

// String implements the Stringer interface.
func (k Kind) String() string {
	switch k {
	case EOF:
		return "EOF"
	case EOL:
		return "EOL"
	case COMMA:
		return "COMMA"
	case COMMENT:
		return "COMMENT"
	case DEPOSITID:
		return "DEPOSITID"
	case FACTGRP:
		return "FACTGRP"
	case FLOAT:
		return "FLOAT"
	case INTEGER:
		return "INTEGER"
	case MINEGRP:
		return "MINEGRP"
	case PARENCL:
		return "PARENCL"
	case PARENOP:
		return "PARENOP"
	case PERCENTAGE:
		return "PERCENTAGE"
	case POPULATION:
		return "POPULATION"
	case PRODUCT:
		return "PRODUCT"
	case QTEXT:
		return "QTEXT"
	case RESEARCH:
		return "RESEARCH"
	case RESOURCE:
		return "RESOURCE"
	case SPACES:
		return "SPACES"
	case TEXT:
		return "TEXT"
	}
	return fmt.Sprintf("token(%d)", k)
}
