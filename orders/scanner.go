// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

import (
	"bytes"
	"encoding/csv"
	"strconv"
	"strings"
)

// scanAll returns all the tokens from a buffer.
// The buffer is a CSV with variable number of fields per record.
// A new-line is included to separate the records.
// If there is no new-line at the end of the buffer, we add one.
func scanAll(buf []byte) (tokens, error) {
	r := csv.NewReader(bytes.NewReader(buf))
	r.FieldsPerRecord, r.ReuseRecord = -1, true

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	var toks []token
	for n, rec := range records {
		line := n + 1
		for _, field := range rec {
			val := strings.TrimSpace(field)
			t := token{Line: line, Kind: "text", Text: val}
			if len(val) == 0 {
				t.Kind = "text"
			} else if val == "\n" {
				t.Kind = "eol"
			} else if n, ok := t.asnum(); ok {
				t.Kind, t.Number = "number", n
			} else if n, ok := t.aspct(); ok {
				t.Kind, t.Number = "percent", n
			}
			toks = append(toks, t)
		}
		toks = append(toks, token{Line: line, Kind: "eol", Text: "\n"})
	}
	if len(toks) == 0 || toks[len(toks)-1].Text != "\n" {
		toks = append(toks, token{Line: len(records) + 1, Kind: "eol", Text: "\n"})
	}

	return toks, nil
}

type tokens []token

func (t tokens) toeol() tokens {
	for len(t) != 0 && t[0].Kind != "eol" {
		t = t[1:]
	}
	return t
}

type token struct {
	Line   int
	Kind   string
	Number int
	Text   string
}

func (t tokens) eof() bool {
	return len(t) == 0
}

func (t tokens) eol() bool {
	return len(t) > 0 && t[0].Text == "\n"
}

func (t tokens) next() (token, tokens) {
	if t.eof() {
		return token{}, nil
	}
	return t[0], t[1:]
}

func (t token) asnum() (int, bool) {
	if n, err := strconv.Atoi(t.Text); err == nil {
		return n, true
	}
	return 0, false
}

func (t token) aspct() (int, bool) {
	if strings.HasSuffix(t.Text, "%") {
		if n, err := strconv.Atoi(t.Text[:len(t.Text)-1]); err == nil {
			return n, true
		}
	}
	return 0, false
}

func (t token) astext() (string, bool) {
	return t.Text, true
}

// Next returns the next token in the buffer along with the remainder of the buffer.
// It skips leading spaces and the token is always trimmed of leading and trailing spaces.
// tokens always end at a comma, end of line, or end of input.
func Next(b []byte) (token []byte, rest []byte) {
	if len(b) != 0 && b[0] == '\n' {
		return b[:1], b[1:]
	} else if len(b) != 0 && b[0] == ',' {
		return b[:1], b[1:]
	}
	// skip leading spaces
	_, b = runof(b, []byte{' ', '\t', '\f', '\r', '\v'})

	if len(b) == 0 { // end of input
		return nil, nil
	} else if b[0] == '\n' { // empty field then end of line
		return []byte{}, b
	} else if b[0] == ',' { // empty field then end of field
		return []byte{}, b
	} else if b[0] == '"' { // quoted text
		token, b = append(token, b[0]), b[1:]
		for len(b) != 0 && b[0] != '\n' {
			ch := b[0]
			token, b = append(token, ch), b[1:]
			if ch == '"' {
				break
			}
		}
		return bytes.TrimSpace(token), b
	}

	// token is all characters up to a comma, end of line, or end of input
	token, rest = runto(b, []byte{',', '\n'})
	return bytes.TrimSpace(token), rest
}

// runof advances through a buffer while the characters are in the accept set.
// it returns the slice accepted and the remainder of the buffer.
func runof(buffer, accept []byte) ([]byte, []byte) {
	length, rest := 0, buffer
	for len(rest) != 0 && bytes.IndexByte(accept, rest[0]) != -1 {
		length, rest = length+1, rest[1:]
	}
	if length == 0 {
		return nil, buffer
	}
	return buffer[:length], rest
}

// runto advances through a buffer until a reject or end of input is found.
// it returns the slice up to the first rejected and the remainder of the buffer.
func runto(buffer, reject []byte) ([]byte, []byte) {
	length, rest := 0, buffer
	for len(rest) != 0 && bytes.IndexByte(reject, rest[0]) == -1 {
		length, rest = length+1, rest[1:]
	}
	if length == 0 {
		return nil, buffer
	}
	return buffer[:length], rest
}
