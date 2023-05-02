// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package lexer_test

import (
	"github.com/mdhender/wraithh/lexer"
	"testing"
)

func TestLexer(t *testing.T) {
	var input []byte
	field, rest := lexer.Next(input)
	if field != "" {
		t.Errorf("nil input: field want \"\", got %q\n", string(field))
	}
	if rest != nil {
		t.Errorf("nil input: rest want nil, got %q\n", string(rest))
	}

	rest = []byte("one, two , \"three\" ,\n four, \n")
	for _, want := range []string{"one", "two", "\"three\"", "\n", "four", "", "\n", ""} {
		field, rest = lexer.Next(rest)
		if want != field {
			t.Errorf("field want %q, got %q\n", want, field)
		}
	}
	if rest != nil {
		t.Errorf("rest  want nil, got %q\n", string(rest))
	}
}
