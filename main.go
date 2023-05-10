// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package main

import (
	"flag"
	"fmt"
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
	tokens := tokenizer.RemoveEmptyLines(tokenizer.RemoveSpaces(tokenizer.RemoveComments(tokenizer.Tokens(input))))
	if debug {
		for _, t := range tokens {
			fmt.Printf("%3d: %10s %q\n", t.Line, t.Kind, string(t.Text))
		}
	}

	parseTree, debugTree, err := orders.Parse(tokens, true)
	if debug {
		fmt.Print("Debug Tree:\n\n", debugTree)
	}
	if err != nil {
		return err
	}
	if debug {
		fmt.Print("Parse Tree:\n\n", parseTree)
	}

	return nil
}
