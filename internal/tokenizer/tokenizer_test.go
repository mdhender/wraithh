// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package tokenizer_test

import (
	"github.com/mdhender/wraithh/internal/tokenizer"
	"testing"
)

func TestTokenizer(t *testing.T) {
	input := []byte("Hello\r\n世界\t55 66%")
	for _, tc := range []struct {
		expect string
		kind   tokenizer.Kind
	}{
		{"Hello", tokenizer.TEXT},
		{"\r", tokenizer.SPACES},
		{"\n", tokenizer.EOL},
		{"世界", tokenizer.TEXT},
		{"\t", tokenizer.SPACES},
		{"55", tokenizer.INTEGER},
		{" ", tokenizer.SPACES},
		{"66%", tokenizer.PERCENTAGE},
		{kind: tokenizer.EOF},
	} {
		var got []byte
		var kind tokenizer.Kind
		kind, got, input = tokenizer.Next(input)
		if string(got) != tc.expect {
			t.Errorf("want %q, got %q\n", tc.expect, string(got))
		}
		if kind != tc.kind {
			t.Errorf("want %d, got %d\n", tc.kind, kind)
		}
	}
	if len(input) != 0 {
		t.Errorf("want eos, got %q\n", string(input))
	}
}
