// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package orders implements a parser for an orders file.
package orders

import (
	"github.com/mdhender/wraithh/internal/parser"
	"github.com/mdhender/wraithh/internal/tokenizer"
)

func Parse(tokens []*tokenizer.Token) (parseTree *parser.Tree, debugTree *parser.DebugTree, err error) {
	b := parser.NewBuilder(tokens)
	if ok := orders(b); !ok {
		return nil, b.DebugTree(), b.Error()
	}
	return b.ParseTree(), b.DebugTree(), nil
}

func orders(b *parser.Builder) (ok bool) {
	b.Enter("orders")
	defer b.Exit(&ok)

	for order(b) || b.Match(tokenizer.Token{Kind: tokenizer.EOL}) {
		//
	}

	return b.Match(tokenizer.Token{Kind: tokenizer.EOF})
}

func order(b *parser.Builder) (ok bool) {
	b.Enter("order")
	defer b.Exit(&ok)

	if assemble(b) {
		return true
	} else if bombard(b) {
		return true
	} else if invade(b) {
		return true
	} else if raid(b) {
		return true
	} else if setup(b) {
		return true
	} else if support(b) {
		return true
	}
	return false
}

func assemble(b *parser.Builder) (ok bool) {
	b.Enter("assemble")
	defer b.Exit(&ok)

	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT, Value: "assemble"}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT}) {
		return false
	}
	if b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) || b.Match(tokenizer.Token{Kind: tokenizer.TEXT}) {
		// optional
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
	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT}) {
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
	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT}) {
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

func xfer_detail(b *parser.Builder) (ok bool) {
	b.Enter("xfer_detail")
	defer b.Exit(&ok)

	if !b.Match(tokenizer.Token{Kind: tokenizer.INTEGER}) {
		return false
	}
	if !b.Match(tokenizer.Token{Kind: tokenizer.TEXT}) {
		return false
	}
	return b.Match(tokenizer.Token{Kind: tokenizer.EOL})
}
