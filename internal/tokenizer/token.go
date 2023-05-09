// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package tokenizer

import "fmt"

// Token is a token from the input.
type Token struct {
	Line  int
	Kind  Kind
	Value string
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
		return "NUMBER"
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
