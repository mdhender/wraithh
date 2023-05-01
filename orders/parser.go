// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

import (
	"os"
)

func ParseFile(name string) (Orders, error) {
	buf, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return Parse(buf)
}

func Parse(buf []byte) (Orders, error) {
	toks, err := scanAll(buf)
	if err != nil {
		return nil, err
	}
	//for _, tok := range toks {
	//	fmt.Printf("%3d: %-12s %q\n", tok.Line, tok.Kind, tok.Text)
	//}

	parsers := []func(tokens) (Order, tokens){
		// combat commands
		parseBombard, parseInvade, parseRaid,
		// support commands
		parseSupport,
		// setup commands
		parseSetup,
		// unknown commands
		parseUnknownOrder,
	}

	var ords Orders
	for len(toks) != 0 {
		var o Order
		for _, parser := range parsers {
			if o, toks = parser(toks); o != nil {
				break
			}
		}
		ords = append(ords, o)

		// consume all tokens up to and including end of line
		for len(toks) != 0 {
			foundEol := toks[0].Kind == "eol"
			toks = toks[1:]
			if foundEol {
				break
			}
		}
	}
	return ords, nil
}
