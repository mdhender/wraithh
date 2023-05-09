// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package tokenizer implements a lexer and tokenizer for order files.
// The tokens returned are compatible with the parser builder.
package tokenizer

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Tokens is a helper function to scan an entire input buffer.
// It returns a slice containing all the tokens found.
// The slice returned will always end with an EOF token.
func Tokens(input []byte) (tokens []*Token) {
	line := 1
	for len(input) != 0 {
		token := &Token{Line: line}
		var lexeme []byte
		token.Kind, lexeme, input = Next(input)
		token.Value = string(lexeme)
		tokens = append(tokens, token)
		if token.Kind == EOL {
			line++
		}
	}
	// force an end of line at the end of the input
	if len(tokens) == 0 || tokens[len(tokens)-1].Kind != EOL {
		tokens = append(tokens, &Token{Line: line, Kind: EOL})
	}
	// add an eof token
	tokens = append(tokens, &Token{Line: line, Kind: EOF})
	return tokens
}

// RemoveComments removes comments from the tokens.
func RemoveComments(input []*Token) (tokens []*Token) {
	for _, token := range input {
		if token.Kind == COMMENT {
			continue
		}
		tokens = append(tokens, token)
	}
	return tokens
}

// RemoveEmptyLines removed empty lines from the tokens.
func RemoveEmptyLines(input []*Token) (tokens []*Token) {
	prior := &Token{Kind: EOL}
	for _, token := range input {
		if token.Kind == EOL && prior.Kind == EOL {
			continue
		}
		tokens = append(tokens, token)
		prior = token
	}
	return tokens
}

// RemoveSpaces removes spaces from the tokens.
func RemoveSpaces(input []*Token) (tokens []*Token) {
	for _, token := range input {
		if token.Kind == SPACES {
			continue
		}
		tokens = append(tokens, token)
	}
	return tokens
}

