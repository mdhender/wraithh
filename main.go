// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package main

import (
	"fmt"
	"github.com/mdhender/wraithh/adapters"
	"github.com/mdhender/wraithh/cli"
	"github.com/mdhender/wraithh/ec"
	"github.com/mdhender/wraithh/parsers/orders"
	"log"
	"os"
	"time"
)

func main() {
	started := time.Now()

	log.SetFlags(log.LstdFlags | log.LUTC)

	if err := dotfiles("WRAITH"); err != nil {
		log.Fatalf("main: %+v\n", err)
	}

	rv := 0
	if err := cli.Execute(); err != nil {
		log.Printf("\n%+v\n", err)
		rv = 2
	}

	log.Printf("\n")
	log.Printf("completed in %v\n", time.Now().Sub(started))

	os.Exit(rv)
}

func run(path string, debug bool) error {
	e, err := ec.LoadGame(path)
	if err != nil {
		return err
	}

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

	err = e.Process()
	if err != nil {
		return err
	}

	return e.SaveGame(path)
}
