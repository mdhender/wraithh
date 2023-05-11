// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package main

import (
	"flag"
	"fmt"
	"github.com/mdhender/wraithh/internal/adapters"
	"github.com/mdhender/wraithh/internal/orders"
	"github.com/mdhender/wraithh/internal/tokenizer"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.LUTC)

	debug := flag.Bool("debug", false, "log detailed parse output")
	flag.Parse()

	if err := run(*debug); err != nil {
		log.Fatal(err)
	}
}

func run(debug bool) error {
	if debug {
		fmt.Println("grammar:", orders.Grammar)
	}

	input, err := os.ReadFile("orders.txt")
	if err != nil {
		return err
	}
	lexemes, err := adapters.Scan(input)
	if err != nil {
		return err
	}
	for n, ods := range adapters.Parse(lexemes) {
		fmt.Println(n, ods)
	}
	if len(lexemes) >= 0 {
		return nil
	}

	tokens := tokenizer.RemoveEmptyLines(tokenizer.RemoveSpaces(tokenizer.RemoveComments(tokenizer.Tokens(input))))
	if debug {
		for _, t := range tokens {
			fmt.Printf("%3d: %10s %q\n", t.Line, t.Kind, t.Text)
		}
	}

	ords, err := orders.Parse(tokens, true, debug)
	if err != nil {
		return err
	}
	fmt.Println(ords)

	return nil
}
