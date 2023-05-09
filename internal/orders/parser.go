// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package orders implements a parser for an orders file.
package orders

import (
	"fmt"
	"github.com/mdhender/wraithh/internal/parser"
	"github.com/mdhender/wraithh/internal/tokenizer"
	"strconv"
	"strings"
)

var stopOnFirstError = true

func Parse(tokens []*tokenizer.Token, stop bool) (parseTree *parser.Tree, debugTree *parser.DebugTree, err error) {
	stopOnFirstError = stop
	b := parser.NewBuilder(tokens)
	if ok := orders(b); !ok {
		return nil, b.DebugTree(), b.Error()
	}
	return b.ParseTree(), b.DebugTree(), nil
}

func iskeyword(token *tokenizer.Token, kw string) bool {
	if token == nil || token.Kind != tokenizer.TEXT {
		return false
	}
	return kw == strings.ToLower(token.Value)
}

func orders(b *parser.Builder) (ok bool) {
	b.Enter("orders")
	defer b.Exit(&ok)

	if stopOnFirstError {
		for order(b) {
			//
		}
	} else {
		for token, ok := b.Peek(1); ok && token.Kind != tokenizer.EOF; token, ok = b.Peek(1) {
			if order(b) {
				continue
			}
			if b.Match(tokenizer.Token{Kind: tokenizer.EOL}) {
				continue
			}
			// error recovery consumes to end of line and save as an unknown node.
			unknown(b)
		}
	}

	return b.Match(tokenizer.Token{Kind: tokenizer.EOF})
}

func order(b *parser.Builder) (ok bool) {
	b.Enter("order")
	defer b.Exit(&ok)

	token, _ := b.Peek(1)
	if iskeyword(token, "assemble") {
		return assemble(b)
	} else if iskeyword(token, "bombard") {
		return bombard(b)
	} else if iskeyword(token, "buy") {
		return buy(b)
	} else if iskeyword(token, "disassemble") {
		return disassemble(b)
	} else if iskeyword(token, "invade") {
		return invade(b)
	} else if iskeyword(token, "probe") {
		return probe(b)
	} else if iskeyword(token, "raid") {
		return raid(b)
	} else if iskeyword(token, "retool") {
		return retool(b)
	} else if iskeyword(token, "sell") {
		return sell(b)
	} else if iskeyword(token, "setup") {
		return setup(b)
	} else if iskeyword(token, "support") {
		return support(b)
	} else if iskeyword(token, "survey") {
		return survey(b)
	} else if iskeyword(token, "transfer") {
		return transfer(b)
	}
	return false
}

func assemble(b *parser.Builder) (ok bool) {
	b.Enter("assemble")
	defer b.Exit(&ok)
	// command
	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "assemble"}) {
		return false
	}
	// csid
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	// optional deposit id or factory group or mining group
	if b.Match(tokenizer.Token{Kind: tokenizer.DEPOSITID}) || b.Match(tokenizer.Token{Kind: tokenizer.FACTGRP}) || b.Match(tokenizer.Token{Kind: tokenizer.MINEGRP}) {
		// do something
	}
	// quantity
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	// material
	if !material(b) {
		return false
	}
	return b.Match(tokenizer.Token{Kind: tokenizer.EOL})
}

func bombard(b *parser.Builder) (ok bool) {
	b.Enter("bombard")
	defer b.Exit(&ok)

	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "bombard"}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.PERCENTAGE}) {
		return false
	}
	return b.Match(tokenizer.Token{Kind: tokenizer.EOL})
}

func buy(b *parser.Builder) (ok bool) {
	b.Enter("buy")
	defer b.Exit(&ok)
	// command
	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "buy"}) {
		return false
	}
	// csid
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	// item may be research or a product
	if b.Match(tokenizer.Token{Kind: tokenizer.RESEARCH}) {
		// do something
	} else if b.Match(tokenizer.Token{Kind: tokenizer.PRODUCT}) {
		// quantity must be specified for product
		if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
			return false
		}
	} else {
		return false
	}
	// price
	if !number(b) {
		return false
	}
	return b.Match(tokenizer.Token{Kind: tokenizer.EOL})
}

func cargo(b *parser.Builder) (ok bool) {
	b.Enter("cargo")
	defer b.Exit(&ok)
	// population or product or research or resource
	return b.Match(tokenizer.Token{Kind: tokenizer.POPULATION}) ||
		b.Match(tokenizer.Token{Kind: tokenizer.PRODUCT}) ||
		b.Match(tokenizer.Token{Kind: tokenizer.RESEARCH}) ||
		b.Match(tokenizer.Token{Kind: tokenizer.RESOURCE})
}

func coordinate(b *parser.Builder) (ok bool) {
	b.Enter("coordinate")
	defer b.Exit(&ok)

	if !b.Match(tokenizer.Token{Kind: tokenizer.PARENOP}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.COMMA}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.COMMA}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	// optional orbit. if present, must be in the range 1..10
	if b.Match(tokenizer.Token{Kind: tokenizer.COMMA}) {
		orbit, _ := b.Peek(1)
		if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
			return false
		}
		// orbit must be 1..10
		if n, _ := strconv.Atoi(orbit.Value); !(1 <= n && n <= 10) {
			return false
		}
	}

	return b.Match(tokenizer.Token{Kind: tokenizer.PARENCL})
}

func disassemble(b *parser.Builder) (ok bool) {
	b.Enter("disassemble")
	defer b.Exit(&ok)
	// command
	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "disassemble"}) {
		return false
	}
	// csid
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	// optional factory group or mining group
	if b.Match(tokenizer.Token{Kind: tokenizer.FACTGRP}) || b.Match(tokenizer.Token{Kind: tokenizer.MINEGRP}) {
		// do something
	}
	// quantity
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	// material
	if !material(b) {
		return false
	}
	return b.Match(tokenizer.Token{Kind: tokenizer.EOL})
}

