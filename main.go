// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package main

import (
	"fmt"
	"github.com/mdhender/wraithh/orders"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.LUTC)
	ords, err := orders.ParseFile("mine.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, ord := range ords {
		fmt.Println(ord)
	}
}
