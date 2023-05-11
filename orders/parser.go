// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

//func ParseFile(name string) (Orders, error) {
//	buf, err := os.ReadFile(name)
//	if err != nil {
//		return nil, err
//	}
//	return Parse(buf)
//}
//
//func Parse(input []byte) (Orders, error) {
//	var toks tokens
//	line := 1
//	for len(input) != 0 {
//		var field string
//		field, input = lexer.Next(input)
//		tok := token{Line: line, Kind: "text", Text: field}
//		if len(field) == 0 {
//			tok.Kind = "text"
//		} else if field == "\n" {
//			tok.Kind, line = "eol", line+1
//		} else if n, ok := tok.asnum(); ok {
//			tok.Kind, tok.Number = "number", n
//		} else if n, ok := tok.aspct(); ok {
//			tok.Kind, tok.Number = "percent", n
//		}
//		toks = append(toks, tok)
//	}
//	// force token buffer to end with a new-line
//	if len(toks) == 0 || toks[len(toks)-1].Text != "\n" {
//		toks = append(toks, token{Line: line, Kind: "eol", Text: "\n"})
//	}
//
//	var orders Orders
//	var order Order
//
//	for len(toks) != 0 {
//		tok := toks[0]
//		toks = toks[1:]
//		switch tok.Text {
//		case "\n":
//			// ignore empty record
//			continue
//		//case "bombard":
//		//	order, toks = parseBombard(toks)
//		case "invade":
//			order, toks = parseInvade(toks)
//		case "raid":
//			order, toks = parseRaid(toks)
//		case "setup":
//			order, toks = parseSetup(toks)
//		case "support-attack":
//			order, toks = parseSupportAttack(toks)
//		case "support-defend":
//			order, toks = parseSupportDefend(toks)
//		default: // unknown command
//			order, toks = parseUnknownOrder(tok, toks)
//		}
//		orders = append(orders, order)
//	}
//
//	return orders, nil
//}
