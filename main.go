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
	"path/filepath"
)

func main() {
	log.SetFlags(log.LstdFlags | log.LUTC)

	path := flag.String("path", "", "path to game data")
	debug := flag.Bool("debug", false, "log detailed parse output")
	flag.Parse()
	if path == nil {
		log.Fatal("missing path")
	} else if *path = filepath.Clean(*path); *path == "." {
		log.Fatal("path can't be .")
	} else {
		log.Printf("path is %q\n", *path)
	}

	if err := run(*path, *debug); err != nil {
		log.Fatal(err)
	}
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
