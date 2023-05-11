// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package main

import (
	"flag"
	"fmt"
	"github.com/mdhender/wraithh/parsers/orders"
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
	input, err := os.ReadFile("orders.txt")
	if err != nil {
		return err
	}
	lexemes, err := orders.Scan(input)
	if err != nil {
		return err
	}
	for n, ods := range orders.Parse(lexemes) {
		fmt.Println(n, ods)
	}

	return nil
}