func invade(b *parser.Builder) (ok bool) {
	b.Enter("invade")
	defer b.Exit(&ok)

	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "invade"}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.PERCENTAGE}) {
		return false
	}
	return b.Match(tokenizer.Token{Kind: tokenizer.EOL})
}

func material(b *parser.Builder) (ok bool) {
	b.Enter("material")
	defer b.Exit(&ok)
	// product or research
	return b.Match(tokenizer.Token{Kind: tokenizer.PRODUCT}) || b.Match(tokenizer.Token{Kind: tokenizer.RESEARCH})
}

func number(b *parser.Builder) (ok bool) {
	b.Enter("number")
	defer b.Exit(&ok)
	// float or integer
	return b.Match(tokenizer.Token{Kind: tokenizer.FLOAT}) || b.Match(tokenizer.Token{Kind: tokenizer.INTEGER})
}

func probe(b *parser.Builder) (ok bool) {
	b.Enter("probe")
	defer b.Exit(&ok)

	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "probe"}) {
		return false
	}
	// csid
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	// orbit or coordinates
	if b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		// orbit
	} else if coordinate(b) {
		// coordinates
	}
	return b.Match(tokenizer.Token{Kind: tokenizer.EOL})
}

func raid(b *parser.Builder) (ok bool) {
	b.Enter("raid")
	defer b.Exit(&ok)

	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "raid"}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.PERCENTAGE}) {
		return false
	}
	if !cargo(b) {
		return false
	}
	return b.Match(tokenizer.Token{Kind: tokenizer.EOL})
}

func retool(b *parser.Builder) (ok bool) {
	b.Enter("retool")
	defer b.Exit(&ok)
	// command
	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "retool"}) {
		return false
	}
	// csid
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	// factgroup
	if !b.Match(tokenizer.Token{Kind: tokenizer.FACTGRP}) {
		return false
	}
	// material
	if !material(b) {
		return false
	}
	return b.Match(tokenizer.Token{Kind: tokenizer.EOL})
}

func sell(b *parser.Builder) (ok bool) {
	b.Enter("sell")
	defer b.Exit(&ok)
	// command
	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "sell"}) {
		return false
	}
	// csid
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	// item may be research or a product
	if b.Match(tokenizer.Token{Kind: tokenizer.RESEARCH}) {
		// do something
	} else if b.Match(tokenizer.Token{Kind: tokenizer.PRODUCT}) {
		// quantity must be specified for product
		if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
			return false
		}
	} else {
		return false
	}
	// price
	if !number(b) {
		return false
	}
	return b.Match(tokenizer.Token{Kind: tokenizer.EOL})
}

func setup(b *parser.Builder) (ok bool) {
	b.Enter("setup")
	defer b.Exit(&ok)

	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "setup"}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	if !coordinate(b) {
		return false
	}
	if !(b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "ship"}) || b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "colony"})) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "transfer"}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.EOL}) {
		return false
	}
	for xfer_detail(b) {
		p1, _ := b.Peek(1)
		if p1 == nil || p1.Kind == tokenizer.EOF {
			break
		} else if p1.Kind == tokenizer.TEXT && p1.Value == "end" {
			break
		}
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "end"}) {
		return false
	}
	return b.Match(tokenizer.Token{Kind: tokenizer.EOL})
}

func support(b *parser.Builder) (ok bool) {
	b.Enter("support")
	defer b.Exit(&ok)

	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "support"}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	if b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		// optional
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.PERCENTAGE}) {
		return false
	}
	return b.Match(tokenizer.Token{Kind: tokenizer.EOL})
}

func survey(b *parser.Builder) (ok bool) {
	b.Enter("survey")
	defer b.Exit(&ok)
	// command
	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "survey"}) {
		return false
	}
	// csid
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	return b.Match(tokenizer.Token{Kind: tokenizer.EOL})
}

func transfer(b *parser.Builder) (ok bool) {
	b.Enter("transfer")
	defer b.Exit(&ok)
	// command
	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "transfer"}) {
		return false
	}
	// csid
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	// quantity
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	// cargo
	if !cargo(b) {
		return false
	}
	// csid
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	return b.Match(tokenizer.Token{Kind: tokenizer.EOL})
}

func unknown(b *parser.Builder) (ok bool) {
	b.Enter("unknown")
	defer b.Exit(&ok)

	for {
		if b.Match(tokenizer.Token{Kind: tokenizer.COMMA}) {
			continue
		}
		if b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
			continue
		}
		if b.Match(tokenizer.Token{Kind: tokenizer.PARENCL}) {
			continue
		}
		if b.Match(tokenizer.Token{Kind: tokenizer.PARENOP}) {
			continue
		}
		if b.Match(tokenizer.Token{Kind: tokenizer.PERCENTAGE}) {
			continue
		}
		if b.Match(tokenizer.Token{Kind: tokenizer.TEXT}) {
			continue
		}
		break
	}
	if b.Match(tokenizer.Token{Kind: tokenizer.EOL}) {
		return true
	}
	if b.Match(tokenizer.Token{Kind: tokenizer.EOF}) {
		return true
	}
	token, _ := b.Peek(1)
	panic(fmt.Sprintf("unknown token %s %q\n", token.Kind, token.Value))
}

func xfer_detail(b *parser.Builder) (ok bool) {
	b.Enter("xfer_detail")
	defer b.Exit(&ok)

	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	if !cargo(b) {
		return false
	}
	return b.Match(tokenizer.Token{Kind: tokenizer.EOL})
}
