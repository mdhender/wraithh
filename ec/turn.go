// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package ec

import (
	"fmt"
	"log"
	"sort"
)

type Orders struct {
	Validated bool
	Handle    string
	Game      string
	Turn      int
	Secret    *Secret
	Orders    []Order
	Error     error
}

func (e *Engine) AddOrders(orders []Order) error {
	eo := &Orders{Orders: orders}
	// gather secrets
	for _, order := range eo.Orders {
		if secret, ok := order.(*Secret); ok {
			if eo.Secret != nil {
				return fmt.Errorf("multiple secrets")
			}
			eo.Secret = secret
			eo.Handle = secret.Handle
			eo.Game = secret.Game
			eo.Turn = secret.Turn
		}
	}
	if eo.Secret == nil {
		return fmt.Errorf("missing secret")
	}
	e.Orders = append(e.Orders, eo)
	return nil
}

func (e *Engine) Process() error {
	// process secrets phase
	for _, orders := range e.Orders {
		err := e.SecretsPhase(orders)
		if err != nil {
			// any error with secrets means the order file should be skipped
			orders.Validated = false
		}
		if orders.Validated {
			log.Printf("secrets: validated %s\n", orders.Handle)
		} else if orders.Error == nil {
			log.Printf("secrets: failed    %s\n", orders.Handle)
		} else {
			log.Printf("secrets: failed    %s %v\n", orders.Handle, orders.Error)
		}
	}

	// sort orders by handle for consistent processing in future phases
	sort.Slice(e.Orders, func(i, j int) bool {
		if !e.Orders[i].Validated {
			return false
		}
		return e.Orders[i].Handle < e.Orders[j].Handle
	})

	return nil
}

func (e *Engine) SecretsPhase(orders *Orders) error {
	if orders.Secret == nil {
		orders.Error = fmt.Errorf("missing secret")
		return nil
	}
	secret := orders.Secret
	orders.Handle = secret.Handle
	orders.Game = secret.Game
	orders.Turn = secret.Turn

	switch secret.Token {
	case "003d626a-27c9-4f92-80f3-880384f22d08":
		if orders.Handle != "mdhender" {
			orders.Error = fmt.Errorf("invalid secret")
			return nil
		}
		if orders.Game != "G1" {
			orders.Error = fmt.Errorf("invalid game")
			return nil
		}
		if orders.Turn != 5 {
			orders.Error = fmt.Errorf("invalid turn")
			return nil
		}
		orders.Validated = true
	default:
		orders.Error = fmt.Errorf("invalid secret")
		return nil
	}
	return nil
}
