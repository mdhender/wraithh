// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package parser

import "bytes"

// isspace returns true if the character is a space.
func isspace(ch byte) bool {
	switch ch {
	case ' ', '\t', '\f', '\r', '\v':
		return true
	}
	return false
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
