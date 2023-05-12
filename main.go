// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package main

import (
	"flag"
	"fmt"
	"github.com/mdhender/wraithh/adapters"
	"github.com/mdhender/wraithh/ec"
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
	e := &ec.Engine{}

	// load all the files
	for _, name := range []string{"orders.txt"} {
		input, err := os.ReadFile(name)
		if err != nil {
			return err
		}
		lexemes, err := orders.Scan(input)
		if err != nil {
			return err
		}
		ods := orders.Parse(lexemes)
		if debug {
			for _, od := range ods {
				fmt.Println(od)
			}
		}
		err = e.AddOrders(adapters.OrdersToEngineOrders(ods))
		if err != nil {
			log.Printf("%s: %v\n", name, err)
		}
	}

	err := e.Process()
	if err != nil {
		return err
	}

	return nil
}