// Next returns the next lexeme from the buffer.
// The lexeme can be a new-line, run of spaces, a comman, an integer, or a percentage.
// If the buffer is empty, returns nil, nil.
// Otherwise, the lexeme and the remainder of the buffer are returned.
func Next(buffer []byte) (kind Kind, lexeme, rest []byte) {
	if len(buffer) == 0 {
		return EOF, nil, nil
	}

	// if it is a delimiter, just return it.
	if buffer[0] == '\n' {
		return EOL, nil, buffer[1:]
	} else if buffer[0] == ',' {
		return COMMA, buffer[:1], buffer[1:]
	} else if buffer[0] == ')' {
		return PARENCL, buffer[:1], buffer[1:]
	} else if buffer[0] == '(' {
		return PARENOP, buffer[:1], buffer[1:]
	}

	// if it is a comment, consume to end of line
	if buffer[0] == ';' {
		for len(buffer) != 0 && buffer[0] != '\n' {
			buffer = buffer[1:]
		}
		return COMMENT, nil, buffer
	}

	r, w := utf8.DecodeRune(buffer)

	// if it is whitespace (which includes invalid runes but not new-lines),
	// return the entire run of whitespace.
	if isspace(r) {
		for len(buffer) != 0 && isspace(r) {
			lexeme, buffer = append(lexeme, buffer[:w]...), buffer[w:]
			r, w = utf8.DecodeRune(buffer)
		}
		return SPACES, lexeme, buffer
	}

	// is it quoted text?
	if r == '"' {
		lexeme, buffer = append(lexeme, buffer[:w]...), buffer[w:]
		r, w = utf8.DecodeRune(buffer)
		for len(buffer) != 0 && r != '\n' && r != '"' {
			if r == '\\' {
				// escaped quotes are accepted as part of the quoted string
				if len(buffer) > 0 || buffer[1] == '"' {
					// consume the escape character
					buffer = buffer[w:]
					r, w = utf8.DecodeRune(buffer)
				}
			}
			lexeme, buffer = append(lexeme, buffer[:w]...), buffer[w:]
			r, w = utf8.DecodeRune(buffer)
		}
		if r == '"' {
			lexeme, buffer = append(lexeme, buffer[:w]...), buffer[w:]
			r, w = utf8.DecodeRune(buffer)
		}
		return QTEXT, lexeme, buffer
	}

	// is it an integer or integer followed by a percent sign?
	if unicode.IsDigit(r) || ((r == '-' || r == '+') && (len(buffer) != 0 && isdigit(buffer[1]))) {
		kind = INTEGER

		lexeme, buffer = append(lexeme, buffer[:w]...), buffer[w:]
		r, w = utf8.DecodeRune(buffer)

		for len(buffer) != 0 && unicode.IsDigit(r) {
			lexeme, buffer = append(lexeme, buffer[:w]...), buffer[w:]
			r, w = utf8.DecodeRune(buffer)
		}
		if r == '.' && len(buffer) > 0 && isdigit(buffer[1]) { // may be a floating point number
			kind = FLOAT
			lexeme, buffer = append(lexeme, buffer[:w]...), buffer[w:]
			r, w = utf8.DecodeRune(buffer)
			for len(buffer) != 0 && unicode.IsDigit(r) {
				lexeme, buffer = append(lexeme, buffer[:w]...), buffer[w:]
				r, w = utf8.DecodeRune(buffer)
			}
		} else if r == '%' {
			kind = PERCENTAGE
			lexeme, buffer = append(lexeme, buffer[:w]...), buffer[w:]
			r, w = utf8.DecodeRune(buffer)
		}
		// must be terminated by a delimiter (space, eol, eof, comma, parentheses, comment)
		if r == '\n' || isspace(r) || r == ',' || r == '(' || r == ')' || r == ';' {
			return kind, lexeme, buffer
		}
		// it's not an integer or percentage, so fall through
	}

	// the lexeme is everything up to the next space or new-line
	for len(buffer) != 0 && !(r == '\n' || isspace(r)) {
		lexeme, buffer = append(lexeme, buffer[:w]...), buffer[w:]
		r, w = utf8.DecodeRune(buffer)
	}

	// most comparisons want lowercase text
	lval := strings.ToLower(string(lexeme))

	// is it a population group?
	switch lval {
	case "civilian":
		return POPULATION, lexeme, buffer
	case "construction-crew":
		return POPULATION, lexeme, buffer
	case "professional":
		return POPULATION, lexeme, buffer
	case "soldier":
		return POPULATION, lexeme, buffer
	case "spy":
		return POPULATION, lexeme, buffer
	case "unskilled-worker", "unsk":
		return POPULATION, lexeme, buffer
	}

	// is it a resource?
	switch lval {
	case "fuel":
		return RESOURCE, lexeme, buffer
	case "gold":
		return RESOURCE, lexeme, buffer
	case "metallics":
		return RESOURCE, lexeme, buffer
	case "non-metallics":
		return RESOURCE, lexeme, buffer
	}

	// is it a deposit id?
	if strings.HasPrefix(lval, "dp-") {
		if _, err := strconv.Atoi(lval[3:]); err == nil {
			return DEPOSITID, lexeme, buffer
		}
	}

	// is it a factory group?
	if strings.HasPrefix(lval, "fg-") {
		if _, err := strconv.Atoi(lval[3:]); err == nil {
			return FACTGRP, lexeme, buffer
		}
	}

	// is it a mining group?
	if strings.HasPrefix(lval, "mg-") {
		if _, err := strconv.Atoi(lval[3:]); err == nil {
			return MINEGRP, lexeme, buffer
		}
	}

	// is it research?
	if lval == "research" {
		return RESEARCH, lexeme, buffer
	} else if strings.HasPrefix(lval, "tl-") {
		if _, err := strconv.Atoi(lval[3:]); err == nil {
			return RESEARCH, lexeme, buffer
		}
	}

	// product will be xxx, xxx-yyy, or xxx-yyy-tl
	// if product includes tl, we must extract it.
	var product string
	if fields := strings.Split(lval, "-"); len(fields) == 1 {
		// product is xxx
		product = lval
	} else {
		// product is xxx-yyy-tl or xxx-yyy
		firstFields, lastField := fields[:len(fields)-1], fields[len(fields)-1]
		if _, err := strconv.Atoi(lastField); err == nil {
			// product is xxx-yyy-tl
			product = strings.Join(firstFields, "-")
		} else {
			// product is xxx-yyy
			product = lval
		}
	}
	switch product {
	case "anti-missile":
		return PRODUCT, lexeme, buffer
	case "assault-craft":
		return PRODUCT, lexeme, buffer
	case "assault-weapons":
		return PRODUCT, lexeme, buffer
	case "automation":
		return PRODUCT, lexeme, buffer
	case "consumer-goods":
		return PRODUCT, lexeme, buffer
	case "energy-shield":
		return PRODUCT, lexeme, buffer
	case "energy-weapon":
		return PRODUCT, lexeme, buffer
	case "factory":
		return PRODUCT, lexeme, buffer
	case "farm":
		return PRODUCT, lexeme, buffer
	case "food":
		return PRODUCT, lexeme, buffer
	case "hyper-engine":
		return PRODUCT, lexeme, buffer
	case "life-support":
		return PRODUCT, lexeme, buffer
	case "lsu", "light-structural-unit":
		return PRODUCT, lexeme, buffer
	case "military-robot":
		return PRODUCT, lexeme, buffer
	case "military-supplies":
		return PRODUCT, lexeme, buffer
	case "mine":
		return PRODUCT, lexeme, buffer
	case "missile":
		return PRODUCT, lexeme, buffer
	case "missile-launcher":
		return PRODUCT, lexeme, buffer
	case "sensor":
		return PRODUCT, lexeme, buffer
	case "space-drive":
		return PRODUCT, lexeme, buffer
	case "su", "structural-unit":
		return PRODUCT, lexeme, buffer
	case "slsu", "super-light-structural-unit":
		return PRODUCT, lexeme, buffer
	case "transport":
		return PRODUCT, lexeme, buffer
	}

	// otherwise it is just plain text
	return TEXT, lexeme, buffer
}

// isdigit returns true if the byte is a digit
func isdigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// isspace returns true if the rune is a space or invalid rune.
// A new-line is not considered a space.
func isspace(r rune) bool {
	return r != '\n' && (unicode.IsSpace(r) || r == utf8.RuneError)
}
