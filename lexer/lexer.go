// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package lexer

import "bytes"

// Next returns the next field along with the remainder of the input.
// it makes two assumptions:
//  1. the input is always at the start of a field
//  2. if the last field of a record is empty, return it instead of eol
//
// returns nil on end of input.
func Next(input []byte) (string, []byte) {
	if len(input) == 0 { // end of input
		return "", nil
	} else if input[0] == '\n' { // end of record
		return "\n", input[1:]
	} else if input[0] == ',' { // end of field
		return "", input[1:]
	}
	field, length, inQuote := input, 0, false
	for len(input) != 0 && input[0] != '\n' {
		if input[0] == '"' {
			inQuote = !inQuote
		} else if input[0] == ',' && !inQuote {
			input = input[1:]
			break
		}
		length, input = length+1, input[1:]
	}
	field = field[:length]
	// trim the field before returning it and the remainder of the input
	return string(bytes.TrimSpace(field)), input
}
